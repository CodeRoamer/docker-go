package api

import "strings"

const (
	Images 		= 1
	Containers 	= 2
)

const (
	List		= 1
	Create		= 2
	Insert 	= 3
	Inspect 	= 4
	Start		= 5
	Stop		= 6

)

type ModuleAPI struct {
	Module int				//	container | images | misc
	Api int 				//	List
	Method string		 	//	POST & GET
	ReqUrl string			//	/containers/json
	ReqArg string			//	all=%d&before=%s&size=%d
	ContentType string		//	application/json
	Result interface {}
}

func GetImagesApi(module, api int, method, reqUrl, reqArg, contentType string ) *ModuleAPI {
	return &ModuleAPI {
		module, api, strings.ToUpper(method), reqUrl, reqArg, contentType, nil,
	}
}
