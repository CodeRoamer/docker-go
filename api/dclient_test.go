package api

import "testing"

func TestDClient_do(t *testing.T) {
	client, err := NewDClient("http://42.96.195.83:4213", "v1.13")
	if err != nil {
		t.Error(err)
		return
	}
	res, status, err := client.do("GET", "/images/json", "application/json", nil)
	if err != nil {
		t.Error(err)
		return
	}
	if status == 404 {
		t.Error("404 request error")
		return
	}
	if res == nil {
		t.Error("request empty")
	}
}
