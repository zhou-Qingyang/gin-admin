package router

import "github.com/zhou-Qingzhang/gin-admin/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
