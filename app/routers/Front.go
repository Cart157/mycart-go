package routers

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    . "wego3/bootstrap/components"
  // c "wego3/app/controllers"
)

// class
type FrontRouter string

func (rtr *FrontRouter) Route(Router *gin.Engine) {
    Router.GET("/", func(c *gin.Context) {
        fmt.Printf("%T\n", Config.Get("haha").String())
        fmt.Println(Config.Get("haha").String())
        fmt.Println(Config.Get("haha").String() == "")
        fmt.Println(Config.GetArray("nihao"))

        fmt.Println(Config.Get("redis.default.host"))
        // Redis.Set("haha", "大爷的", 0).Result()
        // fmt.Println("start……, port:8080, visit:http://localhost:8080")

        // ret, _ := Redis.Get("haha").Result()
        // fmt.Println(ret)
        c.String(http.StatusOK, "Hello World!")
    })

    Router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
}
