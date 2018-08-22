package main

import (
    "os"
    "io"
    "github.com/gin-gonic/gin"
    "wego3/bootstrap"
)

func main() {
    // Log文件，还不会用，可能只能手打
    gin.DisableConsoleColor()

    f, _ := os.Create("./storage/logs/gin.log")
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    app := gin.Default()
    app.Static("/assets", "./public/assets")
    app.Static("/uploads", "./public/uploads")
    app.StaticFile("/favicon.ico", "./public/favicon.ico")

    // 启动路由
    bootstrap.Route(app)

    // 走起
    app.Run() // listen and serve on 0.0.0.0:8080
}
