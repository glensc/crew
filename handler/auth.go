package handler

import (
	"net/http"

	"github.com/Unknwon/macaron"

	"github.com/astaxie/beego/logs"
)

func W1UserSignup(ctx *macaron.Context, log *logs.BeeLogger) (int, []byte) {
	return http.StatusOK, []byte{}
}

func W1UserSignin(ctx *macaron.Context, log *logs.BeeLogger) (int, []byte) {
	return http.StatusOK, []byte{}
}
