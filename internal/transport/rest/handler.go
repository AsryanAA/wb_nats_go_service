package rest

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"wb_nats_go_service/internal/services"
	n "wb_nats_go_service/internal/transport/nats"
	"wb_nats_go_service/pkg"

	// "encoding/json"
	"gorm.io/gorm"
	"net/http"
	"wb_nats_go_service/internal/models"
	// "wb_nats_go_service/pkg"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

func (h Handler) SendToNats(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	var resp models.Response
	resp.Status = http.StatusBadRequest
	err := json.NewDecoder(r.Body).Decode(&order)
	data, err := json.Marshal(order)
	if err != nil {
		resp.Message = pkg.MsgNotParseJSON
	} else {
		nc, err := n.Init()
		if err != nil {
			resp.Message = pkg.MsgDoNotInitNats
			response, _ := json.Marshal(resp)
			w.Write(response)
			return
		} else {
			nc.Subscribe("orders", func(msg *nats.Msg) {
				// fmt.Printf("Сообщение получено: %s\n", string(msg.Data))
			})
			err = nc.Publish("orders", data)
			if err != nil {
				resp.Message = pkg.MsgDoNotSendToNats
			} else {
				err = services.SaveInDB(h.DB, order)
				err = services.SaveInCache(order)
				if err != nil {
					fmt.Println("Сообщение не сохранено в БД")
					return
				}
				resp.Status = http.StatusOK
				resp.Message = pkg.MsgOrderCreated
			}
		}
	}
	response, _ := json.Marshal(resp)
	w.Write(response)
	return
}

// GetOrderById GetUsersByRecordingDateOrAge возвращает пользователей по определенным критериям
func (h Handler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	id := keys.Get("order_id")
	order := services.ReadInCache(id)
	resp, _ := json.Marshal(order)
	fmt.Println(keys)
	w.Write(resp)
}
