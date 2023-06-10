package system

import "github.com/zhou-Qingzhang/gin-admin/service"

type ApiGroup struct {
	BaseApi //登录 用户管理
	AuthorityApi
	AuthorityMenuApi // 动态路由
	OperationRecordApi
	DictionaryApi
	DictionaryDetailApi
	SystemApi
}

var (
	menuService             = service.ServiceGroupApp.SystemServiceGroup.MenuService // 显示动态获取路由
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	authorityService        = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	baseMenuService         = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	operationRecordService  = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
	systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
)
