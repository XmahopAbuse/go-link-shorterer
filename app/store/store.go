package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	databaseURL string
	db *sql.DB
	linkRepository *LinkRepository
}

func New(databaseURL string) *Store {
	return &Store{
		databaseURL: databaseURL,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.databaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Link() *LinkRepository {
	if s.linkRepository != nil {
		return s.linkRepository
	}

	s.linkRepository = &LinkRepository{store: s}

	return s.linkRepository
}
