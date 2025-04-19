package models

type Movie struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	PosterUrl   string   `json:poster_url`
	Genre       string   `json:"genre"`
	Description string   `json:"description"`
	Year        int      `json:"year"`
	Actors      []string `json:"actors"`
}
