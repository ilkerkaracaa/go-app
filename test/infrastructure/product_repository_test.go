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
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()
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

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}
func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}
func TestGetAllProduct(t *testing.T) {
	setup(ctx, dbPool)
	fmt.Println("Test all products")
	clear(ctx, dbPool)
}
