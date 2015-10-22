package handler

import (
	"net/http"

	"gopkg.in/macaron.v1"

	"github.com/astaxie/beego/logs"
)

func W1PostToken(ctx *macaron.Context, log *logs.BeeLogger) (int, []byte) {
	return http.StatusOK, []byte{}
}
