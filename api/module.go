package api

// modules
const (
	Images     = 1
	Containers = 2
	Misc       = 3
)

type ModuleAPI struct {
	Module        int      //	container | images | misc
	Version       []string // supported version ["1.11","1.12"]
	ReqUrl        string   // path for API /images/%s/insert

	Method        string // GET POST DELETE

	ResType        string // application/json  or  application/vnd.docker.raw-stream  or  application/octet-stream  or  nothing
	ReqType        string // application/json  or  nothing

	StatusMap    map[int]string // {200: "no error", 500: "server error"}
}

var ListImagesAPI = ModuleAPI {
	Module: Images,
	Version: []string{"1.11","1.12","1.13"},
	ReqUrl: "/images/json",

	Method: "GET",
	ResType: "application/json",
	ReqType: nil,
	StatusMap: map[int]string {
		200: "no error",
		500: "server error",
	},
}

