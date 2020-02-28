package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/spf13/viper"
	"github.com/xxgail/sql/mysqlconn"
	"github.com/xxgail/sql/redisconn"
	"strconv"
)

func main() {
	//initConfig()
	//
	//initRedis()
	//
	//initMysql()

	fmt.Println("----------------------------------------------------------------")
	var i int = 0
	num := strconv.Itoa(i)
	u := md5.New()
	u.Write([]byte("UserDBUpdate" + num))
	res := hex.EncodeToString(u.Sum(nil))
	u1 := md5.New()
	u1.Write([]byte("UserDBUpdate0"))
	res1 := hex.EncodeToString(u1.Sum(nil))
	fmt.Println(num, res, res1)
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
