package database

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"
import "../config"
import "github.com/pelletier/go-toml"

var (
	DB = New()
)

func New() *gorm.DB {
	config := config.GetConfig
	mysqlDriver := config.Get("database.driver").(string)
	mysqlConfig := config.Get(mysqlDriver).(*toml.Tree)
	user := mysqlConfig.Get("MYSQL_USER").(string)
	password := mysqlConfig.Get("MYSQL_PASSWORD").(string)
	database := mysqlConfig.Get("MYSQL_DATABASE").(string)
	// DB, err := gorm.Open("mysql", "root:gbsayou@/iris?charset=utf8&parseTime=True&loc=Local")
	DB, err := gorm.Open(mysqlDriver, user+":"+password+"@/"+database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("数据库连接错误！")
	}
	// defer DB.Close()
	return DB
}
