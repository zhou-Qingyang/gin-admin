package system

type ServiceGroup struct {
	UserService
	MenuService

	JwtService
	CasbinService // 第三方 权限

	// ApiService

	// InitDBService
	// AutoCodeService
	// BaseMenuService
	// AuthorityService
	// DictionaryService
	// SystemConfigService
	// AutoCodeHistoryService
	// OperationRecordService
	// DictionaryDetailService
	// AuthorityBtnService
}
