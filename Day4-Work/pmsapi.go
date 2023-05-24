package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Product class
type Product struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	ProdPlace string `json:"prod_place"`
	Warranty  int    `json:"warranty"`
}

// ProductManager handles product operations
type ProductManager struct {
	db *sql.DB
}

// NewProductManager creates a new instance of ProductManager
func NewProductManager(db *sql.DB) *ProductManager {
	return &ProductManager{
		db: db,
	}
}

// AddProduct adds a new product
func (pm *ProductManager) AddProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := pm.db.Exec("INSERT INTO product (pname, ptype, prod_place, warranty) VALUES ($1, $2, $3, $4)",
		product.Name, product.Type, product.ProdPlace, product.Warranty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully"})
}

// GetAllProducts retrieves all products
func (pm *ProductManager) GetAllProducts(c *gin.Context) {
	rows, err := pm.db.Query("SELECT * FROM product")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	products := make([]Product, 0)
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.ProdPlace, &product.Warranty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

// GetProduct retrieves a specific product by name
func (pm *ProductManager) GetProduct(c *gin.Context) {
	pname := c.Param("pname")

	var product Product
	err := pm.db.QueryRow("SELECT * FROM product WHERE pname = $1", pname).
		Scan(&product.ID, &product.Name, &product.Type, &product.ProdPlace, &product.Warranty)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// SearchProductWithText searches products based on text
func (pm *ProductManager) SearchProductWithText(c *gin.Context) {
	text := strings.ToLower(c.Query("text"))

	results := make([]Product, 0)
	rows, err := pm.db.Query("SELECT * FROM product WHERE LOWER(pname) LIKE $1 OR LOWER(ptype) LIKE $1 OR LOWER(prod_place) LIKE $1", "%"+text+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.ProdPlace, &product.Warranty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, product)
	}

	c.JSON(http.StatusOK, results)
}

// SearchResultsWithPlace searches products based on place
func (pm *ProductManager) SearchResultsWithPlace(c *gin.Context) {
	place := strings.ToLower(c.Query("place"))

	results := make([]Product, 0)
	rows, err := pm.db.Query("SELECT * FROM product WHERE LOWER(prod_place) = LOWER($1)", place)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.ProdPlace, &product.Warranty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, product)
	}

	c.JSON(http.StatusOK, results)
}

// SearchResultsOutofWarranty searches products that are out of warranty
func (pm *ProductManager) SearchResultsOutofWarranty(c *gin.Context) {
	year := c.Query("year")

	results := make([]Product, 0)
	rows, err := pm.db.Query("SELECT * FROM product WHERE warranty < $1", year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.ProdPlace, &product.Warranty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, product)
	}

	c.JSON(http.StatusOK, results)
}

// Main
func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=testuser dbname=postgres password=XXXX sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pm := NewProductManager(db)

	r := gin.Default()

	r.POST("/products", pm.AddProduct)
	r.GET("/products", pm.GetAllProducts)
	r.GET("/products/:pname", pm.GetProduct)
	r.GET("/products/search", pm.SearchProductWithText)
	r.GET("/products/place", pm.SearchResultsWithPlace)
	r.GET("/products/warranty", pm.SearchResultsOutofWarranty)

	if err := r.Run(":8009"); err != nil {
		log.Fatal(err)
	}
}
