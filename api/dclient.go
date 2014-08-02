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
	"fmt"
	"strings"

	"github.com/Unknwon/com"
	"errors"
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
		// TODO: socketPath = "/"+u.Host+u.Path
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
func (c *DClient) url(path string, arg string) string {
	if arg != "" {
		path = fmt.Sprintf(path, arg)
	}
	return fmt.Sprintf("%s/v%s%s", c.endpoint, c.version, path)
}

// pay attention: path is complete path, should be like this:
// http://endpoint/v1.12/containers
// return APIError
func (c *DClient) get(path string, query interface{}) (*http.Response, error) {

	query_string := ParseStruct2QueryString(query)

	if len(query_string) != 0 {
		path += "?"+query_string
	}

	return c.client.Get(path)
}

// post method, two parts:
// params append to the url, data post in body
func (c *DClient) post(path string, query, jsonParam interface {}) (*http.Response, error) {
	// query string as params
	query_string := ParseStruct2QueryString(query)

	if len(query_string) != 0 {
		path += "?"+query_string
	}
	// post data as body
	var post_data io.Reader = nil
	if jsonParam != nil {
		buf, err := json.Marshal(jsonParam)
		if err = raiseForErr(err); err != nil {
			return nil, err
		}
		post_data = bytes.NewBuffer(buf)
	}

	resp, err := c.client.Post(path, "application/json", post_data)
	if err = raiseForErr(err); err != nil {
		return nil, err
	}
	return resp, nil
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
	if err = raiseForErr(err); err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err = raiseForErr(err); err != nil {
		return nil, err
	}

	return res, nil
}
//param for ?param
//urlParam for /path/%s/..../
//I checkout the docker's doc, I found url only require one param, so I make this args... silly but useful : )
func (client *DClient) Do(module ModuleAPI, param interface{}, urlParam string ) (str_result []byte, err error) {
	if err = checkVersion(module.Version, client.version); err != nil {
		return nil, err
	}
	var resp *http.Response
	switch string(bytes.ToLower([]byte(module.Method))) {
	case "get":
		resp, err = client.get(client.url(module.ReqUrl, urlParam), param)
		if err != nil {
			return nil, err
		}
	case "post":
		resp, err = client.post(client.url(module.ReqUrl, urlParam), param, nil)
		if err != nil {
			return nil, err
		}
	case "delete":
		resp, err = client.delete(client.url(module.ReqUrl, urlParam), param)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Unkown request method.")

	}

	return resultBinary(resp, module.Module)
}

// Ping
func (c *DClient) Ping() (bool, error) {
	resp, err := c.get(c.url("/_ping", ""), nil)
	if err != nil {
		return false, err
	}

	byte_arr, err := resultBinary(resp, 0)

	if strings.Contains(string(byte_arr), "OK") {
		return true, nil
	}else {
		return false, nil
	}
}


// check version
// return APIError
func checkVersion(support []string, curr string) (err error) {
	err = nil

	if !com.IsSliceContainsStr(support, curr) {
		// version not supported
		err = APIError {"docker-go error", 500, "API Version Not Supported"}
	}

	return err
}

// format a response to json(string) or to binary([]byte)
// return APIError
func resultBinary(response *http.Response, module int) ([]byte, error) {
	err := raiseForStatus(response, module)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err = raiseForErr(err); err != nil {
		return nil, err
	}
	return body, nil
}
