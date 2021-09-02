package main

import "fmt"

type Product struct {
	Name string
	Price int
	ID int
}

type Parcel struct {
	Pdt *Product
	ShippedTime time.Time
	Delivered time.Time
}

func (c *Courier) SendProduct(p *Product) *Parcel {
	p := &Parcel{ShippedTime} 
} 