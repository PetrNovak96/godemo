// Package classification of Product API
//
// Documentation for Product API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//  swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PetrNovak96/godemo/product-api/data"
	"github.com/gorilla/mux"
)

// A list of products returned in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// A list of products
	// in:body
	Body []data.Product
}

// A product
// swagger:response productResponse
type productResponseWrapper struct {
	// A product
	// in:body
	Body data.Product
}

// swagger:parameters updateProduct
type productIDParameterWrapper struct {
	// The id of the product to update
	// in: path
	// required: true
	ID int `json:"id"`
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	p.l.Println("volani")
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error writing product list", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

// swagger:route PUT /products{id} products updateProduct
// Updates a product
// responses:
//	200: productResponse
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Error converting id", http.StatusBadRequest)
		return
	}
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
