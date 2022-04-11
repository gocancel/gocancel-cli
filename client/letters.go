package client

import (
	"context"

	"github.com/gocancel/gocancel-go"
)

// LettersService is the interface that wraps the gocancel LettersService.
type LettersService interface {
	List() ([]*gocancel.Letter, error)
	Create(request *gocancel.LetterRequest) (*gocancel.Letter, error)
	Get(letterID string) (*gocancel.Letter, error)
}

type lettersService struct {
	client *gocancel.Client
	ctx    context.Context
}

// NewLettersService builds an instance of LettersService.
func NewLettersService(client *gocancel.Client) LettersService {
	return &lettersService{
		client: client,
		ctx:    context.Background(),
	}
}

func (s *lettersService) List() ([]*gocancel.Letter, error) {
	letters, _, err := s.client.Letters.List(s.ctx, nil)
	if err != nil {
		return nil, err
	}

	return letters, nil
}

func (s *lettersService) Create(request *gocancel.LetterRequest) (*gocancel.Letter, error) {
	letter, _, err := s.client.Letters.Create(s.ctx, request)
	if err != nil {
		return nil, err
	}

	return letter, nil
}

func (s *lettersService) Get(letterID string) (*gocancel.Letter, error) {
	letter, _, err := s.client.Letters.Get(s.ctx, letterID)
	if err != nil {
		return nil, err
	}

	return letter, nil
}
