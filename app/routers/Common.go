package routers

import (
    "github.com/gin-gonic/gin"
    "wego3/app/controllers/common/netease"
)

// use
var ImController = new(netease.ImController)

// class
type CommonRouter string

// public
func (rtr *CommonRouter) Route(Router *gin.Engine) {
    Router.GET("/netease/im/notify",        ImController.Notify)
    Router.POST("/netease/im/notify",       ImController.NotifyPost)
}
