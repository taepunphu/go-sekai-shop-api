package main

import (
	"context"
	"log"
	"os"

	"github.com/taepunphu/go-sekai-shop-api/config"
	"github.com/taepunphu/go-sekai-shop-api/pkg/database"
	"github.com/taepunphu/go-sekai-shop-api/server"
)

func main() {
	ctx := context.Background()

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: no config file path")
		}
		return os.Args[1]
	}())

	// database connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)
	log.Println(db)

	server.Start(ctx, &cfg, db)
}
