package api

// modules
const (
	Images     = 1
	Containers = 2
	Misc       = 3
)

type ModuleAPI struct {
	Module			int      //	container | images | misc
	Version		[]string // supported version ["1.11","1.12"]
	ReqUrl			string   // path for API /images/%s/insert

	ResType		string // application/json  or  application/vnd.docker.raw-stream  or  application/octet-stream  or  nothing
	ReqType		string // application/json  or  nothing
	Method			string //get post delete
}

var ListImagesAPI = ModuleAPI {
	Module: Images,
	Version: []string{"1.11", "1.12", "1.13"},
	ReqUrl: "/images/json",

	ResType: "application/json",
	ReqType: "",
	Method:"get",
}


var InspectImageAPI = ModuleAPI {
	Module: Images,
	Version: []string{"1.11", "1.12", "1.13"},
	ReqUrl: "/images/%s/json",

	ResType: "application/json",
	ReqType: "",
	Method:"get",
}

var InsertFileAPI = ModuleAPI {
	Module: Images,
	Version: []string{"1.11","1.12"},
	ReqUrl: "/images/%s/insert",

	ResType: "application/json",
	ReqType: "",
	Method:"post",
}
