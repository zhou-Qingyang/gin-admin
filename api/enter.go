package api

import "github.com/zhou-Qingzhang/gin-admin/api/system"

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
