// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Order struct {
	ID       string     `json:"id"`
	Products []*Product `json:"products"`
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
