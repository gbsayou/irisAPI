package middleware

import (
	"../services"
	"crypto/md5"
	"encoding/hex"
	"github.com/kataras/iris"
)

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Auth(ctx iris.Context) {
	// checkToken(ctx)
	token := ctx.URLParam("token")
	time := ctx.URLParam("time")
	originSignature := ctx.URLParam("signature")
	preMD5 := "token=" + token + "&time=" + string(time)
	signature := md5V(preMD5)
	if originSignature != signature {
		ctx.JSON("身份校验失败")
		return
	}
	ctx.Next()
}
func CheckToken(ctx iris.Context) {
	token := ctx.URLParam("token")
	user := services.GetUserByToken(token)
	if user == nil {
		ctx.JSON("token不合法")
		return
	}
	ctx.Values().Set("user", user)
	ctx.Next()
}
