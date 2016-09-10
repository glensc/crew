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

//PostUserV1Handler is a handler function of creating a user.
func PostUserV1Handler(ctx *macaron.Context, user User) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//GetUserV1Handler is a handler function of getting a user data.
func GetUserV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PatchUserV1Handler is a handler function of updating a user data.
func PatchUserV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//DeleteUserV1Handler is a handler function of deleting a user in system.
//But in system we never delete data, we just use soft delete function in Gorm.
func DeleteUserV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PostUserPasswordV1Handler is a handler function of changing a user's password.
func PostUserPasswordV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//GetUserEmailsV1Handler is a handler function of listing a user's emails.
func GetUserEmailsV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PostUserEmailsV1Handler is a handler function of adding email address(es).
func PostUserEmailsV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PatchUserEmailsV1Handler is a handler function of setting an email is default email.
func PatchUserEmailsV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//DeleteUserEmailsV1Handler is a handler function of deleting an email.
//But the default email shouldn't be delete.
func DeleteUserEmailsV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}
