package system

import (
	"github.com/gin-gonic/gin"
	api "github.com/zhou-Qingzhang/gin-admin/api"
	"github.com/zhou-Qingzhang/gin-admin/middleware"
)

type DictionaryRouter struct{}

func (s *DictionaryRouter) InitSysDictionaryRouter(Router *gin.RouterGroup) {
	sysDictionaryRouter := Router.Group("sysDictionary").Use(middleware.OperationRecord())
	sysDictionaryRouterWithoutRecord := Router.Group("sysDictionary")
	sysDictionaryApi := api.ApiGroupApp.SystemApiGroup.DictionaryApi
	{
		sysDictionaryRouter.POST("createSysDictionary", sysDictionaryApi.CreateSysDictionary)   // 新建SysDictionary
		sysDictionaryRouter.DELETE("deleteSysDictionary", sysDictionaryApi.DeleteSysDictionary) // 删除SysDictionary
		sysDictionaryRouter.PUT("updateSysDictionary", sysDictionaryApi.UpdateSysDictionary)    // 更新SysDictionary
	}
	{
		sysDictionaryRouterWithoutRecord.GET("findSysDictionary", sysDictionaryApi.FindSysDictionary)       // 根据ID获取SysDictionary
		sysDictionaryRouterWithoutRecord.GET("getSysDictionaryList", sysDictionaryApi.GetSysDictionaryList) // 获取SysDictionary列表
	}
}
