package model

type Payment struct {
	Transaction  string `json:"transaction" db:"transaction"`
	RequestID    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       uint   `json:"amount" db:"amount"`
	PaymentDt    uint   `json:"payment_dt" db:"payment_dt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost uint   `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   uint   `json:"goods_total" db:"goods_total"`
	CustomFee    uint   `json:"custom_fee" db:"custom_fee"`
	OrderUid     string `json:"order_uid" db:"order_uid"`
}
