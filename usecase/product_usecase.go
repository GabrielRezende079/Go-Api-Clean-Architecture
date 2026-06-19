package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	//repository
	repository repository.ProductRepository
}

// GetProducts NewProductUseCase creates a new instance of ProductUseCase with the provided repository.
func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

// GetProducts Use case to get products
func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

// CreateProduct Use case to create a product
func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	// Call the repository to create the product and get the ID of the newly created product
	productId, err := pu.repository.CreateProduct(product)
	// Check for errors while creating the product
	if err != nil {
		return model.Product{}, err
	}

	// Set the ID of the product to the ID returned by the repository
	product.ID = productId

	// Return the created product and nil for the error
	return product, nil
}

func (pu *ProductUseCase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err

	}
	return product, nil
}
