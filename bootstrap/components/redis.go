package components

import (
    "fmt"
    "sync"
    "github.com/go-redis/redis"
)

var Redis *redis.Client

func init() {
    var once sync.Once

    // 初始化Redis客户端
    once.Do(func() {
        Redis = redis.NewClient(&redis.Options{
            Addr:     Config.Get("redis.default.host").String(),
            Password: Config.Get("redis.default.password").String(),
            DB:       int(Config.Get("redis.default.database").Int()),
        })
    })

    pong, err := Redis.Ping().Result()
    fmt.Println(pong, err)
}

// 说明：这个redis库，用redis命令当方法名，用着舒服
// 每步操作后，都可以加.Result()来获得操作的结果
// Set Del 这种不关心结果的可以不加.Result()
