package main

import "fmt"

type observer interface {
	update()
}

type stock struct {
	symbol    string
	price     float32
	observers []observer
}

func (s *stock) setObserver(observer observer) {
	s.observers = append(s.observers, observer)
}

func (s stock) notifyObservers() {
	for _, obs := range s.observers {
		obs.update()
	}
}

func (s *stock) setPrice(price float32) {
	s.price = price
	s.notifyObservers()
}

type statusBar struct {
	stocks []stock
}

func (s *statusBar) addStock(stock *stock) {
	s.stocks = append(s.stocks, *stock)
	stock.setObserver(s)
}

func (s statusBar) show() {
	for _, stock := range s.stocks {
		fmt.Printf("stock: %v\n", stock)
	}
}

func (s *statusBar) update() {
	fmt.Println("Price Changed - Refreshing StatusBar")
}

type stockListView struct {
	stocks []stock
}

func (s *stockListView) addStock(stock *stock) {
	s.stocks = append(s.stocks, *stock)
	stock.setObserver(s)
}

func (s stockListView) show() {
	for _, stock := range s.stocks {
		fmt.Printf("stock: %v\n", stock)
	}
}

func (s *stockListView) update() {
	fmt.Println("Price Changed - Refreshing ListView")
}
