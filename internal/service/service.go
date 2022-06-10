package service

import (
	"fmt"

	"github.com/cornejodev/viator/internal/storage"
)

type Service struct {
	Storage storage.Storage
	Demo    DemoService
}

func New(s storage.Storage) (*Service, error) {
	r, err := s.ProvideRepository()
	if err != nil {
		return nil, fmt.Errorf("error from storage: %v", err)
	}

	return &Service{
		Demo: NewDemoService(r.Demo),
	}, nil
}
