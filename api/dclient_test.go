package api

import (
	"testing"
	"io/ioutil"
	"strings"
)

var host = "http://42.96.195.83:4213"

func TestPing(t *testing.T) {
	client, err := NewDClient(host, "1.13", 20)

	if err != nil {
		t.Fatal(err.Error())
	}

	ok, err := client.Ping()

	if err != nil {
		t.Fatal(err.Error())
	}

	if !ok {
		t.Fatal("ping failed")
	}
}

func TestPost(t *testing.T) {
	client, err := NewDClient(host, "1.13", 20)
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := client.post(host + "/images/create", "fromImage=base", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err.Error())
	}
	if strings.Contains(string(body), "404") {
		t.Fatal("post test get 404")
	}else {
		t.Logf("%s", body)
	}
	res.Body.Close()
}

func TestListImages(t *testing.T) {
	client, err := NewDClient(host, "1.12", 20)

	if err != nil {
		t.Fatal(err.Error())
	}

	var images []ListImagesAPI_Resp
	err = client.ListImages(ListImagesAPI_Query {
		All : false,
	}, &images)

	if err != nil {
		t.Error(err)
		return
	}

	if images == nil {
		t.Error("images list failed")
	}
}

func TestUrl(t *testing.T) {
	client, _ := NewDClient(host, "1.12", 20)
	if client.url("/path/%s/", "hello") != host + "/v1.12/path/hello/" {
		t.Fatal("url error")
	}
}

