package service

import (
	"github.com/zhou-Qingzhang/gin-admin/service/example"
	"github.com/zhou-Qingzhang/gin-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
