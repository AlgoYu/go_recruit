package common

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

var Gedis redis.Conn

func ConnectRedis()  {
	var err error
	Gedis, err = redis.Dial("tcp", beego.AppConfig.String("redis_connect"))
	if err!=nil{
		panic(err)
	}
}

func Put(key string,value interface{})  {
	_, err := Gedis.Do("SET", key, value)
	if err!=nil{
		beego.Error("Redis set error:"+err.Error())
	}
}

func HashPut(key,filed string,value interface{})  {
	_, err := Gedis.Do("HSET", key, filed, value)
	if err!=nil{
		beego.Error("Redis set error:"+err.Error())
	}
}

func Expire(key string,timeValue int)  {
	_, err := Gedis.Do("EXPIRE", key, timeValue)
	if err!=nil{
		beego.Error(err)
	}
}

func PutTimeLimit(key,value string,timeValue int)  {
	_, err := Gedis.Do("SET", key, value, "EX", timeValue)
	if err!=nil{
		beego.Error("Redis set error:"+err.Error())
	}
}

func GetString(key string) string {
	s, err := redis.String(Gedis.Do("GET", key))
	if err!=nil{
		beego.Error("Redis set error:"+err.Error())
		return ""
	}
	return s
}

func CloseRedis()  {
	Gedis.Close()
}