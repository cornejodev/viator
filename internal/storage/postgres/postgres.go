package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cornejodev/viator/config"
	_ "github.com/lib/pq"
)

func New(dbcfg config.Database) (db *sql.DB, err error) {
	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		dbcfg.User, dbcfg.Name, dbcfg.Password)

	db, err = sql.Open(dbcfg.Engine, conn)
	if err != nil {
		return nil, fmt.Errorf("can't open the data base %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("can't do ping %v", err)
	}

	return db, nil
}
