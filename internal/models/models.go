package models

// Repository contains relevant data from Github repo API response
type Repository struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	FullName    string  `json:"full_name"`
	Description *string `json:"description"`
	URL         string  `json:"html_url"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	PushedAt    string  `json:"pushed_at"`
	Language    *string `json:"language"`
}

// InitCache caches whether or not an initial fetch is done
type InitCache struct {
	IsInitialized bool
}

type ResponseJSON struct {
	Code    uint
	Message string
}
