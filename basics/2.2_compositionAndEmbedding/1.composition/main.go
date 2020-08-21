package main

import "fmt"

type Product struct {
	Name string
}

type Customer struct {
	Product Product
	Name    string
	URL     string
}

func (c *Customer) GetCustomerAndProductName() string {
	return c.Name + ":" + c.Product.Name
}

func main() {

	//different objects
	p := Product{Name: "iPhone"}
	c := Customer{Product: p, Name: "Apple", URL: "http://www.apple.com"}
	fmt.Println(c.GetCustomerAndProductName())

	//one object
	cc := Customer{
		Product: Product{Name: "F150"},
		Name:    "ford",
		URL:     "http://ford.com",
	}

	fmt.Println(cc.GetCustomerAndProductName())

}
