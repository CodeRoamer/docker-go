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

	"github.com/coderoamer/docker-go/utils"
	"fmt"
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
func (c *DClient) Do(method, path string, data interface{}) ([]byte, int, error) {
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
	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	} else if method == "POST" {
		req.Header.Set("Content-Type", "plain/text")
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, -1, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 400 {
		return nil, res.StatusCode, errors.New(utils.ReqError)
	}
	return body, res.StatusCode, nil
}

