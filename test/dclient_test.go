package test

import (
	"testing"
	"fmt"
	"github.com/coderoamer/docker-go/api"
)

func TestListImages(t *testing.T) {
	client, err := api.NewDClient("http://222.26.224.50:4213")
	if err != nil {
		fmt.Println(err)
		return
	}
	client.ListImages(true)
}
