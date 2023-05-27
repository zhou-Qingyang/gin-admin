package system

import (
	"github.com/zhou-Qingzhang/gin-admin/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
