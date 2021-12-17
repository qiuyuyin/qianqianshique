package core

import (
	"go.uber.org/zap"
	"net/http"
	"poem_server/global"
	"poem_server/initialize"
	"time"
)

func Server() {
	Router := initialize.Router()
	address := global.POEM_CONFIG.System.Addr

	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	global.POEM_LOG.Info("后端服务开启在 ", zap.String("address", address))

	global.POEM_LOG.Error(s.ListenAndServe().Error())

}
