package api

import "testing"

//func TestDClient_do(t *testing.T) {
//	client, err := NewDClient("unix://var/run/docker.sock", "v1.13", 20)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	res, status, err := client.do("GET", "/images/json", "application/json", nil)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	if status == 404 {
//		t.Error("404 request error")
//		return
//	}
//	if res == nil {
//		t.Error("request empty")
//	}
//
//}

type People struct {
	Hello string `json:"hello"`
	Name string `json:"name"`
	Age int `json:"age"`
}

func TestGet(t *testing.T) {
	client, err := NewDClient("http://127.0.0.1:8000/", "api",20)

	if err != nil {
		t.Error(err)
		return
	}
	status, err := client.get("/announcement", People{
		Hello: "name",
		Name: "lihe",
		Age: 18,
	})

	if err != nil {
		t.Error(err)
		return
	}
	if status == 404 {
		t.Error("404 request error")
		return
	}
}
