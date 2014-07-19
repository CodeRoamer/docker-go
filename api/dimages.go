package api

import "fmt"

const (
	listImages = "/images/json?all=%t"
)

func (client *DClient) ListImages(all bool)  {
	path := "/images/json?all="
	if all {
		path += "1"
	} else {
		path += "0"
	}
	body,_, _ := client.Do("GET", path, nil)
		fmt.Printf("%s", body)
}
