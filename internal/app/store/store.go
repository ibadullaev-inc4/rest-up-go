package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config         *Config
	db             *sql.DB
	UserRepository *UserRepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// User ...
func (s *Store) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		store: s,
	}
	return s.UserRepository
}
