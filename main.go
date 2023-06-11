package main

import (
	"github.com/zhou-Qingzhang/gin-admin/core"
	"github.com/zhou-Qingzhang/gin-admin/global"
	"github.com/zhou-Qingzhang/gin-admin/initialize"
	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper

	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

	global.GVA_DB = initialize.Gorm() // gorm连接数据库
				// 全局 数据库的初始化表
	core.RunWindowsServer()
}
