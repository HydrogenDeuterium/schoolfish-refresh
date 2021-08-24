package models

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"math/rand"
	"os"
	"time"
)

func redisInit() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_address"),  // use default Addr
		Password: os.Getenv("redis_password"), // no password set
		DB:       0,                           // use default DB
	})
	return rdb
}

func mysqlInit() *gorm.DB {
	db, _ := gorm.Open("mysql",
		os.Getenv("mysql_jdbc"))
	return db
}

type DBs struct {
	redis *redis.Client
	mysql *gorm.DB
}

func (dbs DBs) Close() {
	_ = dbs.redis.Close()
	_ = dbs.mysql.Close()
}

func (dbs DBs) RedisGet(email string) string {
	ret, _ := dbs.redis.Get(email).Result()
	return ret
}

func (dbs DBs) RedisSet(k string) error {
	v := getRandomString()
	return dbs.redis.Set(k, v, 5*time.Minute).Err()

}

func getRandomString() interface{} {
	//六位随机数字
	randBytes := make([]byte, 3)
	_, _ = rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func DBInit() DBs {
	return DBs{
		redisInit(),
		mysqlInit(),
	}
}
