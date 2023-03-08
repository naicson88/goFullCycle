package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Order struct {
	ID       string
	Price    float64
	Quantity int
}

func (o *Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

func (o *Order) setPrice(price float64) float64 {
	o.Price = price
	fmt.Println("Price dentro do setPrice: ", o.Price)
	return o.Price
}

func NewOrder() *Order {
	return &Order{}
}

func main() {
	fmt.Println(uuid.New().String())
}
