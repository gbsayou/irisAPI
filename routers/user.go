package routers

import (
	"../controllers"
	"github.com/kataras/iris"
)

func UserRouter(app *iris.Application) {
	app.Get("/users/id/{id}", controllers.GetUser)
	app.Post("/users", controllers.UserRegister)
	app.Get("/users/name/{name}", controllers.GetUserByName)
}
