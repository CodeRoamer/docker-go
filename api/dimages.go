package api

import (
	"encoding/json"

	"github.com/Unknwon/com"
)

type ListImagesAPI_Query struct {
	All bool    `json:"all"`
	Filters map[string][]string `json:"filters,omitempty"`
}

type ListImagesAPI_Resp struct {
	ID          string   `json:"Id"`
	RepoTags    []string `json:",omitempty"`
	Created     int64
	Size        int64
	VirtualSize int64
	ParentId    string `json:",omitempty"`
	Repository  string `json:",omitempty"`
	Tag         string `json:",omitempty"`
}

func (client *DClient) ListImages(json ListImagesAPI_Query) ([]ListImagesAPI_Resp, error) {
	if !com.IsSliceContainsStr(ListImagesAPI.Version, client.version) {
		// version not supported
		return nil, nil
	}
	resp, err := client.get(client.url(ListImagesAPI.ReqUrl), json)


	var images []ListImagesAPI_Resp
	err = json.Unmarshal(client.result_binary(resp), &images)
	if err != nil {
		return nil, err
	}
	return images, nil
}
