package bootstrap

import (
	"7youo-wms/global"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Server(addr string, router *gin.Engine)  {

	s := endless.NewServer(addr, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20

	time.Sleep(10 * time.Microsecond)
	global.G_LOG.Info("server run success on ", zap.String("address", addr))

	err := s.ListenAndServe()
	if err != nil {
		global.G_LOG.Error("ListenAndServe failed", zap.Any("err", err))
	}
}
