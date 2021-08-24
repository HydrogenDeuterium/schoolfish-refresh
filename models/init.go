package models

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

func RedisInit() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_address"),  // use default Addr
		Password: os.Getenv("redis_password"), // no password set
		DB:       0,                           // use default DB
	})
	return rdb
}

func MysqlInit() *gorm.DB {
	db, _ := gorm.Open("mysql",
		os.Getenv("mysql_jdbc"))
	return db
}

func RedisTest() {
	client := RedisInit()

	//过期时间：1 秒
	err := client.Set("key1", "random-value", 1*time.Second).Err()
	if err != nil {
		panic(err)
	}

	//1 秒之内，可以获取到
	val, err := client.Get("key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1=", val)

	//睡眠两秒，过期了 获取不到了。
	time.Sleep(2 * time.Second)
	val, err = client.Get("key1").Result()
	if err == redis.Nil {
		fmt.Println("key1 does not exist now")
	} else if err != nil {
		panic(err)
	}
	fmt.Println("key1", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist even")
	} else if err != nil {
		panic(err)
	}
	fmt.Println("key2=", val2)

	// Output: key value
	// key2 does not exist
}
