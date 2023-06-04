package internal

import (
	"log"
	"os"
	"time"

	"github.com/zhou-Qingzhang/gin-admin/global"
	"gorm.io/gorm/schema"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBBASE interface {
	GetLogMode() string
}

// 定义全局变量
var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	// 创建默认日志记录器
	defaultLogger := logger.Default

	var logMode DBBASE
	logMode = &global.GVA_CONFIG.Mysql

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		config.Logger = defaultLogger.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = defaultLogger.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = defaultLogger.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = defaultLogger.LogMode(logger.Info)
	default:
		config.Logger = defaultLogger.LogMode(logger.Info)
	}

	// 创建新的日志记录器并输出到控制台
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	config.Logger = newLogger

	return config
}
