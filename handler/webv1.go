package handler

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

//GetIndexPageV1Handler is
func GetIndexPageV1Handler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}
