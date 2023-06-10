package system

type ServiceGroup struct {
	UserService
	MenuService
	JwtService
	CasbinService // 第三方 权限
	AuthorityService
	BaseMenuService // 基础菜服务
	OperationRecordService
	SystemConfigService //服务器信息
	DictionaryService
	DictionaryDetailService
}
