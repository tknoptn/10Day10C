package main

import (
	"fmt"
	"strings"
)

// Product class in java
type Product struct {
	pname      string
	ptype      string
	prod_place string
	warranty   int
}

// ProductManager is serviceperson who will maintain the list of products, display the products ,add a product ,  give the product
// In Java it will be ArrayList
type ProductManager struct {
	products []Product // products will contain list of Product
}

func NewProductManager() *ProductManager {
	return &ProductManager{
		products: make([]Product, 0),
	}

}

// AddProduct method -- adds a new product to the list
func (pm *ProductManager) AddProduct(product Product) {
	pm.products = append(pm.products, product)
}

// GetAllProdcuts method --  Display a list of products and return slice of Products
func (pm *ProductManager) GetAllProducts() []Product {
	return pm.products
}

// GetProduct  Method -- Display a Particular Product from the list of Products which returns the required product
// if nothing found  it reurns null
func (pm *ProductManager) GetProduct(pname string) *Product {
	for _, product := range pm.products { // for eachloop in java
		if product.pname == pname { // isequals in java
			return &product
		}
	}
	return nil
}

// SearchProductWithText  -- Returns a list of products based on search keyword
func (pm *ProductManager) SearchProductWithText(text string) []Product {
	results := make([]Product, 0)
	for _, product := range pm.products {
		if strings.Contains(strings.ToLower(product.pname), strings.ToLower(text)) ||
			strings.Contains(strings.ToLower(product.ptype), strings.ToLower(text)) ||
			strings.Contains(strings.ToLower(product.prod_place), strings.ToLower(text)) {
			results = append(results, product)
		}
	}
	return results
}

// SearchResultsWithPlace -- Returns a list of product available at that specific place
func (pm *ProductManager) SearchResultsWithPlace(place string) []Product {
	results := make([]Product, 0)
	for _, product := range pm.products {
		if strings.EqualFold(product.prod_place, place) {
			results = append(results, product)
		}
	}
	return results
}

// SearchResultsOutofWarranty -- Returns list of products which are out of warranty
func (pm *ProductManager) SearchResultsOutofWarranty(year int) []Product {
	results := make([]Product, 0)
	for _, product := range pm.products {
		if product.warranty < year {
			results = append(results, product)
		}
	}
	return results
}

func main() {
	pm := NewProductManager()

	// Adding  to list
	product1 := Product{"Chair", "Furniture", "Mumbai", 2024}
	product2 := Product{"VoltasAC", "AirConditioner", "Chennai", 2022}
	product3 := Product{"VoltasAC", "AirConditioner", "Kolkata", 2022}
	pm.AddProduct(product1)
	pm.AddProduct(product2)
	pm.AddProduct(product3)

	//Invoke the GetAllProducts method
	products := pm.GetAllProducts()
	fmt.Println("Printing  All Products in one line", products)

	//Pritnts  one by one the Product from list of products
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
	product := pm.GetProduct("Product 1")
	if product != nil {
		fmt.Println("Returned Selected Prodcut", product)
		fmt.Printf("Product Name: %s\n", product.pname)
		fmt.Printf("Product Type: %s\n", product.ptype)
		fmt.Printf("Product Place: %s\n", product.prod_place)
		fmt.Printf("Product Warranty: %d\n", product.warranty)
		fmt.Println("-----------")

	}

	//SearchProductWithText
	fmt.Println("=============Search Products for a Particular text================")
	fmt.Println("Enter what you want to search with out spaces")
	var searchProduct string
	fmt.Scan(&searchProduct)
	searchResults := pm.SearchProductWithText(searchProduct)
	//	searchResults := pm.SearchProductWithText("Chair")
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

	fmt.Println("=============Search products  at a Place  ================")
	fmt.Println("Enter what you want to search")
	var searchPlace string
	fmt.Scan(&searchPlace)
	searchResultsByPlace := pm.SearchResultsWithPlace(searchPlace)

	//	searchResultsByPlace := pm.SearchResultsWithPlace("Chennai")
	if len(searchResultsByPlace) > 0 {
		for _, product := range searchResultsByPlace {
			fmt.Printf("Product Name: %s\n", product.pname)
			fmt.Printf("Product Type: %s\n", product.ptype)
			fmt.Printf("Product Place: %s\n", product.prod_place)
			fmt.Printf("Product Warranty: %d\n", product.warranty)
			fmt.Println("-----------")
		}

	} else {
		fmt.Println("No products found in the specified place ")
	}

	fmt.Println("=============Search for  products  out of warranty ================")
	fmt.Println("Enter current Year")
	var searchWaranty int
	fmt.Scan(&searchWaranty)
	searchResultsByWarranty := pm.SearchResultsOutofWarranty(searchWaranty)
	if len(searchResultsByWarranty) > 0 {
		for _, product := range searchResultsByWarranty {
			fmt.Println("products out of warranty")
			fmt.Printf("Product Name: %s\n", product.pname)
			fmt.Printf("Product Type: %s\n", product.ptype)
			fmt.Printf("Product Place: %s\n", product.prod_place)
			fmt.Printf("Product Warranty: %d\n", product.warranty)
			fmt.Println("-----------")
		}

	} else {
		fmt.Println("All products found are in warranty range")
	}

}
