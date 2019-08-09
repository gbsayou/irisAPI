package main

import (
	"./middleware"
	"./routers"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	myHand := middleware.MyHand{}
	app.Router.MyHand = myHand
	app.Use(middleware.Auth)
	app.Use(middleware.CheckToken)
	// app.Use(middleware.MyLogger)
	routers.UserRouter(app)
	app.Run(iris.Addr(":8080")) // 跑起来
}
