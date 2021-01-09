package main

import (
	"testing"

	"github.com/PetrNovak96/godemo/client/client"
	"github.com/PetrNovak96/godemo/client/client/products"
)

func TestClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	_, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}
}
