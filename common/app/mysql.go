package app

import (
	"common/tool"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func MysqlInit() {
	// 日志打印
	newLogger := logger.Default
	tool.MyLog.Println("Mysql:初始化！")
	v, err := gorm.Open(mysql.Open(Config.MysqlUrl), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 定义表前缀
			SingularTable: true, // true不在表后面+ s，
		},
	})
	tool.AssertErrWithErrInfo(err, "Mysql:初始化失败!")
	Db = v
}
