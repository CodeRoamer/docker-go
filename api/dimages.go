package api

import "fmt"

const (
	listImages = "/images/json?all=%t"
)

func (client *DClient) ListImages(all bool)  {
	body, status, err := client.Do("GET", fmt.Sprintf(listImages, all), nil)
	fmt.Println(status, err)
	fmt.Printf("%s", body)
}
