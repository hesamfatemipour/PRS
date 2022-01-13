package main

type Product struct {
	id         int
	name       string
	attributes []string
}
type ShoppingCart map[int]Product
