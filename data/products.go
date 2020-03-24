package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

type Products []*Product


func(p *Products) ToJSON(w io.Writer) error{
	err := json.NewEncoder(w)
	return err.Encode(p)
	
}
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID: 1,
		Name: "string",
		Description: "string",
		Price: 32,
		SKU: "string",
		CreatedOn: "string",
		UpdatedOn: "string",
	},
	&Product{
		ID: 2,
		Name: "string",
		Description: "string",
		Price: 32,
		SKU: "string",
		CreatedOn: "string",
		UpdatedOn: "string",
	},

}
