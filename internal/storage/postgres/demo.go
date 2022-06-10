package postgres

import (
	"database/sql"

	"github.com/cornejodev/viator/internal/domain"
)

// DemoRepository (before DemoDAO) it's implementation of DemoDAO interface of service/.
type DemoRepository struct {
	db *sql.DB
}

func NewDemoRepository(db *sql.DB) *DemoRepository {
	return &DemoRepository{
		db: db,
	}
}

func (r DemoRepository) Create(demo *domain.Demo) error {
	stmt, err := r.db.Prepare("INSERT INTO boilerplates (name) VALUEs ($1) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(demo.Name).Scan(&demo.ID)
	if err != nil {
		return err
	}

	return nil
}
