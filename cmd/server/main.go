package main

import (
	"github.com/sergeykhomenko/test-health-checks/internal/app"
	"log"
)

func main() {
	router := app.NewRouter()
	err := router.Run("0.0.0.0:8000")
	if err != nil {
		log.Panic("Can't run a server")
	}
}
