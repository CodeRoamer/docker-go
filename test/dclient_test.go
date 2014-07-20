package test

import (
	"testing"
	"github.com/coderoamer/docker-go/api"
)

func TestListImages(t *testing.T) {
	client, err := api.NewDClient("http://222.26.224.50:4213")
	if err != nil {
		t.Error(err)
		return
	}
	err = client.ListImages(false)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestPing(t *testing.T) {
	client, err := api.NewDClient("http://222.26.224.50:4213")
	if err != nil {
		t.Error(err)
		return
	}
	err = client.Ping()
	if err != nil {
		t.Error("ping error")
		return
	}

	client, err = api.NewDClient("http://222.26.224.51:4213")
	if err != nil {
		t.Error(err)
		return
	}
	err = client.Ping()
	if err != nil {
		return
	}else {
		t.Error("error ")
	}
}
