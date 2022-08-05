package db

import (
	"fmt"
	"github.com/nfjBill/gorm-driver-dm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-docker/config"
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
	db, err := gorm.Open(dm.Open(dbConfig.GetString("dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "DE_", //看了源码了，dm驱动初始化默认设成空了。大坑，这在驱动之前改的没用
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	//db.Config.NamingStrategy = schema.NamingStrategy{
	//	TablePrefix: "DE_",
	//}
	return db
}
