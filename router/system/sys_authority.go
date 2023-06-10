package system

import (
	"github.com/gin-gonic/gin"
	api "github.com/zhou-Qingzhang/gin-admin/api"
	"github.com/zhou-Qingzhang/gin-admin/middleware"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityApi := api.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority) // 创建角色
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority) // 删除角色
		authorityRouter.PUT("updateAuthority", authorityApi.UpdateAuthority)  // 更新角色
	}
	{
		authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}
}
