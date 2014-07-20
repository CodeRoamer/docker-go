package api

import (
	"fmt"
	"encoding/json"
)

type APIImages struct {
	ID          string   `json:"Id"`
	RepoTags    []string `json:",omitempty"`
	Created     int64
	Size        int64
	VirtualSize int64
	ParentId    string `json:",omitempty"`
	Repository  string `json:",omitempty"`
	Tag         string `json:",omitempty"`
}

func (client *DClient) ListImages(all bool) ([]APIImages, error) {
	arg := fmt.Sprintf("all=%v", all)
	api := GetImagesApi(Images, List, "Get", "/images/json", arg, "application/json")
	body, err := client.Do(api)
	if err != nil {
		return  nil, err
	}
	var images []APIImages
	err = json.Unmarshal(body, &images)
	if err != nil {
		return nil, err
	}
	return images, nil
}
