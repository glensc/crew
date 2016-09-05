package middleware

import (
	"gopkg.in/macaron.v1"
)

//SetMiddlewares is
func SetMiddlewares(m *macaron.Macaron) {
	m.Use(macaron.Static("static", macaron.StaticOptions{
		Expires: func() string { return "max-age=0" },
	}))

	//Set log
	m.Use(logger())

	m.Use(macaron.Recovery())
}
