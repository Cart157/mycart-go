package routers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    // "wego3/app/controllers/api"
)

// class
type ApiRouter string

func (rtr *ApiRouter) Route(Router *gin.Engine) {
    api := Router.Group("/api")
    {
        api.GET("/test", func(c *gin.Context) {
            c.String(http.StatusOK, "测试api组路由")
        })
    }
}
