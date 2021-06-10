package rd

import (
	"time"

	"github.com/go-redis/redis"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

// RdSet: store 1 data to redis
func RdSet(key, value string, exp time.Duration) (err error) {

	err = redisClient.Set(key, value, exp).Err()
	if err != nil {
		return err
	}

	return nil
}

// RdGet: get 1 data from redis
func RdGet(key string) (res string, err error) {

	res, err = redisClient.Get(key).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

// RdDel: delete 1 data from redis
func RdDel(key string) (err error) {

	_, err = redisClient.Del(key).Result()
	if err != nil {
		return err
	}

	return nil
}
