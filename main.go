package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xxgail/sql/mysqlconn"
	"github.com/xxgail/sql/redisconn"
)

func main() {
	a := 5
	//b := 100
	fmt.Println(float64(a) * 0.01)
	//initConfig()
	//
	//initRedis()
	//
	//initMysql()
}

// 初始化redis配置，启动redis
func initRedis() {
	redisconn.InitNewClient()
}

// 初始化MySQL配置，启动MySQL
func initMysql() {
	mysqlconn.InitDB()
}

func initConfig() {
	viper.SetConfigName("config/sql")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
