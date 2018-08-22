package bootstrap

import (
    "github.com/gin-gonic/gin"
    "wego3/app/routers"
)

// use
var ApiRouter    = new(routers.ApiRouter)
var FrontRouter  = new(routers.FrontRouter)
var CommonRouter = new(routers.CommonRouter)

// 路由整个应用
func Route(app *gin.Engine) {
    // 注册路由
    ApiRouter.Route(app)
    FrontRouter.Route(app)
    CommonRouter.Route(app)
}
