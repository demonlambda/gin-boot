package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//根据启动参数中指定的环境信息，加载相应环境下的配置文件
func init() {
	fmt.Printf("run args:%v\n", os.Args)

	//加载配置文件
	env := flag.String("env", "", "配置文件")
	flag.Parse()

	cwd, _ := os.Getwd()
	configFilePath := cwd + string(os.PathSeparator) + "config" + string(os.PathSeparator) + "config-" + *env + ".yaml"
	commonConfigFilePath := cwd + string(os.PathSeparator) + "config" + string(os.PathSeparator) + "config.yaml"

	//加载默认配置
	setDefualtConfig()
	if isValidEnv(*env) {
		viper.SetConfigFile(commonConfigFilePath)
		viper.ReadInConfig()

		viper.SetConfigFile(configFilePath)
		viper.MergeInConfig()
	} else {
		panic("empty or valid  env:" + *env)
	}
}

func isValidEnv(env string) bool {
	validEnvs := map[string]int{"dev": 1, "test": 1, "prd": 1}
	if _, ok := validEnvs[env]; ok {
		return true
	}

	return false
}

func setDefualtConfig() {
	viper.SetDefault("port", "8090")
	viper.SetDefault("mysql.maxIdleConns", 10)
	viper.SetDefault("mysql.maxOpenConns", 100)
}
