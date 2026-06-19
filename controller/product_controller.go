package controller

import (
	"go-api/model"
	"go-api/usecase"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//use case
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

// GetProduct
func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

	}
	ctx.JSON(http.StatusOK, products)
}

// CreateProduct Controller to create a product
func (p *productController) CreateProduct(ctx *gin.Context) {
	// Bind the JSON payload from the request body to a product model
	var product model.Product
	// Check for errors while binding the JSON payload
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	// Call the use case to create the product and get the inserted product with the ID
	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	// Return the inserted product as JSON response with status OK
	ctx.JSON(http.StatusOK, insertedProduct)
}

// GetProductById
func (p *productController) GetProductsById(ctx *gin.Context) {

	//
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "O Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "O Id do produto precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return

	}

	ctx.JSON(http.StatusOK, product)

}
