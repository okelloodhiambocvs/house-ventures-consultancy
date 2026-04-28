package models

type CV struct {
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Skills     []string `json:"skills"`
	Experience []string `json:"experience"`
	Education  []string `json:"education"`
}