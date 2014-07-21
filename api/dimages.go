package api

import (
	"encoding/json"
	"fmt"
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



type InspectImageAPI_Resp struct {
	Created    string
	Container  string
	ContainerConfig struct {
		Hostname     string
		User         string
		Memory       int
		MemorySwap   int
		AttachStdin  bool
		AttachStdout bool
		AttachStderr bool
		PortSpecs    string
		Tty          bool
		OpenStdin    bool
		StdinOnce    bool
		Env          []string
		Cmd          []string
		Dns          []string
		Image        string
		Volumes      []string
		VolumesFrom  string
		WorkingDir   string
	}
	Id         string
	Parent     string
	Size       int
}

type InsertFileAPI_Query struct {
	Path string `json:"path,omitempty"`
	Url  string`json:"url,omitempty"`
}

type InsertFileAPI_Resp struct {
	Status   string `json:"status,omitempty"`
	Progress string `json:"progress,omitempty"`
	ProgressDetail struct {Current int `json:"current,omitempty"`} `json:"progressDetail,omitempty"`
	Error    string `json:"error,omitempty"`
}

func (client *DClient) ListImages(json_param ListImagesAPI_Query) (images []ListImagesAPI_Resp, err error) {
	if err = checkVersion(ListImagesAPI.Version, client.version); err != nil {
		return
	}

	// get response
	resp, err := client.get(client.url(ListImagesAPI.ReqUrl), json_param)
	if err != nil {
		return
	}

	// get byte, check response
	byte_arr, err := resultBinary(resp, ListImagesAPI.Module)
	if err != nil {
		return
	}

	// marshal bytes into struct
	if err = raiseForErr(json.Unmarshal(byte_arr, &images)); err != nil {
		return
	}

	// return images
	return
}


func (client *DClient) InspectImage(name string) (image InspectImageAPI_Resp, err error) {
	if err = checkVersion(InspectImageAPI.Version, client.version); err != nil {
		return
	}

	// get response
	resp, err := client.get(client.url(fmt.Sprintf(InspectImageAPI.ReqUrl, name)), nil)
	if err != nil {
		return
	}

	// get byte, check response
	byte_arr, err := resultBinary(resp, InspectImageAPI.Module)
	if err != nil {
		return
	}

	// marshal bytes into struct
	if err = raiseForErr(json.Unmarshal(byte_arr, &image)); err != nil {
		return
	}

	// return images
	return
}

func (client *DClient) InsertFile(name string, query_param InsertFileAPI_Query) (ret InsertFileAPI_Resp, err error) {
	if err = checkVersion(InsertFileAPI.Version, client.version); err != nil {
		return
	}

	// get response
	resp, err := client.post(client.url(fmt.Sprintf(InsertFileAPI.ReqUrl, name)), query_param, nil)
	if err != nil {
		return
	}

	// get byte, check response
	byte_arr, err := resultBinary(resp, InsertFileAPI.Module)
	if err != nil {
		return
	}

	// marshal bytes into struct
	if err = raiseForErr(json.Unmarshal(byte_arr, &ret)); err != nil {
		return
	}

	// return images
	return
}
