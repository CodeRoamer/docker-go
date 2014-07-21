package api

import "testing"

type People struct {
	Hello string `json:"hello"`
	Name  string `json:"name"`
	Age   string `json:"age"`
}

func TestPing(t *testing.T) {
	client, err := NewDClient("http://42.96.195.83:4213", "1.13", 20)

	if err != nil {
		t.Fatal(err.Error())
	}

	str, err := client.Ping()

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(str)
}

func TestListImages(t *testing.T) {
	client, err := NewDClient("http://42.96.195.83:4213", "1.12", 20)

	if err != nil {
		t.Fatal(err.Error())
	}

	str, err := client.ListImages(ListImagesAPI_Query {
		All : true,
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(str)
}
