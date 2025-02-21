package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	description string `json:"description"`
	Completed   bool   `json:"completed"`
}
