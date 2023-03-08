package entity

import (
	"errors"
	"fmt"
)

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64, finalPrice float64) (*Order, error) {
	order := &Order{
		ID:         id,
		Price:      price,
		Tax:        tax,
		FinalPrice: finalPrice,
	}

	err := order.Validate()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) getTotal() float64 {
	return o.Price * float64(o.Tax)
}

func (o *Order) setPrice(price float64) float64 {
	o.Price = price
	fmt.Println("Price dentro do setPrice: ", o.Price)
	return o.Price
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("ID is required")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price * o.Tax

	err := o.Validate()

	if err != nil {
		return err
	}

	return nil
}
