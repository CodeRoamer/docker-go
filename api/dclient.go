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

type DClient struct {
	endpoint			string
	endpointURL		*url.URL
	client				*http.Client
	scheme				string
}

func NewDClient(endpoint string) (*DClient, error) {
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
		scheme: u.Scheme}, nil
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
	if c.scheme == "http" {
		path = fmt.Sprintf("%s%s", c.endpointURL.String(), path)
	}
	req, err := http.NewRequest(method, path, params)
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
//		result, status, err = c.do("Get", fmt.Sprintf("%s?%s", api.ReqUrl, api.ReqArg), api.ContentType, nil)
		result, status, err = c.do("Get", api.ReqUrl, api.ContentType, nil)
	}else if strings.ToLower(api.Method) == "post" {
		//TODO
	}
	retError := GetGeneralStatusError(status, api)
	if  retError == NoError {
		return result, err
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
