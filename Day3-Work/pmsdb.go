package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

// Product class in java
type Product struct {
	id         int
	pname      string
	ptype      string
	prod_place string
	warranty   int
}

// ProductManager is serviceperson who will maintain the list of products, display the products, add a product, give the product
// In Java it will be ArrayList
type ProductManager struct {
	db *sql.DB
}

func NewProductManager(db *sql.DB) *ProductManager {
	return &ProductManager{
		db: db,
	}
}

// AddProduct method -- adds a new product to the list
func (pm *ProductManager) AddProduct(product Product) {
	_, err := pm.db.Exec("INSERT INTO product (pname, ptype, prod_place, warranty) VALUES ($1, $2, $3, $4)",
		product.pname, product.ptype, product.prod_place, product.warranty)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllProducts method -- Display a list of products and return slice of Products
func (pm *ProductManager) GetAllProducts() []Product {
	rows, err := pm.db.Query("SELECT * FROM product")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	products := make([]Product, 0)
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.id, &product.pname, &product.ptype, &product.prod_place, &product.warranty)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products
}

// GetProduct Method -- Display a Particular Product from the list of Products which returns the required product
// if nothing found it returns nil
func (pm *ProductManager) GetProduct(pname string) *Product {
	var product Product
	err := pm.db.QueryRow("SELECT * FROM product WHERE pname = $1", pname).
		Scan(&product.id, &product.pname, &product.ptype, &product.prod_place, &product.warranty)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		log.Fatal(err)
	}
	return &product
}

// SearchProductWithText -- Returns a list of products based on search keyword
func (pm *ProductManager) SearchProductWithText(text string) []Product {
	results := make([]Product, 0)
	rows, err := pm.db.Query("SELECT * FROM product WHERE LOWER(pname) LIKE $1 OR LOWER(ptype) LIKE $1 OR LOWER(prod_place) LIKE $1", "%"+strings.ToLower(text)+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.id, &product.pname, &product.ptype, &product.prod_place, &product.warranty)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, product)
	}

	return results
}

// SearchResultsWithPlace -- Returns a list of products available at a specific place
func (pm *ProductManager) SearchResultsWithPlace(place string) []Product {
	results := make([]Product, 0)
	rows, err := pm.db.Query("SELECT * FROM product WHERE LOWER(prod_place) = LOWER($1)", place)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.id, &product.pname, &product.ptype, &product.prod_place, &product.warranty)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, product)
	}

	return results
}

// SearchResultsOutofWarranty -- Returns a list of products that are out of warranty
func (pm *ProductManager) SearchResultsOutofWarranty(year int) []Product {
	results := make([]Product, 0)
	rows, err := pm.db.Query("SELECT * FROM product WHERE warranty < $1", year)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.id, &product.pname, &product.ptype, &product.prod_place, &product.warranty)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, product)
	}

	return results
}

func main() {
	db, err := sql.Open("postgres", "host= 127.0.0.1 port=5432 user=testuser dbname=postgres password=1234 sslmode=disable")
	//db, err := sql.Open("postgres", "postgres://postgres:your_password@localhost/product_management?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pm := NewProductManager(db)

	// Adding to list
	product1 := Product{pname: "Chair", ptype: "Furniture", prod_place: "Mumbai", warranty: 2024}
	product2 := Product{pname: "VoltasAC", ptype: "AirConditioner", prod_place: "Chennai", warranty: 2022}
	product3 := Product{pname: "VoltasAC", ptype: "AirConditioner", prod_place: "Kolkata", warranty: 2022}
	pm.AddProduct(product1)
	pm.AddProduct(product2)
	pm.AddProduct(product3)

	// Invoke the GetAllProducts method
	products := pm.GetAllProducts()
	fmt.Println("Printing All Products in one line:", products)

	// Prints one by one the Product from the list of products
	for _, product := range products {
		fmt.Println("-----------")
		fmt.Println(product)
		fmt.Printf("Product Name: %s\n", product.pname)
		fmt.Printf("Product Type: %s\n", product.ptype)
		fmt.Printf("Product Place: %s\n", product.prod_place)
		fmt.Printf("Product Warranty: %d\n", product.warranty)
		fmt.Println("-----------")
	}

	// GetProduct
	product := pm.GetProduct("Chair")
	if product != nil {
		fmt.Println("Returned Selected Product:", product)
		fmt.Printf("Product Name: %s\n", product.pname)
		fmt.Printf("Product Type: %s\n", product.ptype)
		fmt.Printf("Product Place: %s\n", product.prod_place)
		fmt.Printf("Product Warranty: %d\n", product.warranty)
		fmt.Println("-----------")
	}

	// SearchProductWithText
	fmt.Println("=============Search Products for a Particular text================")
	fmt.Println("Enter what you want to search without spaces:")
	var searchProduct string
	fmt.Scan(&searchProduct)
	searchResults := pm.SearchProductWithText(searchProduct)
	if len(searchResults) > 0 {
		fmt.Println("Search Results:")
		for _, product := range searchResults {
			fmt.Printf("Product Name: %s\n", product.pname)
			fmt.Printf("Product Type: %s\n", product.ptype)
			fmt.Printf("Product Place: %s\n", product.prod_place)
			fmt.Printf("Product Warranty: %d\n", product.warranty)
			fmt.Println("-----------")
		}
	} else {
		fmt.Println("No products found matching the search criteria.")
	}

	fmt.Println("=============Search products at a Place==============")
	fmt.Println("Enter what you want to search:")
	var searchPlace string
	fmt.Scan(&searchPlace)
	searchResultsByPlace := pm.SearchResultsWithPlace(searchPlace)

	if len(searchResultsByPlace) > 0 {
		for _, product := range searchResultsByPlace {
			fmt.Printf("Product Name: %s\n", product.pname)
			fmt.Printf("Product Type: %s\n", product.ptype)
			fmt.Printf("Product Place: %s\n", product.prod_place)
			fmt.Printf("Product Warranty: %d\n", product.warranty)
			fmt.Println("-----------")
		}

	} else {
		fmt.Println("No products found in the specified place.")
	}

	fmt.Println("=============Search for products out of warranty==============")
	fmt.Println("Enter current Year:")
	var searchWarranty int
	fmt.Scan(&searchWarranty)
	searchResultsByWarranty := pm.SearchResultsOutofWarranty(searchWarranty)
	if len(searchResultsByWarranty) > 0 {
		for _, product := range searchResultsByWarranty {
			fmt.Println("Products out of warranty:")
			fmt.Printf("Product Name: %s\n", product.pname)
			fmt.Printf("Product Type: %s\n", product.ptype)
			fmt.Printf("Product Place: %s\n", product.prod_place)
			fmt.Printf("Product Warranty: %d\n", product.warranty)
			fmt.Println("-----------")
		}

	} else {
		fmt.Println("All products found are within the warranty range.")
	}
}
