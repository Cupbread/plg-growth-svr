package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var CfgMysql *viper.Viper
var CfgLog *viper.Viper

// Setup 载入配置文件
func Setup(path string) {
	viper.SetConfigFile(path)
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	CfgMysql = viper.Sub("mysql")
	if CfgMysql == nil {
		panic("config not found mysql.database")
	}
}

func SetApplicationIsInit() {
	SetConfig("./config", "settings.application.isInit", false)
}

func SetConfig(configPath string, key string, value interface{}) {
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	viper.WriteConfig()
}
