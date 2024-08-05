package store

import (
	"database/sql"
	"fmt"
	"log"
)

type Store struct {
	db *sql.DB
}

func (s *Store) GetDB() *sql.DB {
	return s.db
}

func NewStore(db *sql.DB, setupFuncs []func(*sql.DB) error) *Store {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Db connection sucessful....")
	for _, setupFunc := range setupFuncs {
		if err := setupFunc(db); err != nil {
			log.Fatal("failed in setup: ", err)
		}
	}
	fmt.Println("Db setup sucessful....")
	fmt.Println("new store created")

	store := Store{
		db: db,
	}
	return &store
}
