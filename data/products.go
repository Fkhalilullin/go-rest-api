package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

//Product define
type Product struct {
	ID          int    `json:"id"`
	Price       int    `jsos:"price"`
	Name        string `jsos:"name"`
	Description string `jsos:"description"`
	CreatedOn   string `jsos:"created"`
	UpdatedOn   string `jsos:"updated"`
	DeletedOn   string `jsos:"deleted"`
}

type Products []*Product

var ErrProductNotFound = fmt.Errorf("Product not found")

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)

	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = Products{
	{
		ID:          1,
		Price:       250,
		Name:        "Eugene Onegin",
		Description: "Novel in verse written by Alexander Pushkin",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Price:       300,
		Name:        "War and Peace",
		Description: "Literary work mixed with chapters on history and philosophy by the Russian author Leo Tolstoy",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
