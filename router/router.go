package router

import (
	"github.com/Unknwon/macaron"

	"github.com/containerops/crew/handler"
)

func SetRouters(m *macaron.Macaron) {
	//Crew User And Organization V1 API

	m.Group("/w1", func() {
		//Signin and Signup
		m.Post("/signup", handler.W1Signup)
		m.Post("/signin", handler.W1Signin)

		//User Router
		m.Group("/user", func() {
			m.Get("/list", handler.W1GetUserList)
			m.Put("/:userKey/profile", handler.W1PutUserProfile)
			m.Get("/:userKey/profile", handler.W1GetUserProfile)
			m.Post("/:userKey/gravatar", handler.W1PostUserGravatar)
			m.Put("/:userKey/passwd", handler.W1PutUserPasswd)
		})

		//Organization Router
		m.Group("/org", func() {
			m.Post("/", handler.W1PostOrganization)
			m.Put("/:orgKey", handler.W1PutOrganization)
			m.Get("/:orgKey", handler.W1GetOrganization)
			m.Delete("/:orgKey", handler.W1DeleteOrganization)

			//Team Router
			m.Group("/:orgKey/team", func() {
				m.Post("/", handler.W1PostTeam)
				m.Get("/list", handler.W1GetTeams)
				m.Put("/:teamKey", handler.W1PutTeam)
				m.Get("/:teamKey", handler.W1GetTeam)
				m.Delete("/:teamKey", handler.W1DeleteTeam)

				//User Management
				m.Group("/:teamKey/user", func() {
					m.Get("/list", handler.W1GetTeamUsers)
					m.Put("/:userKey", handler.W1PutTeamUser)
					m.Delete("/:userKey", handler.W1DeleteTeamUser)
				})
			})
		})

	})
}
