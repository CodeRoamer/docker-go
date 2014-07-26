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

	var images []ListImagesAPI_Resp
	err = client.ListImages(ListImagesAPI_Query {
		All : false,
	}, &images)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(images)
}

func TestInspectImages(t *testing.T) {
	var image InspectImageAPI_Resp
	client, err := NewDClient("http://42.96.195.83:4213", "1.13", 20)

	if err != nil {
		t.Fatal(err.Error())
	}
	err = client.InspectImage("ubuntu", &image)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(image)
}
