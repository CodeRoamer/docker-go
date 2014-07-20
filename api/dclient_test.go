package api

import (
	"testing"
	"fmt"
)


func TestDClient_do(t *testing.T) {
	client, err := NewDClient("http://42.96.195.83:4213")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, status, err := client.do("Get", "/images/json", "application/json", nil)
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
