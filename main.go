package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/skylineCodes/bank_app/api"
	db "github.com/skylineCodes/bank_app/db/sqlc"
	"github.com/skylineCodes/bank_app/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

  store := db.NewStore(connPool)
  server := api.NewServer(util.Config, store)

  err = server.Start(config.ServerAddress)

  if err != nil {
    log.Fatal("cannot start server:", err)
  }
}