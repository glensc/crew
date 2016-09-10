package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/macaron.v1"
)

//GetSearchUsersV1Handler is handler function of finding user via various criteria.  (This method returns up to 100 results per page.)
func GetSearchUsersV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}
