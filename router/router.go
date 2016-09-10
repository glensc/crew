package router

import (
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"

	"github.com/containerops/crew/handler"
)

//SetRouters is
func SetRouters(m *macaron.Macaron) {
	//Web API
	m.Get("/", handler.GetIndexPageV1Handler)

	//V1 Version API
	m.Group("/v1", func() {
		//User
		m.Group("/user", func() {
			m.Post("/", binding.Bind(handler.User{}, handler.PostUserV1Handler))
			m.Get("/:user", handler.GetUserV1Handler)
			m.Patch("/:user", handler.PatchUserV1Handler)
			m.Delete("/:user", handler.DeleteUserV1Handler)

			m.Post("/:user/password", handler.PostUserPasswordV1Handler)
			m.Get("/:user/emails", handler.GetUserEmailsV1Handler)
			m.Post("/:user/emails", handler.PostUserEmailsV1Handler)
			m.Patch("/:user/emails/default", handler.PatchUserEmailsV1Handler)
			m.Delete("/:user/emails", handler.DeleteUserEmailsV1Handler)

			m.Get("/:user/ssh_keys", handler.GetUserSSHKeysV1Handler)
			m.Get("/:user/ssh_keys/:id", handler.GetUserSSHKeyV1Handler)
			m.Post("/:user/ssh_keys", handler.PostUserSSHKeyV1Handler)
			m.Delete("/:user/ssh_keys/:id", handler.DeleteUserSSHKeyV1Handler)

			m.Get("/:user/gpg_keys", handler.GetUserGPGKeysV1Handler)
			m.Get("/:user/gpg_keys/:id", handler.GetUserGPGKeyV1Handler)
			m.Post("/:user/gpg_keys", handler.PostUserGPGKeyV1Handler)
			m.Delete("/:user/gpg_keys/:id", handler.DeleteUserGPGKeyV1Handler)

			m.Get("/teams", handler.GetUserTeamsV1Handler)
			m.Get("/orgs", handler.GetUserOrgsV1Handler)
		})

		//Organization
		m.Group("/orgs", func() {
			m.Post("/", handler.PostOrganizationV1Handler)
			m.Get("/:org", handler.GetOrganizationV1Handler)
			m.Patch("/:org", handler.PatchOrganizationV1Handler)
			m.Delete("/org", handler.DeleteOrganizationV1Handler)

			m.Group("/members", func() {
				m.Get("/", handler.GetOrgMembersV1Handler)
				m.Get("/:user", handler.CheckOrgMemberV1Handler)
				m.Put("/:user", handler.PutOrgMemberV1Handler)
				m.Delete("/:user", handler.DeleteOrgMemberV1Handler)
			})

			//Team
			m.Get("/teams", handler.GetTeamsV1Handler)

			m.Group("/team", func() {
				m.Post("/", handler.PostTeamV1Handler)
				m.Get("/:team", handler.GetTeamV1Handler)
				m.Patch("/:team", handler.PatchTeamV1Handler)
				m.Delete("/:team", handler.DeleteTeamV1Handler)

				m.Group("/:team", func() {
					m.Get("/members", handler.GetTeamMembersV1Handler)

					m.Group("/memberships", func() {
						m.Get("/:user", handler.GetTeamUserMembershipV1Handler)
						m.Put("/:user", handler.PutTeamUserMembershipV1Handler)
						m.Delete("/:user", handler.DeleteTeamUserMembershipV1Handler)
					})

					m.Group("/repos", func() {
						m.Get("/", handler.GetTeamReposV1Handler)
						m.Put("/:repo", handler.PutTeamRepoV1Handler)
						m.Delete("/:repo", handler.DeleteTeamRepoV1Handler)
					})
				})
			})
		})

		//Search
		m.Group("/search", func() {
			m.Get("/users", handler.GetSearchUsersV1Handler)
		})
	})
}
