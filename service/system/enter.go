package system

type ServiceGroup struct {
	UserService
	MenuService
	JwtService
	CasbinService // 第三方 权限
	AuthorityService
	BaseMenuService // 不知道 这个是干嘛的
	// DictionaryService
	// SystemConfigService
	// AutoCodeHistoryService
	// OperationRecordService
	// DictionaryDetailService
	// AuthorityBtnService
}
