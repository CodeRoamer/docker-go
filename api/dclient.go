// Copyright 2013 CodeHolic org.
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.


package api

import (
	"net/url"
	"net/http"
	"encoding/json"
	"bytes"
	"net"
	"io/ioutil"
	"io"
	"errors"
	"fmt"
	"strings"
)

const (
	DefaultDockerAPIVersion = "1.13"
	DefaultTimeoutSeconds = 60
	StreamHeaderSizeBytes = 8
)

type DClient struct {
	endpoint			string
	endpointURL		*url.URL
	client				*http.Client
	scheme				string
	version			string
	timeout			int
}

func NewDClient(endpoint, version string, timeout int) (*DClient, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	var client *http.Client

	switch u.Scheme{
	case "unix":
		httpTransport := &http.Transport{}
		socketPath := u.Path
		unixDial := func(proto string, addr string) (net.Conn, error) {
			return net.Dial("unix", socketPath)
		}
		httpTransport.Dial = unixDial
		// Override the main URL object so the HTTP lib won't complain
		client = &http.Client{Transport: httpTransport}
	case "http":
		client = http.DefaultClient
	}
	u.Path = ""

	return &DClient{
		endpoint: endpoint,
		endpointURL: u,
		client: client,
		scheme: u.Scheme,
		version: version,
		timeout: timeout}, nil
}

func (c *DClient) url(path string) string {
	return fmt.Sprintf("%s/v%s%s", c.endpoint, c.version, path)
}

func (c *DClient) get(path string, options interface {}) (*http.Response, error) {
	var param url.Values

	if options != nil {
		if buf, err := json.Marshal(options); err == nil {
			var data map[string]string

			fmt.Println(string(buf))

			if err = json.Unmarshal([]byte(string(buf)), &data); err == nil {
				for key, value := range data {
					param.Add(key, value)
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	req.Form = param

	return c.client.Do(req)

}

//args: method:get/post, path:request path data:post data(json data)
//return: body, status, error
func (c *DClient) do(method, path, contentType string, data []byte) ([]byte, int, error) {
	buffer := bytes.NewBuffer(data)
	path = fmt.Sprintf("/%s%s", c.version, path)
	if c.scheme == "http" {
		path = fmt.Sprintf("%s%s", c.endpointURL.String(), path)
	}
	req, err := http.NewRequest(strings.ToUpper(method), path, buffer)
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

	result, status, err = c.do(api.Method,
		fmt.Sprintf("%s?%s", api.ReqUrl, api.ReqArg),
		api.ContentType, nil)

	retError := GetGeneralStatusError(status, api)
	if err != nil {
		return nil, err
	}
	if  retError == NoError {
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
