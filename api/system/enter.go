package system

import "github.com/zhou-Qingzhang/gin-admin/service"

type ApiGroup struct {
	BaseApi //登录 用户管理
	AuthorityApi
	AuthorityMenuApi // 动态路由
	// JwtApi
	// DBApi
	// SystemApi
	// CasbinApi
	// AutoCodeApi
	// SystemApiApi
	// DictionaryApi
	// OperationRecordApi
	// AutoCodeHistoryApi
	// DictionaryDetailApi
	// AuthorityBtnApi
}

var (
	menuService      = service.ServiceGroupApp.SystemServiceGroup.MenuService // 显示动态获取路由
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	baseMenuService  = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService

// authorityBtnService = service.ServiceGroupApp.SystemServiceGroup.AuthorityBtnService
// jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
// initDBService           = service.ServiceGroupApp.SystemServiceGroup.InitDBService
// casbinService           = service.ServiceGroupApp.SystemServiceGroup.CasbinService
// autoCodeService         = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
// baseMenuService         = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
//
// dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
// systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
// operationRecordService  = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
// autoCodeHistoryService  = service.ServiceGroupApp.SystemServiceGroup.AutoCodeHistoryService
// dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
//
)
