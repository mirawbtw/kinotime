package types

type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
}
