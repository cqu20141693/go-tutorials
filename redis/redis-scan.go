package redis

import (
	"context"
	"github.com/cqu20141693/go-service-common/redis"
)

func Scan(prefix string) (allKeys []string) {
	background := context.Background()
	iter := redis.RedisDB.Scan(background, 0, prefix, 10).Iterator()
	for iter.Next(background) {
		allKeys = append(allKeys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return
}

func ScanAllKey(prefix string) (allKeys []string) {
	background := context.Background()
	var cursor uint64
	for {
		var keys []string
		var err error

		keys, cursor, err = redis.RedisDB.Scan(background, cursor, prefix, 10).Result()
		if err != nil {
			panic(err)
		}
		allKeys = append(allKeys, keys...)
		if cursor == 0 {
			break
		}
	}
	return
}
