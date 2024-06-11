package infrastructure

import (
	"context"
	"fmt"
	"goapp/common/postgresql"
	"goapp/datalayer"
	"goapp/domain"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
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
	expected := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Lambader",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Dekorasyon Sarayı",
		},
	}
	t.Run("Get all products", func(t *testing.T) {
		actualProducts := productRepositorty.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expected, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestGetAllProductByStore(t *testing.T) {
	setup(ctx, dbPool)
	expected := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
	}
	t.Run("Get all products by store", func(t *testing.T) {
		actualProducts := productRepositorty.GetAllProductsByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, expected, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestAddProduct(t *testing.T) {
	expected := []domain.Product{
		{
			Id:       1,
			Name:     "Kupa",
			Price:    100,
			Discount: 0.0,
			Store:    "Kırtasiye",
		},
	}
	newProduct := domain.Product{
		Name:     "Kupa",
		Price:    100,
		Discount: 0.0,
		Store:    "Kırtasiye",
	}
	t.Run("Add Product", func(t *testing.T) {
		productRepositorty.AddProduct(newProduct)
		products := productRepositorty.GetAllProducts()
		assert.Equal(t, 1, len(products))
		assert.Equal(t, expected, products)
	})
	clear(ctx, dbPool)
}

func TestGetProductById(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("Get Product by Id", func(t *testing.T) {
		product, _ := productRepositorty.GetById(1)
		_, err := productRepositorty.GetById(5)
		assert.Equal(t, domain.Product{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		}, product)
		assert.Equal(t, "Error getting with id 5", err.Error())
	})
	clear(ctx, dbPool)
}
