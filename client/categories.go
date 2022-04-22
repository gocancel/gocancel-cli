package client

import (
	"context"

	"github.com/gocancel/gocancel-go"
)

// CategoriesService is the interface that wraps the gocancel CategoriesService.
type CategoriesService interface {
	List(opts *gocancel.CategoriesListOptions) ([]*gocancel.Category, error)
	Get(categoryID string) (*gocancel.Category, error)
}

type categoriesService struct {
	client *gocancel.Client
	ctx    context.Context
}

// NewCategoriesService builds an instance of CategoriesService.
func NewCategoriesService(client *gocancel.Client) CategoriesService {
	return &categoriesService{
		client: client,
		ctx:    context.Background(),
	}
}

func (s *categoriesService) List(opts *gocancel.CategoriesListOptions) ([]*gocancel.Category, error) {
	categories, _, err := s.client.Categories.List(s.ctx, opts)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *categoriesService) Get(categoryID string) (*gocancel.Category, error) {
	category, _, err := s.client.Categories.Get(s.ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return category, nil
}
