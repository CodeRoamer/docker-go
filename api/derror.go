package api

const (
	ReqError = "request error"
	NoError = "Success"
)

func GetGeneralStatusError(code int, module *ModuleAPI) (res string) {
	switch code{
	case 200: case 204:
		res = NoError
	case 500:
		res = "server error"
	case 400:
		res = "bad parameter"
	case 404:
		if module.Module == Containers {
			res = "no such container or request 404"
		}else if module.Module == Images {
			res = "no such image or request 404"
		}else {
			res = "no possible error"
		}
	case 406:
		if module.Module == Containers {
			res = "impossible to attach (container not running)"
		}else {
			res = "no possible error"
		}
	case 304:
		if module.Module == Containers && module.Api == Start {
			res = "container already started"
		}else if module.Module == Containers && module.Api == Stop {
			res = "container already stopped"
		}
	case 409:
		res = "conflict"
	default:
		res = "no possible error"
	}
	return
}
