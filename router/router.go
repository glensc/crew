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
		})
	})
}
