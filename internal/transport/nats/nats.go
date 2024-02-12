package nats

import "github.com/nats-io/nats.go"

// Init Функция для подключения к серверу NATS
func Init() (*nats.Conn, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	return nc, nil
}
