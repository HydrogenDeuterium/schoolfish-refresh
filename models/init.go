package models

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
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
	db, err := gorm.Open("mysql",
		os.Getenv("mysql_jdbc"))
	if err != nil {
		fmt.Printf("%s", err)
	}
	return db
}

type DBGroup struct {
	redis *redis.Client
	Mysql *gorm.DB
}

func (dbs DBGroup) Close() {
	_ = dbs.redis.Close()
	_ = dbs.Mysql.Close()
}

func (dbs DBGroup) RedisGet(email string) (string, error) {
	ret, err := dbs.redis.Get(email).Result()
	return ret, err
}

func (dbs DBGroup) RedisSet(email string) string {
	code := "123456"

	//*测试用
	if os.Getenv("test") != "True" {
		code = getRandomString()
	}
	// 测试用*/

	err := dbs.redis.Set(email, code, 5*time.Minute).Err()
	if err != nil {
		return ""
	}
	return code

}

func (dbs DBGroup) RedisDel(email string) {
	dbs.redis.Del(email)
}

func getRandomString() string {
	//六位随机数字
	randBytes := make([]byte, 3)
	_, _ = rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func DBInit() DBGroup {
	return DBGroup{
		redisInit(),
		mysqlInit(),
	}
}
