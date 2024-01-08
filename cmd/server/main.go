package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sergeykhomenko/test-health-checks/internal/app"
	"github.com/sergeykhomenko/test-health-checks/internal/app/services"
	"log/slog"
	"strconv"
)

func main() {
	pingTimeout := flag.Int("timeout", 500, "Single request timeout")
	port := flag.Int("port", 8000, "Port to listen")
	debug := flag.Bool("debug", false, "Run in debug mode")

	flag.Parse()
	services.SetPingbackTimeout(*pingTimeout)

	app.InitLogger(*debug)
	gin.SetMode(gin.ReleaseMode)
	if *debug {
		gin.SetMode(gin.DebugMode)
	}

	router := app.NewRouter()
	slog.Info("Running server on port " + strconv.Itoa(*port))
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		panic("Can't run a server")
	}
}
