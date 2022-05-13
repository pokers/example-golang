package sql

import (
	"example-restapi+db/src/util"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Type     string
	Username string
	Pass     string
	URL      string
	Port     uint
	DBName   string
}

type MySqlInst struct {
	Sql      *gorm.DB
	isOpened bool
}

var inst MySqlInst

func InitDB(cfg *DBConfig) {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.Username, cfg.Pass, cfg.URL, cfg.Port, cfg.DBName)
	// util.PrintInfo("DSN : %s", dataSourceName)
	inst.Sql, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		util.PrintError("Failed to open DB. Should be reboot....")
		panic(err.Error())
		return
	}
	mysqlDB, err := inst.Sql.DB()
	if err != nil {
		util.PrintError("Failed to get msqlDB. Should be reboot....")
		panic(err.Error())
		return
	}
	mysqlDB.SetConnMaxLifetime(time.Hour)
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(100)
	inst.isOpened = true
}

func GetConnection() MySqlInst {
	if !inst.isOpened {
		panic("DB is not opened")
	}
	return inst
}

func CloseDB() {
	mysqlDB, err := inst.Sql.DB()
	if inst.isOpened && err != nil {
		mysqlDB.Close()
		inst.isOpened = false
	}
	util.PrintInfo("MySql DB : closed...")
}
