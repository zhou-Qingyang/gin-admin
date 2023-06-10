package system

type RouterGroup struct {
	BaseRouter
	UserRouter
	MenuRouter
	AuthorityRouter
	OperationRecordRouter
	DictionaryRouter
	DictionaryDetailRouter
	SysRouter
}
