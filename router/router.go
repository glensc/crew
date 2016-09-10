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
			m.Get("/:username", handler.GetUserV1Handler)
			m.Patch("/:username", handler.PatchUserV1Handler)
			m.Delete("/:username", handler.DeleteUserV1Handler)

			m.Post("/:username/password", handler.PostUserPasswordV1Handler)
			m.Get("/:username/emails", handler.GetUserEmailsV1Handler)
			m.Post("/:username/emails", handler.PostUserEmailsV1Handler)
			m.Patch("/:username/emails/default", handler.PatchUserEmailsV1Handler)
			m.Delete("/:username/emails", handler.DeleteUserEmailsV1Handler)
		})

		//Organization
		m.Group("/orgs", func() {
			m.Post("/", handler.PostOrganizationV1Handler)
			m.Get("/:org", handler.GetOrganizationV1Handler)
			m.Patch("/:org", handler.PatchOrganizationV1Handler)
			m.Delete("/org", handler.DeleteOrganizationV1Handler)

			//Team
			m.Group("/team", func() {
				m.Post("/", handler.PostTeamV1Handler)
				m.Get("/:team", handler.GetTeamV1Handler)
				m.Patch("/:team", handler.PatchTeamV1Handler)
				m.Delete("/:team", handler.DeleteTeamV1Handler)
			})
		})

		//Search
		m.Group("/search", func() {
			m.Get("/users", handler.GetSearchUsersV1Handler)
		})
	})
}
