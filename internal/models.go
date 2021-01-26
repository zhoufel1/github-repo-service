package main

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
