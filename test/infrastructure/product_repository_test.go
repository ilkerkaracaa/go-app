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
	fmt.Println("Before tests")
	exitCode := m.Run()
	fmt.Println("After tests")
	os.Exit(exitCode)
}

func TestGetAllProduct(t *testing.T) {

}
