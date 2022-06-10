package postgres

import (
	"database/sql"
	"fmt"
	"log"

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

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("can't do ping %v", err)
	}

	log.Println("Connected to postgres!")
	return db, nil
}
