package entities

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}
