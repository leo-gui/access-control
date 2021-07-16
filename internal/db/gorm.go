package db

import (
	"access-control/internal"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wonderivan/logger"
)


func GormCoon() (db *gorm.DB) {
	m:= internal.InitConfigByViper()
	dsn := m.Username + ":" + m.Password + "@(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	if db, err := gorm.Open("mysql",dsn); err != nil {
		logger.Alert("数据库连接失败")
		return nil
	} else {
		sqlDB := db.DB()
		if err := sqlDB.Ping(); err == nil {
			sqlDB.SetMaxIdleConns(m.MaxIdleConns)
			sqlDB.SetMaxOpenConns(m.MaxOpenConns)
			return db
		}else {
			return nil
		}
	}
}

