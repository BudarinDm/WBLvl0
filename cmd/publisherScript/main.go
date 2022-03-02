package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"time"
)

func main() {
	sc, _ := stan.Connect("prod", "publisher")
	defer sc.Close()

	for i := 0; ; i++ {
		json := fmt.Sprintf(`{
  "order_uid": "b563feb7b2b84b6test%d",
  "track_number": "WBILMTESTTRACK%d",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira %d",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test%d",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817%d,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930%d,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}`, i, i, i, i, i, i)

		//badJson := "{fdsfsdf: sdfsdfdsf , affsasfsfa}"

		time.Sleep(2 * time.Second)
		sc.Publish("order", []byte(json))

		//if i%5 == 2 {
		//	sc.Publish("order", []byte(badJson))
		//}
	}

}
