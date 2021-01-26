package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// Store contains functions for working with the database
type Store struct {
	db *gorm.DB
}

func generateDSN() string {
	return fmt.Sprintf(
		"%s:%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PROTOCOL"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_NAME"),
	)
}

// NewStore creates a new Store struct
func NewStore() (*Store, error) {
	dsn := generateDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Repository{})
	return &Store{db}, nil
}

// Create
func (s Store) Create(repos []Repository) {
	s.db.Create(&repos)
}

// Retrieve
func (s Store) Retrieve() []Repository {
	repos := []Repository{}
	s.db.Find(&repos)
	return repos
}
