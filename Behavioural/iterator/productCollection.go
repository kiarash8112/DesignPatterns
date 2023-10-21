package iterator

/*
This class only allows us to add a product to a collection. Once we
add a bunch of products to a collection, there is no way to iterate
that collection and print the products.
Implement this feature using the iterator pattern.
*/

type Product struct {
	id   int
	name string
}

type ProductCollection struct {
	products []Product
}

func (p *ProductCollection) add(product Product) {
	p.products = append(p.products, product)
}
