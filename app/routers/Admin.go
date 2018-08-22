package routers

// import (
//   c "wego2/app/controllers"
//     "github.com/devfeel/dotweb"
//     "github.com/devfeel/middleware/cors"
// )

// // class & construct
// type AdminRouter string
// func NewAdminRouter() AdminRouter {
//     return "AdminRouter"
// }

// var GoodsController = c.NewGoodsController()
// var UserController = c.NewUserController()

// func (r AdminRouter) Register(Route *dotweb.HttpServer) {
//     // 设置cors选项中间件, 并使用默认的跨域配置
//     option := cors.NewConfig().UseDefault()

//     // 后台管理
//     admin := Route.Group("/admin").Use(cors.Middleware(option))

//     admin.GET("/goods", GoodsController.Index)
//     admin.GET("/goods/:id", GoodsController.Show)
//     admin.POST("/goods", GoodsController.Store)
//     admin.PUT("/goods/:id", GoodsController.Update)
//     admin.DELETE("/goods/:id", GoodsController.Destroy)
// }
