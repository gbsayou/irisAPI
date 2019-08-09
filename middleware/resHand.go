package middleware

import (
	// "fmt"
	"../logger"
	"github.com/kataras/iris"
	"time"
)

type MyHand struct {
}
type logBody struct {
	Time   time.Time
	Url    string
	Method string
	Result string
	User   interface{} `json:"User"`
}

func (ctx MyHand) HandResponseWriter(result string, context iris.Context) string {
	user := context.Values().Get("user")
	content := &logBody{
		Time:   time.Now(),
		Url:    context.Request().URL.Path,
		Method: context.Request().Method,
		Result: result,
		User:   user,
	}
	logger.Logger(content)
	return result
}
