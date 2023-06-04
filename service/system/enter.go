package system

type ServiceGroup struct {
	UserService
	JwtService
	CasbinService // 第三方 权限

	// ApiService
	// MenuService
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
