package api

import (
	"fmt"
	"encoding/json"
	"time"
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

type DImage struct {
	ID              string    `json:"id"`
	Parent          string    `json:"parent,omitempty"`
	Comment         string    `json:"comment,omitempty"`
	Created         time.Time `json:"created"`
	Container       string    `json:"container,omitempty"`
//	ContainerConfig Config    `json:"containerconfig,omitempty"`
	DockerVersion   string    `json:"dockerversion,omitempty"`
	Author          string    `json:"author,omitempty"`
//	Config          *Config   `json:"config,omitempty"`
	Architecture    string    `json:"architecture,omitempty"`
	Size            int64
}

type ProgressDetail struct {
	Current	int
	Total		int
}

type DImageStatus struct {
	Id					string
	Status				string
	progressDetail	ProgressDetail
}

func (client *DClient) ListImages(all bool) ([]APIImages, error) {
	arg := fmt.Sprintf("all=%v", all)
	api := GetImagesApi(Images, List, "get", "/images/json", arg, "application/json")
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

//set nil if not needed
func (client *DClient) CreateImages(fromImage, tag, frmSrc, repo, registry string )(error){
	var arg string
	if fromImage != "" {
		arg = fmt.Sprintf("fromImage=%s&", fromImage)
	}
	if tag != "" {
		arg = fmt.Sprintf("%stag=%s&", arg, tag)
	}
	if frmSrc != "" {
		arg = fmt.Sprintf("%sfrmSrc=%s&", arg, frmSrc)
	}
	if repo != "" {
		arg = fmt.Sprintf("%srepo=%s&", arg, repo)
	}
	if registry != "" {
		arg = fmt.Sprintf("%s registry=%s", arg, registry)
	}
	api := GetImagesApi(Images, Create, "post", "/images/create", arg, "application/json")
	body, err := client.Do(api)
	if err != nil {
		return  err
	}
	status := new(DImageStatus)
	err = json.Unmarshal(body, status)
	if err != nil {
		return err
	}
	fmt.Println(status)
	return nil
}
