package initialize

import (
	"github.com/zhou-Qingzhang/gin-admin/config"
	"github.com/zhou-Qingzhang/gin-admin/global"

	"github.com/zhou-Qingzhang/gin-admin/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库

func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn1(), // DSN data source name
		DefaultStringSize:         191,      // string 类型字段的默认长度
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		global.GVA_LOG.Info("mysql has initialized")

		return db
	}
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置   //这个不知道 干嘛的
func GormMysqlByConfig(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
