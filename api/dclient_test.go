package api

import (
	"testing"
	"fmt"
)


func TestDClient_do(t *testing.T) {
	client, err := NewDClient("http://127.0.0.1:4213")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, status, err := client.Do("Get", "/images/json?all=0", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res, status)
}
