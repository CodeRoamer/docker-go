package test

import (
	"testing"
	"fmt"
	"github.com/coderoamer/docker-go/api"
)

func TestListImages(t *testing.T) {
	client, err := api.NewDClient("http://127.0.0.1:4213")
	if err != nil {
		fmt.Println(err)
		return
	}
	client.ListImages(true)
}
