package router

import (
	"github.com/Unknwon/macaron"

	"github.com/containerops/crew/handler"
)

func SetRouters(m *macaron.Macaron) {
	//Crew User And Organization V1 API

	m.Group("/w1", func() {
		m.Post("/signup", handler.W1Signup)
	})
}
