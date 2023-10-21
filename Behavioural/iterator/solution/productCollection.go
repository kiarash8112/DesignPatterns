package solution

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

//if type of products changed only this file need change not main one

type iterator interface {
	next()
	hasNext() bool
	current() Product
}

type ProductIterator struct {
	currentIndex int
	products     *[]Product
}

func newProductIterator(p ProductCollection) ProductIterator {
	return ProductIterator{
		currentIndex: 0,
		products:     &p.products,
	}
}

func (p *ProductIterator) next() {
	p.currentIndex++
}

func (p *ProductIterator) current() Product {
	products := *p.products
	return products[p.currentIndex]
}

func (p *ProductIterator) hasNext() bool {
	products := *p.products
	return p.currentIndex < len(products)
}
