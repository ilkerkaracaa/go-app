package infrastructure

import (
	"context"
	"fmt"
	"goapp/common/postgresql"
	"goapp/datalayer"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var productRepositorty datalayer.IProductRepository
var dbPool *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()
	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "goapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})
	productRepositorty = datalayer.NewProductRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestGetAllProduct(t *testing.T) {
	fmt.Println("Test get all products")
	fmt.Println(productRepositorty)
	fmt.Println(dbPool)
}
