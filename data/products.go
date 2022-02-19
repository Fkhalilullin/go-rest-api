package data

import (
	"encoding/json"
	"io"
	"time"
)

//Product define
type Product struct {
	ID          int    `json:"id"`
	Price       int    `jsos:"price"`
	Name        string `jsos:"name"`
	Description string `jsos:"descrition"`
	CreatedOn   string `jsos:"-"`
	UpdatedOn   string `jsos:"-"`
	DeletedOn   string `jsos:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = Products{
	&Product{
		ID:          1,
		Price:       200,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Price:       300,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
