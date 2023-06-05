package system

import "github.com/zhou-Qingzhang/gin-admin/service"

type ApiGroup struct {
	BaseApi
	// JwtApi
	// DBApi
	// SystemApi
	// CasbinApi
	// AutoCodeApi
	// SystemApiApi
	// AuthorityApi
	// DictionaryApi
	AuthorityMenuApi
	// OperationRecordApi
	// AutoCodeHistoryApi
	// DictionaryDetailApi
	// AuthorityBtnApi
}

var (
	menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	// jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
// initDBService           = service.ServiceGroupApp.SystemServiceGroup.InitDBService
// casbinService           = service.ServiceGroupApp.SystemServiceGroup.CasbinService
// autoCodeService         = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
// baseMenuService         = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
// authorityService        = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
// dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
// systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
// operationRecordService  = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
// autoCodeHistoryService  = service.ServiceGroupApp.SystemServiceGroup.AutoCodeHistoryService
// dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
// authorityBtnService     = service.ServiceGroupApp.SystemServiceGroup.AuthorityBtnService
)
