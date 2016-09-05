package router

import (
	"gopkg.in/macaron.v1"

	"github.com/containerops/crew/handler"
)

//SetRouters is
func SetRouters(m *macaron.Macaron) {
	//Web API
	m.Get("/", handler.GetIndexPageV1Handler)
}
