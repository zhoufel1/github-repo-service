package store

import (
	"errors"
	"fmt"
	"github-repo-service/internal/models"
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
	db.AutoMigrate(&models.Repository{}, &models.InitialFetchCheck{})
	fetched := models.InitialFetchCheck{}
	err = db.First(&fetched).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.Create(&models.InitialFetchCheck{Fetched: false})
	}
	return &Store{db}, nil
}

func (s Store) IsInitiallyFetched() bool {
	fetched := models.InitialFetchCheck{}
	s.db.First(&fetched)
	return fetched.Fetched
}

func (s Store) SetFetched() {
	var fetched models.InitialFetchCheck
	s.db.First(&fetched)
	s.db.Model(&fetched).Where("fetched = ?", false).Update("fetched", true)
}

// Create
func (s Store) Create(repos []models.Repository) {
	s.db.Create(&repos)
}

// Retrieve
func (s Store) Retrieve() []models.Repository {
	repos := []models.Repository{}
	s.db.Find(&repos)
	return repos
}
