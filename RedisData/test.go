package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: "",
		DB:       viper.GetInt("redis.DB"),
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	val, err := client.Exists(keyUserBalancer).Result()
	if val == 1 {
		var balance int64 = 500
		tp := false
		res := SetUserBalance(balance, tp)
		fmt.Println("res", res)
	} else {
		err = client.Set(keyUserBalancer, "0", 0).Err()
	}
	fmt.Println(val)

	//err = client.SAdd("key1","value11").Err()

	//err = client.Set("MBUser:Balancer:ECC903ECCC3D597F8971147BB01E7875", "0", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val,err := client.SMembers("key1").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key",val)
	//
	//val2,err := client.Get("key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exists")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2",val2)
	//}
}

//$balance = intval($balance);
//if($type){
//$after = Redis::incrby($this->keyUserBalancer,$balance); //递增
//$before = $after - $balance;
//}else{
//$after = Redis::decrby($this->keyUserBalancer,$balance); //递减
//$before = $after + $balance;
//}
//
////        Redis::set($this->keyUserDB,$this->user_id);
//Redis::sadd($this->keyUserDB,$this->user_id);
//
//return [
//'after'=>$after,
//'before'=>$before
//];

var keyUserBalancer = "MBUser:Balancer:ECC903ECCC3D597F8971147BB01E7875"

func SetUserBalance(balance int64, tp bool) map[string]int64 {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: "",
		DB:       viper.GetInt("redis.DB"),
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	val, err := client.Exists(keyUserBalancer).Result()
	fmt.Println(val)

	var (
		after  int64
		before int64
	)
	if tp == true {
		after, err = client.IncrBy(keyUserBalancer, balance).Result()
		if err != nil {
			panic(err)
		}
		before = after - balance
	} else {
		after, err = client.DecrBy(keyUserBalancer, balance).Result()
		if err != nil {
			panic(err)
		}
		before = after + balance
	}

	res := map[string]int64{
		"after":  after,
		"before": before,
	}
	return res
}
