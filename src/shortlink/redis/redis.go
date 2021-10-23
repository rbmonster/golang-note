package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	URLIDKEY = "next.url.id"

	ShortLinkKey = "shortlink:%s:url"

	URLHashKey = "urlHash:%s:url"

	ShortLinkDetailKey = "shortLink:%s:detail"
)

type URLDetail struct {
	URL                 string        `json:"url"`
	CreatedAt           string        `json:"created_at"`
	ExpirationInMinutes time.Duration `json:"expiration_in_minutes
"`
}

func DoRedisConnect(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "121.5.73.118:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

//func main () {
//	con, err := redis.Dial("tcp", "121.5.73.118:6379")
//	if err != nil {
//		panic("")
//	}
//	err1 := con.Send("SET", "check", "test")
//	if err1 != nil {
//		fmt.Println(err1)
//		return
//	}
//	con.Flush()
//	con.Receive()
//}
