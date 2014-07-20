package test

import (
	"testing"
	"github.com/coderoamer/docker-go/api"
)

func TestListImages(t *testing.T) {
	client, err := api.NewDClient("http://42.96.195.83:4213")
	if err != nil {
		t.Error(err)
		return
	}
	images, err := client.ListImages(true)
	if err != nil {
		t.Error(err)
		return
	}
	if images == nil {
		t.Error("images not fetch")
	}
}

func TestPing(t *testing.T) {
	client, err := api.NewDClient("http://42.96.195.83:4213")
	if err != nil {
		t.Error(err)
		return
	}
	err = client.Ping()
	if err != nil {
		t.Error("ping error")
		return
	}
	client, err = api.NewDClient("http://42.96.195.83:4214")
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
