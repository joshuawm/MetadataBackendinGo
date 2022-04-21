package database

import (
	// "github.com/go-redis/redis"

	// "context"
	"errors"
	"log"

	"github.com/alovn/go-bloomfilter"

	"github.com/go-redis/redis/v8"
)

// //this is for upstash
// //about Redis Compatiablity, visit https://docs.upstash.com/redis/overall/rediscompatibility
// import(
// 	"context"
// 	"github.com/go-redis/redis/v8"

// )

// func ConnectRedis(url string) (bool,*redis.Client) {
// 	opt,_ := redis.ParseURL(url)
// 	client:=redis.NewClient(opt)
// 	var ok bool
// 	_, err := client.Ping()
// 	if(err!=nil){
// 		ok = false
// 	}else{
// 		ok =true
// 	}
// 	return ok,client
// }

var RedisClient *redis.Client

// var ct context.Context = context.Background()

func ConnectRedis(Addr string, Password string, DB int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       DB,
	})

	return rdb
}

type BFilter struct {
	BFClient bloomfilter.BloomFilter
}

func (this *BFilter) Create(Addr string, Password string, DB int, key string) {
	client := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       DB,
	})
	this.BFClient = bloomfilter.NewRedisBloomFilter(client, key, 1000000)
	log.Println("redis initializition finished")
}
func (this *BFilter) BloomFilterAdd(value string) (bool, error) {
	var e error
	if exist, _ := this.BFClient.MightContain([]byte(value)); exist {
		e = errors.New("该数据已存在")
	} else {
		e = this.BFClient.Put([]byte(value))
	}
	if e != nil {
		return false, e
	} else {
		return true, e
	}

}

func (this *BFilter) BloomFilterExist(value string) (bool, error) {
	return this.BFClient.MightContain([]byte(value))
}
