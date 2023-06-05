package system

import (
	"github.com/zhou-Qingzhang/gin-admin/global"
	"github.com/zhou-Qingzhang/gin-admin/model/common/response"
	"github.com/zhou-Qingzhang/gin-admin/model/system"
	systemRes "github.com/zhou-Qingzhang/gin-admin/model/system/response"
	"github.com/zhou-Qingzhang/gin-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct{}

// GetMenu
// @Tags      AuthorityMenu
// @Summary   获取用户动态路由
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      request.Empty                                                  true  "空"
// @Success   200   {object}  response.Response{data=systemRes.SysMenusResponse,msg=string}  "获取用户动态路由,返回包括系统菜单详情列表"
// @Router    /menu/getMenu [post]
func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	if menus == nil {
		menus = []system.SysMenu{}
	}
	response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
}
