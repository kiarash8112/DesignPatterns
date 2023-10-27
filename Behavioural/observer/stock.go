package observer

import "fmt"

/*
We’re building an application for watching the price of various
stocks. There are two places in our application where we need to
display the stocks:
- StatusBar: shows the popular stocks
- StockListView: shows the complete list of stocks
When the price of a stock changes, the corresponding views
(StatusBar and/or StockListView) need to be refreshed to reflect the
latest price.
Our application currently does not have the ability to communicate
the change in stock prices to the corresponding views.
Use the observer pattern to solve this problem.
*/

type stock struct {
	symbol string
	price  float32
}

type statusBar struct {
	stocks []stock
}

func (s *statusBar) addStock(stock stock) {
	s.stocks = append(s.stocks, stock)
}

func (s statusBar) show() {
	for _, stock := range s.stocks {
		fmt.Printf("stock: %v\n", stock)
	}
}

type stockListView struct {
	stocks []stock
}

func (s *stockListView) addStock(stock stock) {
	s.stocks = append(s.stocks, stock)
}

func (s stockListView) show() {
	for _, stock := range s.stocks {
		fmt.Printf("stock: %v\n", stock)
	}
}
