package router

import (
	"github.com/zhou-Qingzhang/gin-admin/router/example"
	"github.com/zhou-Qingzhang/gin-admin/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
