package paymentRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PaymentRepositoryService interface {}	

	paymentRepository struct {
		db *mongo.Client
	}
)

func NewPaymentRepositoryService(db *mongo.Client) PaymentRepositoryService {
	return &paymentRepository{
		db: db,
	}
}

func (r *paymentRepository) PaymentDbConn(pctx *context.Context) *mongo.Database {
	return r.db.Database("payment_db")
}