package service

import (
	"github.com/cornejodev/viator/internal/domain"
	"github.com/cornejodev/viator/internal/storage"
)

// DemoService (before DemoDAO) interface is InterfaceDAO in Pattern DAO.
// http://chuwiki.chuidiang.org/index.php?title=Patr%C3%B3n_DAO
type DemoService interface {
	Add(*domain.Demo) error
}

type demoService struct {
	repo storage.DemoRepository
}

func NewDemoService(repo storage.DemoRepository) DemoService {
	return &demoService{repo}
}

func (ds demoService) Add(demo *domain.Demo) error {
	if !demo.HasName() {
		return domain.ErrDemoHasNoName
	}
	if err := ds.repo.Create(demo); err != nil {
		return err
	}
	return nil
}
