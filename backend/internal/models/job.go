package models

type Job struct {
	Title       string   `json:"title"`
	Company     string   `json:"company"`
	Skills      []string `json:"skills"`
	Description string   `json:"description"`
}