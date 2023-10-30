package proxy

import "fmt"

type Product struct {
	id   int
	name string
}

type DbContext struct {
	updateObjects map[int]Product
}

func (DbContext) getProduct(id int) Product {
	// Automatically generate SQL statements
	// to read the product with the given ID.
	fmt.Printf("SELECT * FROM products WHERE product_id = %d \n", id)

	// Simulate reading a product object from a database.
	product := Product{id: id}
	return product
}

func (db *DbContext) saveChanges() {
	// Automatically generate SQL statements
	// to update the database.
	for key, updateObj := range db.updateObjects {
		fmt.Printf("UPDATE products SET name = '%s' WHERE product_id = %d \n", updateObj.name, updateObj.id)
		delete(db.updateObjects, key)
	}
}

func (db *DbContext) markAsChanged(product Product) {
	db.updateObjects[product.id] = product
}
