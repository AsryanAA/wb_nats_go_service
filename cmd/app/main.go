package main

import (
	"log"
	"wb_nats_go_service/internal/app"
)

// точка входа в приложение
func main() {
	err := app.Run()
	if err != nil {
		log.Fatalf("Service can not started, %s", err)
	}
}
