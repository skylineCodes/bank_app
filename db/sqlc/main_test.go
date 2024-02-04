package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
  dbDriver = "postgres"
  dbSource = "postgresql://root:simplebankgolang@localhost:5432/simple_bank?sslmode=disable"
)

var testStore SQLStore

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), dbSource)

  if err != nil {
    log.Fatal("Cannot connect to db:", err)
  }

  testStore = *NewStore(connPool)

  os.Exit(m.Run())
}