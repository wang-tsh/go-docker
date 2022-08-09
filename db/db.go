package db

import (
	"fmt"
	"github.com/nfjBill/gorm-driver-dm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-docker/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	DB = initDB(config.DbConfig)

}
func initDB(dbConfig *viper.Viper) *gorm.DB {
	//dbConfig
	log.Infoln("to init db")
	var db *gorm.DB
	var err error
	if "mysql" == config.DbConfig.GetString("type") {
		db, err = gorm.Open(mysql.Open(dbConfig.GetString("dsn")), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "DE_", //看了源码了，dm驱动初始化默认设成空了。大坑，这在驱动之前改的没用
			},
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else if "dm" == config.DbConfig.GetString("type") {
		db, err = gorm.Open(dm.Open(dbConfig.GetString("dsn")), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "DE_", //看了源码了，dm驱动初始化默认设成空了。大坑，这在驱动之前改的没用
			},
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		log.Errorf("error db type")
	}

	if err != nil {
		fmt.Println(err)
	}
	//db.Config.NamingStrategy = schema.NamingStrategy{
	//	TablePrefix: "DE_",
	//}
	return db
}
