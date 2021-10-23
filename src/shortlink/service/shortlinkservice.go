package main

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jxskiss/base62"
	redisInfo "golang-note/src/shortlink/redis"
	"io"
	"time"
)

type ShortLinkService struct {
	Cli *redis.Client
}

func main() {
	connect := redisInfo.DoRedisConnect("localhost:6379")
	service := ShortLinkService{Cli: connect}
	shorten, err := service.Shorten("www.baidu.com", 120)
	if err != nil {
		return
	}
	fmt.Printf(shorten)
	fmt.Println(service.UnShorten(shorten))
	fmt.Println(service.ShortLinkInfo(shorten))
}

var ctx = context.Background()

func (service *ShortLinkService) Shorten(url string, exp int64) (string, error) {

	h := toSha1(url)

	res, err := service.Cli.Get(ctx, fmt.Sprintf(redisInfo.URLHashKey, h)).Result()

	if err == redis.Nil {
		return "", err
	} else if err != nil {
		return "", err
	} else {
		if res != "" {
			return res, nil
		}
	}

	incr := service.Cli.Incr(ctx, redisInfo.URLIDKEY)
	if incr.Err() != nil {
		return "", incr.Err()
	}

	key := service.Cli.Get(ctx, redisInfo.URLIDKEY)
	if key.Err() != nil {
		return "", key.Err()
	}
	var data = []byte(key.String())
	eid := string(base62.Encode(data))

	setRes := service.Cli.Set(ctx, fmt.Sprintf(redisInfo.ShortLinkKey, eid), url, 0)
	if setRes.Err() != nil {
		return "", setRes.Err()
	}

	setRes = service.Cli.Set(ctx, fmt.Sprintf(redisInfo.URLHashKey, h), eid, 0)
	if setRes.Err() != nil {
		return "", setRes.Err()
	}

	detail, err := json.Marshal(
		&redisInfo.URLDetail{
			URL:                 url,
			CreatedAt:           time.Now().String(),
			ExpirationInMinutes: time.Duration(exp),
		},
	)
	if err != nil {
		return "", err
	}
	setRes = service.Cli.Set(ctx, fmt.Sprintf(redisInfo.ShortLinkDetailKey, eid), detail, 0)
	if setRes.Err() != nil {
		return "", setRes.Err()
	}

	return eid, nil
}

func toSha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func (service *ShortLinkService) ShortLinkInfo(eid string) (interface{}, error) {
	result, err := service.Cli.Get(ctx, fmt.Sprintf(redisInfo.ShortLinkDetailKey, eid)).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (service *ShortLinkService) UnShorten(eid string) (string, error) {
	result, err := service.Cli.Get(ctx, fmt.Sprintf(redisInfo.ShortLinkKey, eid)).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "nil", err
	} else {
		return result, nil
	}
}

//type storage interface {
//	Shorten(url string, exp int64) (string, error)
//	ShortLinkInfo(eid string) (interface{}, error)
//	UnShorten(eid string) (string, error)
//}
