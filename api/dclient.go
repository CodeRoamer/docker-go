// 	Copyright 2013 CodeHolic org
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.


package api

import (
	"net/url"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"io"
	"errors"
	"fmt"
	"strings"
)

const (
	DefaultDockerAPIVersion = "1.13"
	DefaultTimeoutSeconds   = 60
	StreamHeaderSizeBytes   = 8
)

type DClient struct {
	endpoint              string
	endpointURL           *url.URL
	client                *http.Client
	scheme                string
	version               string
	timeout               int
}


// create a new DClient with the given endpoint and version,
// with additional timeout param
func NewDClient(endpoint, version string, timeout int) (*DClient, error) {
	var client *http.Client

	// with trailing slash?
	if strings.HasSuffix(endpoint, "/") {
		endpoint = strings.TrimSuffix(endpoint, "/")
	}

	// parse endpoint to url struct
	endpoint_url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	switch endpoint_url.Scheme{
	case "unix":
		//		httpTransport := &http.Transport{}
		//		// TODO: socketPath = "/"+u.Host+u.Path
		//		socketPath := u.Path
		//		unixDial := func(/*proto string, addr string*/) (net.Conn, error) {
		//			return net.Dial("unix", socketPath)
		//		}
		//		httpTransport.Dial = unixDial
		//		// Override the main URL object so the HTTP lib won't complain
		//		client = &http.Client{Transport: httpTransport}
		client = nil
	case "http":
		client = http.DefaultClient
	}

	return &DClient{
		endpoint: endpoint,
		endpointURL: endpoint_url,
		client: client,
		scheme: endpoint_url.Scheme,
		version: version,
		timeout: timeout,
	}, nil
}

// create a url with the given path
// form with a endpoint, api version and path
func (c *DClient) url(path string) string {
	return fmt.Sprintf("%s/v%s%s", c.endpoint, c.version, path)
}


// format a response to json(string) or to binary([]byte)
func (c *DClient) resultBinary(response *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err {
		return nil, err
	}
	return body, nil
}

// format a response to json(string) or to binary([]byte)
func (c *DClient) resultJson(response *http.Response) (string, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err {
		return nil, err
	}
	return string(body), nil
}

// raise an error for http status
func (c *DClient) raiseForStatus(statusCode int, module ModuleAPI)  (err error) {
	err = nil

	if statusCode < 400 {
		err = errors.New(module.StatusMap[statusCode])
	}

	return err
}

// pay attention: path is complete path, should be like this:
// http://endpoint/v1.12/containers
func (c *DClient) get(path string, query interface{}) (*http.Response, error) {

	query_string := ParseStruct2QueryString(query)

	if len(query_string) != 0 {
		path += "?"+query_string
	}

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	return c.client.Do(req)

}


// post method, two parts:
// params append to the url, data post in body
func (c *DClient) post(path string, query, json interface {}) (*http.Response, error) {
	// query string as params
	query_string := ParseStruct2QueryString(query)

	if len(query_string) != 0 {
		path += "?"+query_string
	}

	// post data as body
	var post_data io.Reader = nil
	if json != nil {
		buf, err := json.Marshal(json)
		if err != nil {
			return nil, err
		}
		post_data = bytes.NewBuffer(buf)
	}

	// create request
	req, err := http.NewRequest("POST", path, post_data)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return c.client.Do(req)
}


// delete
// pay attention: path is complete path, should be like this:
// http://endpoint/v1.12/containers
func (c *DClient) delete(path string, options interface{}) (*http.Response, error) {

	query_string := ParseStruct2QueryString(options)

	if len(query_string) != 0 {
		path += "?"+query_string
	}

	req, err := http.NewRequest("DELETE", path, nil)

	if err != nil {
		return nil, err
	}

	return c.client.Do(req)

}



//args: method:get/post, path:request path data:post data(json data)
//return: body, status, error
func (c *DClient) do(method, path, contentType string, data interface{}) ([]byte, int, error) {
	var params io.Reader
	if data != nil {
		buf, err := json.Marshal(data)
		if err != nil {
			return nil, -1, err
		}
		params = bytes.NewBuffer(buf)
	}
	path = fmt.Sprintf("/%s/%s", c.version, path)
	if c.scheme == "http" {
		path = fmt.Sprintf("%s%s", c.endpointURL.String(), path)
	}
	req, err := http.NewRequest(strings.ToUpper(method), path, params)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Set("Content-Type", contentType)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, res.StatusCode, err
}




func (c *DClient) Do(api *ModuleAPI) ([]byte, error) {
	var result []byte
	var status int
	var err error

	if strings.ToLower(api.Method) == "get" {
		result, status, err = c.do(api.Method, fmt.Sprintf("%s?%s", api.ReqUrl, api.ReqArg), api.ContentType, nil)

	}else if strings.ToLower(api.Method) == "post" {
		//TODO
	}
	retError := GetGeneralStatusError(status, api)
	if err != nil {
		return nil, err
	}
	if retError == NoError {

		return result, nil
	}else {

		return nil, errors.New(retError)
	}
}

func (c *DClient) Ping() error {
	path := "/_ping"
	_, status, err := c.do("GET", path, "application/json", nil)
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		return nil
	}
	return nil
}
