package core

import (
	"fmt"
	"time"

	"github.com/zhou-Qingzhang/gin-admin/global"
	"github.com/zhou-Qingzhang/gin-admin/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	//初始化
	initialize.OtherInit()

	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	// 从db加载jwt数据
	// if global.GVA_DB != nil {
	// 	system.LoadAll()
	// }
	// Router.Static("/form-generator", "./resource/page")  生成代码的位置

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	欢迎使用 gin-admin
	当前版本:v2.5.5
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
