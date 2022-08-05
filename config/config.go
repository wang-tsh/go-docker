package config

import (
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
)

const (
	DbFileName  = "db"
	FileExtName = "yml"
	AppFileName = "application"
)

type APPCONFIG struct {
	LogPath      string
	LogFileName  string
	MaxAge       time.Duration
	RotationTime time.Duration
	ContextPath  string
}

var (
	DbConfig  *viper.Viper
	AppConfig *APPCONFIG
)

func init() {
	DbConfig = initDBConfig()
	AppConfig = initAppConfig()
}
func initDBConfig() *viper.Viper {
	workDir, _ := os.Getwd()
	appViper := viper.New()
	appViper.SetConfigName(DbFileName)
	appViper.SetConfigType(FileExtName)
	appViper.AddConfigPath(path.Join(workDir, "config"))
	err := appViper.ReadInConfig()
	if err != nil {
		panic("Failed to get config！")
	}
	return appViper
}
func initAppConfig() *APPCONFIG {
	workDir, _ := os.Getwd()
	appViper := viper.New()
	appViper.SetConfigName(AppFileName)
	appViper.SetConfigType(FileExtName)
	appViper.AddConfigPath(path.Join(workDir, "config"))
	err := appViper.ReadInConfig()
	if err != nil {
		panic("Failed to get config！")
	}
	appConfig := APPCONFIG{
		LogPath:      appViper.GetString("log.logPath"),
		LogFileName:  appViper.GetString("log.LogFileName"),
		MaxAge:       appViper.GetDuration("log.maxAge"),
		RotationTime: appViper.GetDuration("log.rotationTime"),
		ContextPath:  appViper.GetString("app.contextPath"),
	}
	return &appConfig
}
