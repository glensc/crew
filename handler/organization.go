package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/macaron.v1"
)

//PostOrganizationV1Handler is a handler function which creating an organization.
func PostOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//GetOrganizationV1Handler is a handler function which getting an organization data.
func GetOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PatchOrganizationV1Handler is a handler function which updating an organization data.
func PatchOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//DeleteOrganizationV1Handler is a handler function which delete an organization data.
func DeleteOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//GetOrgMembersV1Handler is a handler function which is getting users of an organization.
func GetOrgMembersV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//CheckOrgMemberV1Handler is a handler function which is checking member of an organization.
func CheckOrgMemberV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PutOrgMemberV1Handler is a handler function which is adding member of an organization.
func PutOrgMemberV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//DeleteOrgMemberV1Handler is a handler function which is deleting member of an organization.
func DeleteOrgMemberV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}
