package system

import (
	"github.com/gin-gonic/gin"
	api "github.com/zhou-Qingzhang/gin-admin/api"
	"github.com/zhou-Qingzhang/gin-admin/middleware"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouterWithoutRecord := Router.Group("menu")
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	authorityMenuApi := api.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("addBaseMenu", authorityMenuApi.AddBaseMenu)           // 新增菜单
		menuRouter.POST("deleteBaseMenu", authorityMenuApi.DeleteBaseMenu)     // 删除菜单
		menuRouter.POST("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)     // 更新菜单
		menuRouter.POST("addMenuAuthority", authorityMenuApi.AddMenuAuthority) //	角色管理 增加menu和角色关联关系
	}
	{
		menuRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu)                   // 动态路由
		menuRouterWithoutRecord.POST("getMenuList", authorityMenuApi.GetMenuList)           // 菜单管理 所有菜单列表
		menuRouterWithoutRecord.POST("getBaseMenuById", authorityMenuApi.GetBaseMenuById)   // 根据id获取菜单、
		menuRouterWithoutRecord.POST("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)   //角色管理 获取所有的信息
		menuRouterWithoutRecord.POST("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // 根据改角色获取 含有的惨淡
	}
	return menuRouter
}
