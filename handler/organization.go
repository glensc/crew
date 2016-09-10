package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/macaron.v1"
)

//PostOrganizationV1Handler is a handler function of creating an organization.
func PostOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//GetOrganizationV1Handler is a handler function of getting an organization data.
func GetOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PatchOrganizationV1Handler is a handler function of updating an organization data.
func PatchOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//DeleteOrganizationV1Handler is a handler function of delete an organization data.
func DeleteOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}
