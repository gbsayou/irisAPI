package middleware

import "github.com/kataras/iris"
import "fmt"

func MyLogger(ctx iris.Context) {
	// result := ctx.Next()
	fmt.Println("before")
	ctx.Next()
	fmt.Println("after")
	fmt.Println(ctx.Params())
	// fmt.Println(ctx.Response)
}
