package app

import (
	"fmt"
	"net/http"
	"wb_nats_go_service/internal/config"
	"wb_nats_go_service/internal/database/postgresql"
	"wb_nats_go_service/internal/services"
	"wb_nats_go_service/internal/transport/rest"
)

func Run() error {
	// инициализация config файла
	cfg := config.MustLoad()

	// инициализация подключения к базе данных
	db := postgresql.Init(cfg)

	//
	h := rest.New(db)

	services.LoadInCacheInDB(db)

	http.HandleFunc("/order_by_id", h.GetOrderById)
	http.HandleFunc("/send_order", h.SendToNats)
	// TODO: реализовать потом логирование
	fmt.Printf("Starting service on port %v\n", cfg.HTTPServer.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.HTTPServer.Host, cfg.HTTPServer.Port), nil)
	if err != nil {
		return err
	}

	return nil
}
