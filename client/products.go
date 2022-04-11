package client

import (
	"context"

	"github.com/gocancel/gocancel-go"
)

// ProductsService is the interface that wraps the gocancel ProductsService.
type ProductsService interface {
	Get(productID string) (*gocancel.Product, error)
}

type productsService struct {
	client *gocancel.Client
	ctx    context.Context
}

// NewProductsService builds an instance of ProductsService.
func NewProductsService(client *gocancel.Client) ProductsService {
	return &productsService{
		client: client,
		ctx:    context.Background(),
	}
}

func (s *productsService) Get(productID string) (*gocancel.Product, error) {
	product, _, err := s.client.Products.Get(s.ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
