package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

const CONFIG_TYPE = "yaml"

func ViperConfig(path, name string) {
	viper.SetConfigName(name)        // name of config file (without extension)
	viper.SetConfigType(CONFIG_TYPE) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)        // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("配置文件出现错误: %w", err))
	}
}
