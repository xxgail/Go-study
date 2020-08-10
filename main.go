package main

import (
	"fmt"
	"github.com/spf13/viper"
	"xxgail/sql/mysqlconn"
	"xxgail/sql/redisconn"
)

func main() {
	initConfig()

	//
	//initRedis()
	//
	//initMysql()

	fmt.Println("----------------------------------------------------------------")

	//router := gin.Default()
	//routers.Init(router)

	//task.Init()
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
