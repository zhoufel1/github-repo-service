package store

import (
	"errors"
	"fmt"
	"github-repo-service/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// ReposDB contains functions for working with the database
type ReposDB struct {
	db            *gorm.DB
	IsInitialized bool
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

// NewReposDB creates a new Store struct
func NewReposDB() (*ReposDB, error) {
	dsn := generateDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Repository{}, &models.InitializationCache{})

	initCache := models.InitializationCache{}
	err = db.First(&initCache).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		db.Create(&models.InitializationCache{IsInitialized: false})
		return &ReposDB{db: db, IsInitialized: false}, nil
	}
	if !initCache.IsInitialized {
		return &ReposDB{db: db, IsInitialized: false}, nil
	}
	return &ReposDB{db: db, IsInitialized: true}, nil
}

func (d ReposDB) Initialize(repos []models.Repository) {
	var fetched models.InitializationCache
	d.db.First(&fetched)
	d.db.Model(&fetched).Where("is_initialized = ?", false).Update("is_initialized", true)
	d.db.Create(&repos)
}

// RetrieveRepos
func (d ReposDB) RetrieveRepos() []models.Repository {
	repos := []models.Repository{}
	d.db.Find(&repos)
	return repos
}

func (d ReposDB) UpdateRepos(new []models.Repository) {
	current := []models.Repository{}
	d.db.Find(&current)
	d.db.Delete(&current)
	d.db.Create(&new)
}
