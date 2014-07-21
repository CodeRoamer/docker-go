package api

import (
	"encoding/json"
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

func (client *DClient) ListImages(json_param ListImagesAPI_Query) ([]ListImagesAPI_Resp, error) {
	if err := checkVersion(ListImagesAPI.Version, client.version); err != nil {
		return nil, err
	}

	// get response
	resp, err := client.get(client.url(ListImagesAPI.ReqUrl), json_param)
	if err != nil {
		return nil, err
	}

	// get byte, check response
	byte_arr, err := resultBinary(resp, ListImagesAPI.Module)
	if err != nil {
		return nil, err
	}

	// marshal bytes into struct
	var images []ListImagesAPI_Resp
	if err = raiseForErr(json.Unmarshal(byte_arr, &images)); err != nil {
		return nil, err
	}

	// return images
	return images, nil
}
