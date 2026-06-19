package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	//db connection
	connection *sql.DB
}

// NewProductRepository creates a new instance of ProductRepository with the provided database connection.
func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}

}

// GetProducts
func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	// Define the SQL query to retrieve products from the database
	query := "SELECT id, product_name, price FROM product"
	// Execute the query and get the result set
	rows, err := pr.connection.Query(query)

	// Check for errors while executing the query
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Product{}, err
	}

	var productList []model.Product // Initialize an empty slice to hold the products
	var productObj model.Product    // Create a variable to hold each product as we iterate through the rows

	// Iterate through the result set and scan each row into the productObj
	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Product_name,
			&productObj.Price,
		)

		// Check for errors while scanning the row
		if err != nil {
			fmt.Println("Error executing query:", err)
			return []model.Product{}, err
		}

		// Append the scanned product object to the productList slice
		productList = append(productList, productObj)

	}

	// Close the rows to free up resources
	rows.Close()

	// Return the list of products and nil for the error
	return productList, nil
}

// CreateProduct and return the ID of the newly created product
func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product " +
		"(product_name, price)" +
		"VALUES ($1, $2) RETURNING id")
	// Check for errors while preparing the query
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return 0, err
	}
	// Execute the query with the provided product details
	err = query.QueryRow(product.Product_name, product.Price).Scan(&id)

	// Check for errors while executing the query
	if err != nil {
		fmt.Println("Error executing query:", err)
		return 0, err
	}

	// Close the query to free up resources
	query.Close()
	// Return the ID of the newly created product and nil for the error
	return id, nil
}

// GetProduct by ID
func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {
	// Executes the SQL query searching for ID
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")

	// Look for erros while preparing the query
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Product_name,
		&produto.Price,
	)

	//error catch for "product not found"
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Query Closed
	query.Close()
	return &produto, nil

}
