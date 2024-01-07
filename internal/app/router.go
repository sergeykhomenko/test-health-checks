package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sergeykhomenko/test-health-checks/internal/app/handlers"
)

func NewRouter() *gin.Engine {
	ping := new(handlers.PingHandlers)

	r := gin.Default()
	r.POST("/ping-urls", ping.PingUrls)

	return r
}
