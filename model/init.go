package model

import (
	"fmt"
	"gin-api-boilerplate/config"
	"gin-api-boilerplate/pkg/log"
	"go.uber.org/zap"

	// MySQL driver.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Master *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Error("Database connection failed. Database name: "+name, zap.Error(err))
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(config.C.DbConfig.Gorm.LogMode)
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

// used for cli
func InitSelfDB() *gorm.DB {
	return openDB(config.C.DbConfig.Username,
		config.C.DbConfig.Password,
		config.C.DbConfig.Host,
		config.C.DbConfig.Database)
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func (db *Database) Init() {
	DB = &Database{
		Master: GetSelfDB(),
	}
}

func (db *Database) Close() {
	DB.Master.Close()
}
