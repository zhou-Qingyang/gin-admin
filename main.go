package main

import (
	"github.com/zhou-Qingzhang/gin-admin/core"
	"github.com/zhou-Qingzhang/gin-admin/global"
	"github.com/zhou-Qingzhang/gin-admin/initialize"
	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	// initialize.OtherInit() 		//JWT 配置 相关

	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	// initialize.Timer()                // 初始化 定时器
	// initialize.DBList()   //数据库初始化
	// if global.GVA_DB != nil {
	// 	initialize.RegisterTables(global.GVA_DB) // 初始化表
	// 	// 程序结束前关闭数据库链接
	// 	db, _ := global.GVA_DB.DB()
	// 	defer db.Close()
	// }					// 全局 数据库的初始化表
	core.RunWindowsServer()
}
