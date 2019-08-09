package controllers

import "github.com/kataras/iris"
import "fmt"

type MyRes struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func WrapRes(result interface{}, status int, msg string) *MyRes {
	res := &MyRes{Status: status, Result: result, Msg: msg}
	return res
}
func writeLog(ctx iris.Context, result interface{}) {
	fmt.Println(ctx.Params())
	fmt.Println(result)
}
