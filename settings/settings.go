package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Init 将配置信息加载到 viper 的对象中
func Init() (err error) {
	viper.SetConfigFile("./config.yaml")
	if err = viper.ReadInConfig(); err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed,err: %v \n", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("setting file changed.")
	})
	return err
}
