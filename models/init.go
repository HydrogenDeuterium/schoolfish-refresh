package models

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"os"
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

func DBInit() DBs {
	return DBs{
		redisInit(),
		mysqlInit(),
	}
}
