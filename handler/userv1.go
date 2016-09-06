package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/macaron.v1"
)

//User is a form schema for create a user account.
type User struct {
	Name     string `from:"name" binding:"Required"`     //Name of user account
	Email    string `from:"email" binding:"Required"`    //Email of user account
	FullName string `from:"fullname" binding:"Required"` //Full name of account
	Password string `from:"Password" binding:"Required"` //Password for the user
}

//PostUserV1Handler is create a user.
func PostUserV1Handler(ctx *macaron.Context, user User) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//GetUserV1Handler is
func GetUserV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PatchUserV1Handler is
func PatchUserV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//DeleteUserV1Handler is create a user.
func DeleteUserV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PostUserPasswordV1Handler is create a user.
func PostUserPasswordV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}
