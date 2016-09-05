package middleware

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/macaron.v1"

	"github.com/containerops/dockyard/setting"
)

func logger() macaron.Handler {
	return func(ctx *macaron.Context) {
		if setting.RunMode == "dev" {
			log.Info("------------------------------------------------------------------------------")
			log.Info(time.Now().String())
		}

		log.WithFields(log.Fields{
			"Method": ctx.Req.Method,
			"URL":    ctx.Req.RequestURI,
		}).Info(ctx.Req.Header)

	}
}
