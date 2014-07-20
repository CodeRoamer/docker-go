package api

import (
	"testing"
	"fmt"
)


func TestDClient_do(t *testing.T) {
	client, err := NewDClient("http://42.96.195.83:4243")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, status, err := client.do("Get", "/images/json?all=0", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res, status)
}
