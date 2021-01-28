package store

import (
	"errors"
	"fmt"
	"github-repo-service/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// RepoDB handler for database operations
type RepoDB struct {
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

// NewRepoDB create a new RepoDB
func NewRepoDB() (*RepoDB, error) {
	dsn := generateDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Repository{}, &models.InitCache{})

	initCache := models.InitCache{}
	err = db.First(&initCache).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		db.Create(&models.InitCache{IsInitialized: false})
		fmt.Println(initCache.IsInitialized)
	}
	if !initCache.IsInitialized {
		return &RepoDB{db: db, IsInitialized: false}, nil
	}
	return &RepoDB{db: db, IsInitialized: true}, nil
}

// Initialize store data from initial fetch
func (d RepoDB) Initialize(repos []models.Repository) {
	var fetched models.InitCache
	d.db.First(&fetched)
	d.db.Model(&fetched).Where("is_initialized = ?", false).Update("is_initialized", true)
	d.db.Create(&repos)
}

// RetrieveRepos retrieve data from store
func (d RepoDB) RetrieveRepos() []models.Repository {
	repos := []models.Repository{}
	d.db.Find(&repos)
	return repos
}

// UpdateRepos update database entries
func (d RepoDB) UpdateRepos(new []models.Repository) {
	current := []models.Repository{}
	d.db.Find(&current)
	d.db.Delete(&current)
	d.db.Create(&new)
}
