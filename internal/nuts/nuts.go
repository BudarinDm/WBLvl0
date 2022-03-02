package nuts

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"wblvl0/internal/model"
	"wblvl0/internal/repository"
)

func NewNutsSubscriber(conn stan.Conn, repos *repository.Repository) stan.Subscription {

	handler := func(msg *stan.Msg) {
		var order model.Order

		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Printf("unmarshal error: %s", err.Error())
		}

		_, err = repos.CreateOrder(order)
		if err != nil {
			log.Printf("create order db error: %s", err.Error())
		}
	}

	sub, err := conn.Subscribe(
		"order",
		handler,
		stan.DurableName("stand1"),
	)
	if err != nil {
		log.Fatalf("create subscribe error: %s", err.Error())
	}

	return sub
}
