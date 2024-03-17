package itemRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ItemRepositoryService interface{}

	itemRepository struct {
		db *mongo.Client
	}
)

func NewItemRepositoryService(db *mongo.Client) ItemRepositoryService {
	return &itemRepository{
		db: db,
	}
}

func (r *itemRepository) itemDbconn(pctx context.Context) *mongo.Database {
	return r.db.Database("item_db")
}
