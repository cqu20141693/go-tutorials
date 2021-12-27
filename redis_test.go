package go_tutorials

import (
	"context"
	"fmt"
	"github.com/cqu20141693/go-service-common/boot"
	ccredis "github.com/cqu20141693/go-service-common/redis"
	credis "github.com/cqu20141693/go-tutorials/redis"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"testing"
	"time"
)

func init() {
	fmt.Println("start test")
	boot.Task()
}

/**
scan ： 时间复杂度O（1）
每次扫描需要执行cursor 和count
cursor 为0表示首次扫描，count 表示一次扫描多少底层数组槽（redis key底层是一个Map,底层是数组槽（相同hash）+链表）
所以一次命令可能返回0个数据，也可能多余count个数据
当返回的cursor=0表示遍历完成

缺点，可能会返回重复的key
*/
func TestScan(t *testing.T) {

	prefix := "sips:*"

	ctx := context.Background()
	_, err := ccredis.RedisDB.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 100; i++ {
			pipe.Set(ctx, "sips:"+strconv.Itoa(i), "test", time.Second*30)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	allKey := credis.ScanAllKey(prefix)
	fmt.Println("ScanAllKey size=", len(allKey))
	for i, s := range allKey {
		fmt.Printf("index=%d value=%s \n", i, s)
	}
	fmt.Println("scan size=", len(allKey))
	keys := credis.Scan(prefix)
	for i, key := range keys {
		fmt.Printf("index=%d key=%s \n", i, key)
	}
	_, err = ccredis.RedisDB.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 100; i++ {
			pipe.Del(ctx, "sips:"+strconv.Itoa(i))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

}

func TestHash(t *testing.T) {
	key := "sip:test"
	field := "camera1"
	formatInt := strconv.FormatInt(time.Now().UnixMilli(), 10)
	hSet := ccredis.RedisDB.HSet(context.Background(), key, field, formatInt)
	if hSet.Err() != nil {
		log.Println("hSet err", hSet.Err())
	}
	data, err := ccredis.RedisDB.HGet(context.Background(), key, field).Result()
	if err != nil {
		log.Println("hGet err", err)
	}
	parseInt, _ := strconv.ParseInt(data, 10, 64)
	duration := time.Duration(parseInt)
	t1 := time.Now().UnixMilli() - parseInt
	log.Println("hGet data=", data, duration, t1)

}
