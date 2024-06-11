package datalayer

import (
	"context"
	"errors"
	"fmt"
	"goapp/domain"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
	AddProduct(prdouct domain.Product) error
	GetById(productId int64) (domain.Product, error)
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool: dbPool,
	}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {

	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "Select * From products")
	if err != nil {
		log.Error("Query not working", err)
		return []domain.Product{}
	}
	return extractProductsFromRows(productRows)
}

func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {

	ctx := context.Background()
	getProductsByStoreNameQuery := `Select * From products where store = $1`
	productRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNameQuery, storeName)
	if err != nil {
		log.Error("Query not working", err)
		return []domain.Product{}
	}
	return extractProductsFromRows(productRows)
}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()
	insert_query := `insert into products (name,price,discount,store) VALUES ($1,$2,$3,$4)`
	addNewProduct, err := productRepository.dbPool.Exec(ctx, insert_query, product.Name, product.Price, product.Discount, product.Store)
	if err != nil {
		log.Error("faild to add new product", err)
		return err
	}
	log.Info(fmt.Printf("Product added with %v", addNewProduct))
	return nil
}

func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string
	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}
	return products
}

func (productRepository *ProductRepository) GetById(productId int64) (domain.Product, error) {
	ctx := context.Background()
	getByIdQuery := `Select * From products where id = $1`
	queryRow := productRepository.dbPool.QueryRow(ctx, getByIdQuery, productId)
	var id int64
	var name string
	var price float32
	var discount float32
	var store string
	err := queryRow.Scan(&id, &name, &price, &discount, &store)
	if err != nil {
		return domain.Product{}, errors.New(fmt.Sprintf("Error getting with id %d", productId))
	}
	return domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}, nil
}
