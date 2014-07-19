package api

import (
	"net/url"
	"net/http"
	"encoding/json"
	"bytes"
	"net"
	"net/http/httputil"
	"io/ioutil"
	"io"
	"errors"

	"github.com/coderoamer/docker-go/utils"
)

type DClient struct {
	endpoint            string
	endpointURL         *url.URL
	client              *http.Client
}


func NewDClient(endpoint string) (*DClient, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	return &DClient{endpoint, u, http.DefaultClient}, nil
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

	req, err := http.NewRequest(method, path, params)
	if err != nil {
		return nil, -1, err
	}
	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	} else if method == "POST" {
		req.Header.Set("Content-Type", "plain/text")
	}
	var resp *http.Response
	protocol := c.endpointURL.Scheme
	address := c.endpointURL.Path

	switch protocol {
	case "unix":
		dial, err := net.Dial(protocol, address)
		if err != nil {
			return nil, -1, err
		}
		defer dial.Close()
		clientconn := httputil.NewClientConn(dial, nil)
		resp, err = clientconn.Do(req)
		if err != nil {
			return nil, -1, err
		}
		defer clientconn.Close()
	case "http":
		resp, err = c.client.Do(req)
	}

	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, resp.StatusCode, errors.New(utils.ReqError)
	}
	return body, resp.StatusCode, nil
}
