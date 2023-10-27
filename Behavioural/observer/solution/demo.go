package main

func main() {
	statusBar := statusBar{}
	listView := stockListView{}

	stock1 := stock{symbol: "stock1", price: 10}
	stock2 := stock{symbol: "stock2", price: 20}
	stock3 := stock{symbol: "stock3", price: 30}

	statusBar.addStock(&stock1)
	statusBar.addStock(&stock2)

	listView.addStock(&stock1)
	listView.addStock(&stock2)
	listView.addStock(&stock3)

	stock2.setPrice(12)
	stock3.setPrice(13)
}
