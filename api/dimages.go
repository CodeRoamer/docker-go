package api

import "encoding/json"

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

func (client *DClient) ListImages(param ListImagesAPI_Query, images *[]ListImagesAPI_Resp /*return*/) (err error) {

	str_result, err := client.Do(ListImagesAPI, param,  "")
	if err != nil {
		return err
	}
	err = raiseForErr(json.Unmarshal(str_result, &images))
	return
}

func (client *DClient) InspectImage(name string, image *InspectImageAPI_Resp /*return*/) ( err error) {
	str_result, err := client.Do(InspectImageAPI, nil, name)
	if err != nil {
		return err
	}
	err = raiseForErr(json.Unmarshal(str_result, image))
	return
}
