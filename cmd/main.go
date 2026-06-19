package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default() // Cria um novo servidor Gin

	dbConnection, err := db.ConnerctDB() // Conecta ao banco de dados
	if err != nil {
		panic(err)
	}

	// Camada de Repository
	productRepository := repository.NewProductRepository(dbConnection)

	// Camada de Use Case
	productUseCase := usecase.NewProductUseCase(productRepository)

	//Camada de Controller
	productController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.POST("/products", productController.CreateProduct)

	server.GET("/products/:productId", productController.GetProductsById)

	server.Run(":8080") // http://localhost:8080/ping
}
