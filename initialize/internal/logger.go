package internal

import (
	"fmt"
	"io"
	"log"

	"github.com/zhou-Qingzhang/gin-admin/global"
)

// type writer struct {
// 	logger.Writer
// }

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
// func NewWriter(w logger.Writer) *writer {
// 	return &writer{Writer: w}
// }

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
// func (w *writer) Printf(message string, data ...interface{}) {
// var logZap bool
// switch global.GVA_CONFIG.System.DbType {
// case "mysql":
// 	logZap = global.GVA_CONFIG.Mysql.LogZap
// case "pgsql":
// 	logZap = global.GVA_CONFIG.Pgsql.LogZap
// }
// 	var logZap bool
// 	logZap = global.GVA_CONFIG.Mysql.LogZap
// 	if logZap {
// 		global.GVA_LOG.Info(fmt.Sprintf(message+"\n", data...))
// 	} else {
// 		w.Writer.Printf(message, data...)
// 	}
// }

type MyWriter struct {
	io.Writer // 继承 io.Writer 接口
}

func (w *MyWriter) Write(p []byte) (int, error) {
	log.Printf("%s", p) // 使用 log 包输出日志信息
	return len(p), nil
}

func (w *MyWriter) Printf(message string, data ...interface{}) {
	var logZap bool
	logZap = global.GVA_CONFIG.Mysql.LogZap
	if logZap {
		global.GVA_LOG.Info(fmt.Sprintf(message+"\n", data...))
	}

}
