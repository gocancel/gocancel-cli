package client

import (
	"context"

	"github.com/gocancel/gocancel-go"
)

// OrganizationsService is the interface that wraps the gocancel OrganizationsService.
type OrganizationsService interface {
	List(opts *gocancel.OrganizationsListOptions) ([]*gocancel.Organization, error)
	Get(organizationID string) (*gocancel.Organization, error)
	ListProducts(organizationID string, opts *gocancel.OrganizationProductsListOptions) ([]*gocancel.Product, error)
	GetProduct(organizationID string, productID string) (*gocancel.Product, error)
}

type organizationsService struct {
	client *gocancel.Client
	ctx    context.Context
}

// NewOrganizationsService builds an instance of OrganizationsService.
func NewOrganizationsService(client *gocancel.Client) OrganizationsService {
	return &organizationsService{
		client: client,
		ctx:    context.Background(),
	}
}

func (s *organizationsService) List(opts *gocancel.OrganizationsListOptions) ([]*gocancel.Organization, error) {
	organizations, _, err := s.client.Organizations.List(s.ctx, opts)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

func (s *organizationsService) Get(organizationID string) (*gocancel.Organization, error) {
	organization, _, err := s.client.Organizations.Get(s.ctx, organizationID)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *organizationsService) ListProducts(organizationID string, opts *gocancel.OrganizationProductsListOptions) ([]*gocancel.Product, error) {
	products, _, err := s.client.Organizations.ListProducts(s.ctx, organizationID, opts)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *organizationsService) GetProduct(organizationID string, productID string) (*gocancel.Product, error) {
	product, _, err := s.client.Organizations.GetProduct(s.ctx, organizationID, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
