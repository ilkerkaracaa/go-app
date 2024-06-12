package main

import (
	"context"
	"goapp/common/app"
	"goapp/common/postgresql"
	"goapp/controller"
	"goapp/datalayer"
	"goapp/service"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	server := echo.New()
	configurationManager := app.NewConfigurationManagger()
	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostreSqlConfig)
	prodcutRepository := datalayer.NewProductRepository(dbPool)
	productService := service.NewProductService(prodcutRepository)
	productController := controller.NewProductController(productService)
	productController.RegisterRoutes(server)
	server.Start("localhost:8080")
}
