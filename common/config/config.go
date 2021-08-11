package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	v *viper.Viper
)

func init() {
	// todo 后面可以支持 consul 或 ectd
	// 读取yaml文件
	v = viper.New()
	//v.Debug()

	configName := os.Getenv("CONFIG_NAME")
	if configName == "" { // 如果全局变量没有配置,使用默认的配置文件
		configName = "config" // 默认配置文件名称
	}

	// 设置读取的配置文件名
	v.SetConfigName(configName)

	v.AddConfigPath("./config") // windows环境下为%GOPATH，linux环境下为$GOPATH
	// 设置配置文件类型
	v.SetConfigType("yaml")

	// 加载配置文件
	if err := v.ReadInConfig();err != nil {
		fmt.Printf("config read err: %s\n", err)
	}
}

// 返回字符类型配置
func String(key string) string {
	return v.GetString(key)
}

// 返回Int类型配置
func Int(key string) int {
	return v.GetInt(key)
}

// 返回布尔类型配置
func Bool(key string) bool {
	return v.GetBool(key)
}

// 返回 map[string]interface{} 配置
func Map(key string) map[string]interface{} {
	return v.GetStringMap(key)
}

// 返回 字符型数组 配置
func StringArray(key string) []string {
	return v.GetStringSlice(key)
}

// 返回 整型数组 配置
func IntArray(key string) []int {
	return v.GetIntSlice(key)
}

// 返回 浮点型 配置
func Float64(key string) float64 {
	return v.GetFloat64(key)
}