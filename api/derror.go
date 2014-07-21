package api

import (
	"fmt"
)


func GetGeneralStatusError(code int, module int) (res string) {
	switch code{
	case 200, 201, 204:
		res = "no error"
	case 500:
		res = "server error"
	case 400:
		res = "bad parameter"
	case 404:
		if module == Containers {
			res = "no such container or request 404"
		}else if module == Images {
			res = "no such image or request 404"
		}else {
			res = "no possible error"
		}
	case 406:
		if module == Containers {
			res = "impossible to attach (container not running)"
		}else {
			res = "no possible error"
		}
	case 304:
		if module == Containers {
			res = "container already in the state you request"
		}
	case 409:
		res = "conflict"
	default:
		res = "no possible error"
	}
	return
}


type APIError struct {
	ErrorMessage string
	StatusCode int
	Explanation string
}

func (err APIError) Error() string {
	return fmt.Sprintf("%d %s: %s", err.StatusCode, err.ErrorMessage, err.Explanation)
}

func (err APIError) IsClientError() bool {
	return err.StatusCode >= 400 && err.StatusCode < 500
}

func (err APIError) IsServerError() bool {
	return err.StatusCode >= 500 && err.StatusCode < 600
}

