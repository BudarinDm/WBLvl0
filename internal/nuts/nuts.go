package nuts

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"log"
	"wblvl0/internal/cache"
	"wblvl0/internal/model"
	"wblvl0/internal/repository"
)

func NewConnection() (stan.Conn, error) {
	natsConnection, err := stan.Connect(viper.GetString("nuts.stanclusterid"), viper.GetString("nuts.clientid"))
	if err != nil {
		return nil, err
	}
	return natsConnection, err
}

func NewNutsSubscriber(conn stan.Conn, repos *repository.Repository, cache *cache.Cache) stan.Subscription {

	handler := func(msg *stan.Msg) {
		var order model.Order

		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Printf("unmarshal error: %s", err.Error())
			return
		}

		_, err = repos.CreateOrder(order)
		if err != nil {
			log.Printf("create order db error: %s", err.Error())
			return
		}
		cache.Add(order)
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
