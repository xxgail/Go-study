package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/techoner/gophp/serialize"
	"github.com/xxgail/sql/mysqlconn"
	"github.com/xxgail/sql/redisconn"
	"time"
)

func main() {
	initConfig()

	initRedis()

	initMysql()
	redisClient := redisconn.GetClient()
	var bidUserId int64 = 12
	var bidPrice int64 = 1000
	bidLog := make(map[string]interface{})
	bidLog["bid_user_id"] = bidUserId
	bidLog["bid_price"] = bidPrice
	bidLog["bid_before"] = 12222
	bidLog["bid_after"] = 11222
	bidLog["bid_time"] = time.Now().Format("2006-01-02 15:04:05")
	bidLog["step"] = 100

	//bidLog["last_user_id"] = lastUserId
	//bidLog["last_mar_before"] = lastUserMarBefore
	//bidLog["last_mar_after"] = lastUserMarAfter
	//bidLog["last_com_before"] = lastUserComBefore
	//bidLog["last_com_after"] = lastUserComAfter
	//
	//bidLog["invite_user_id"] = lastInviteUserId
	//bidLog["invite_before"] = lastInviteRewBefore
	//bidLog["invite_after"] = lastInviteRewAfter

	info, _ := serialize.Marshal(bidLog)
	redisClient.RPush("aa", info)
	//userIds,err := redisClient.SMembers("MBUser:UserDBUpdate:362104ec9caf13ade9754ccc918a3229").Result()
	////fmt.Println(val , val,reflect.TypeOf(val))
	//
	////userIds, err := redisClient.SMembers(keyUserDB).Result()
	//if err != nil {
	//	fmt.Println("获取需要更新的用户信息失败",err)
	//}
	//if len(userIds) == 0 {
	//	fmt.Println("没有需要更新的用户信息")
	//}else{
	//	for _,v := range userIds{
	//		fmt.Println("----- 当前更新的用户ID为 ",v)
	//		userId,_ := strconv.Atoi(v)
	//		//user := mbuser.InitUser(userId,false)
	//		//userInfo := user.GetUserInfo()
	//
	//		//开启事务
	//		tx, err := mysqlconn.GetMysqlConn().Begin()
	//		if err != nil{
	//			fmt.Println("tx fail")
	//		}
	//		//准备sql语句
	//		stmt, err := tx.Prepare("UPDATE users SET balance = ?, income = ?, bid_count = ? WHERE id = ?")
	//		if err != nil{
	//			fmt.Println("Prepare fail")
	//		}
	//		//设置参数以及执行sql语句
	//		res, err := stmt.Exec(1001,1000, 10, userId)
	//		if err != nil{
	//			fmt.Println("Exec fail")
	//		}
	//		//提交事务
	//		tx.Commit()
	//		fmt.Println(res.RowsAffected())
	//	}
	//}

	//----- 插入bid_log
	val, _ := redisClient.LPop("aa").Result()
	log, _ := serialize.UnMarshal([]byte(val))
	bidLog1 := log.(map[string]interface{})
	tx, err := mysqlconn.GetMysqlConn().Begin()
	queryBid := "INSERT INTO auction_bid_log (`auction_id`,`user_id`,`bid`,`commission`,`group_id`,`created_at`,`updated_at`,`margin`) VALUES (?,?,?,?,?,?,?,?)"
	stemBid, err := tx.Prepare(queryBid)
	if err != nil {
		fmt.Println("Prepare fail")
	}
	resBid, err := stemBid.Exec(1, bidLog1["bid_user_id"], bidLog1["bid_price"], 110, 1, "2020-01-02 14:22:22", 1578385660, 1)
	if err != nil {
		fmt.Println("Exec fail")
	}

	queryBid1 := "INSERT INTO auction_bid_log (`auction_id`,`user_id`,`bid`,`commission`,`group_id`,`created_at`,`updated_at`,`margin`) VALUES (?,?,?,?,?,?,?,?)"
	stemBid1, err := tx.Prepare(queryBid1)
	if err != nil {
		fmt.Println("Prepare fail")
	}
	resBid1, err := stemBid1.Exec(21, 21, 100, 110, 1, "2020-01-02 14:22:22", 1578385660, 1)
	if err != nil {
		fmt.Println("Exec fail")
	}

	//将事务提交
	tx.Commit()
	fmt.Println(resBid.LastInsertId())
	fmt.Println(resBid1.LastInsertId())
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
