package controllers

import (
	"../services"
	"github.com/kataras/iris"
	"strconv"
)

func GetUser(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		ctx.JSON("传值不对")
		return
	}
	result, err := services.GetUserByID(id)
	if err != nil {
		// fmt.Println(err)
		ctx.JSON(WrapRes(nil, 200, "没有找到用户"))
		return
	}
	ctx.JSON(WrapRes(result, 200, "操作成功"))
}

func UserRegister(ctx iris.Context) {
	name := ctx.FormValue("name")
	age, _ := strconv.Atoi(ctx.FormValueDefault("age", "20"))
	height, _ := strconv.Atoi(ctx.FormValueDefault("height", "170"))
	if err := services.UserRegister(name, age, height); err != nil {
		panic("注册用户失败！")
	}
	ctx.JSON("用户注册成功！")
}

func GetUserByName(ctx iris.Context) {
	name := ctx.Params().Get("name")
	ctx.JSON(services.GetUserByName(name))
}
