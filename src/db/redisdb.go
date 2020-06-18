package db

import (
	"github.com/go-redis/redis"
)

type RedisDB struct {
	_init     bool
	_instance *redis.Client
}

var redisinst *RedisDB

//返回redis
func (this *RedisDB) GetDB() (*redis.Client, error) {
	if this._init == false || this._instance == nil {
		ConnectDB(this)
	}
	if _, err := this._instance.Ping().Result(); err != nil {
		return nil, err
	}
	return this._instance, nil
}

func ConnectDB(db *RedisDB) {
	db._instance = redis.NewClient(&redis.Options{
		Addr:     "106.54.46.165:6379",
		Password: "Xiaonan7147", // no password set
		DB:       0,             //
	})
	db._init = true
}

func GetRedisManager() *RedisDB {
	if redisinst == nil {
		redisinst = &RedisDB{
			_init:     false,
			_instance: nil,
		}
		ConnectDB(redisinst)
	}
	return redisinst
}
