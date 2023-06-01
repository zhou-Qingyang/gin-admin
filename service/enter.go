package service

import (
	"github.com/zhou-Qingzhang/gin-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
