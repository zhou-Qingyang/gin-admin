package system

import "github.com/zhou-Qingzhang/gin-admin/global"

type SysAuthorityBtn struct {
	global.GVA_MODEL
	AuthorityId      uint `gorm:"comment:角色ID"`
	SysMenuID        uint `gorm:"comment:菜单ID"`
	SysBaseMenuBtnID uint `gorm:"comment:菜单按钮ID"`
}
