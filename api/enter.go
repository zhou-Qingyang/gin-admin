package api

import "github.com/zhou-Qingzhang/gin-admin/api/system"

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	// ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
