package repository

import (
	"github.com/jmoiron/sqlx"
	"wblvl0/internal/model"
	"wblvl0/internal/repository/postgres"
)

type Order interface {
	CreateOrder(order model.Order) (string, error)
	GetOrder(uid string) (model.Order, error)
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: postgres.NewOrderPostgres(db),
	}
}
