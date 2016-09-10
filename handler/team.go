package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/macaron.v1"
)

//GetTeamsV1Handler is a handler function of list teams of an organization.
func GetTeamsV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PostTeamV1Handler is a handler function of creating a team.
func PostTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//GetTeamV1Handler is a handler function of getting a team data.
func GetTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PatchTeamV1Handler is a handler function of updating a team data.
func PatchTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//DeleteTeamV1Handler is a handler function of deleting a team data.
func DeleteTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//GetTeamMembersV1Handler is a handler function of getting members of a team.
func GetTeamMembersV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//GetTeamUserMembershipV1Handler is a handler function of getting a user's membership with a team, the team must be visible to the authenticated user.
func GetTeamUserMembershipV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PutTeamUserMembershipV1Handler is a handler function of adding user to a team.
//If the user is already a member of the team's organization, this endpoint will add the user to the team. In order to add a membership between an organization member and a team, the authenticated user must be an organization owner or a maintainer of the team.
//If the user is unaffiliated with the team's organization, this endpoint will send an invitation to the user via email. This newly-created membership will be in the "pending" state until the user accepts the invitation, at which point the membership will transition to the "active" state and the user will be added as a member of the team. In order to add a membership between an unaffiliated user and a team, the authenticated user must be an organization owner.
func PutTeamUserMembershipV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//DeleteTeamUserMembershipV1Handler is a handler function of deleting user from a team.
//In order to remove a membership between a user and a team, the authenticated user must have 'admin' permissions to the team or be an owner of the organization that the team is associated with.
func DeleteTeamUserMembershipV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//GetTeamReposV1Handler is a handler function of listing repositories in a team.
func GetTeamReposV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//PutTeamRepoV1Handler is a handler function of adding or updating repository in a team.
func PutTeamRepoV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

//DeleteTeamRepoV1Handler is a handler function of deleting repository in a team.
func DeleteTeamRepoV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}
