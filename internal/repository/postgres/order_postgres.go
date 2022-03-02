package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"wblvl0/internal/model"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) GetOrder(uid string) (model.Order, error) {
	var order model.Order

	query := fmt.Sprintf(`SELECT * FROM orders WHERE order_uid=$1`)
	err := r.db.Get(&order, query, uid)

	delivery, err := r.GetDelivery(uid)
	items, err := r.GetItems(uid)
	payment, err := r.GetPayments(uid)

	order.Delivery = delivery
	order.Items = items
	order.Payment = payment

	return order, err
}

func (r *OrderPostgres) CreateOrder(order model.Order) (string, error) {
	var uid string

	query := `INSERT INTO orders
				(
					order_uid,
					track_number,
					entry,
					locale,
					internal_signature,
					customer_id,
					delivery_service,
					shardkey,
					sm_id,
					date_created,
					oof_shard
				)
			    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := r.db.Exec(query, order.UID, order.TrackNumber, order.Entry, order.Locale,
		order.InternalSignature, order.CustomerId, order.DeliveryService, order.ShardKey,
		order.SmId, order.DateCreated, order.OofShard)

	err = r.CreatePayment(order.Payment, order.UID)
	if err != nil {
		return "", err
	}

	err = r.CreateDelivery(order.Delivery, order.UID)
	if err != nil {
		return "", err
	}

	for item := range order.Items {
		err = r.CreateItem(order.Items[item], order.UID)
		if err != nil {
			return "", err
		}
	}

	return uid, nil
}

func (r *OrderPostgres) CreateDelivery(delivery model.Delivery, uid string) error {
	query := `INSERT INTO deliveries
		(
			name,
			phone,
			zip,
			city,
			address,
			region,
			email,
			order_uid
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(query, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address,
		delivery.Region, delivery.Email, uid)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderPostgres) CreateItem(item model.Item, uid string) error {
	query := `INSERT INTO items
		(
			chrt_id,
			track_number,
			price,
			rid,
			name,
			sale,
			size,
			total_price,
			nm_id,
			brand,
			status,
			order_uid
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := r.db.Exec(query, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale,
		item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status, uid)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderPostgres) CreatePayment(payment model.Payment, uid string) error {
	query := `INSERT INTO payments
		(
			transaction,
			request_id,
			currency,
			provider,
			amount,
			payment_dt,
			bank,
			delivery_cost,
			goods_total,
			custom_fee,
			order_uid
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := r.db.Exec(query, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider,
		payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal,
		payment.CustomFee, uid)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderPostgres) GetDelivery(uid string) (model.Delivery, error) {
	var delivery model.Delivery

	query := fmt.Sprintf("SELECT * FROM deliveries WHERE order_uid=$1")
	err := r.db.Get(&delivery, query, uid)

	return delivery, err
}

func (r *OrderPostgres) GetPayments(uid string) (model.Payment, error) {
	var payment model.Payment

	query := fmt.Sprintf("SELECT * FROM payments WHERE order_uid=$1")
	err := r.db.Get(&payment, query, uid)

	return payment, err
}

func (r *OrderPostgres) GetItems(uid string) ([]model.Item, error) {
	var items []model.Item

	query := fmt.Sprintf("SELECT * FROM items WHERE order_uid=$1")
	err := r.db.Get(&items, query, uid)

	return items, err
}
