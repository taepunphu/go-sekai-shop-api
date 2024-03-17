package playerRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PlayerRepositoryService interface{}

	palyerRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepositoryService(db *mongo.Client) PlayerRepositoryService {
	return &palyerRepository{
		db: db,
	}
}

func (r *palyerRepository) playerDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("player_db")
}
