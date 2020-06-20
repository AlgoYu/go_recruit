package common

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

var Gedis redis.Conn

func ConnectRedis() {
	var err error
	Gedis, err = redis.Dial("tcp", beego.AppConfig.String("redis_connect"))
	if err != nil {
		panic(err)
	}
}

func RedisSet(key string, value interface{}) {
	_, err := Gedis.Do("SET", key, value)
	if err != nil {
		beego.Error("Redis set error:" + err.Error())
	}
}

func RedisHSet(key, filed string, value interface{}) {
	_, err := Gedis.Do("HSET", key, filed, value)
	if err != nil {
		beego.Error("Redis set error:" + err.Error())
	}
}

func RedisExists(key string) bool {
	exists, _ := redis.Bool(Gedis.Do("EXISTS", key))
	return exists
}

func RedisExpire(key string, timeValue int) {
	_, err := Gedis.Do("EXPIRE", key, timeValue)
	if err != nil {
		beego.Error(err)
	}
}

func RedisDelete(key string) {
	_, err := Gedis.Do("DEL", key)
	if err != nil {
		beego.Error(err)
	}
}

func RedisPutTimeLimit(key, value string, timeValue int) {
	_, err := Gedis.Do("SET", key, value, "EX", timeValue)
	if err != nil {
		beego.Error("Redis set error:" + err.Error())
	}
}

func RedisGetString(key string) string {
	s, err := redis.String(Gedis.Do("GET", key))
	if err != nil {
		beego.Error("Redis set error:" + err.Error())
		return ""
	}
	return s
}

func CloseRedis() {
	Gedis.Close()
}
