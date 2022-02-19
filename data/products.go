package data

import "time"

//Product define
type Product struct {
	ID          int
	Price       int
	Name        string
	Description string
	CreateOn    string
	UpdateOn    string
	DeleteOn    string
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Price:       320,
		Name:        "Latte",
		Description: "Milky coffee",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Price:       400,
		Name:        "Sandwich",
		Description: "Very tasty food",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
