package model

import (
	"context"
	"database/sql"

	"kinotime/internal/types"
)

type MovieRepository struct {
	DB *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (repo *MovieRepository) CreateMovie(ctx context.Context, title, genre, description string) error {
	_, err := repo.DB.ExecContext(ctx, "INSERT INTO movies (title, genre, description) VALUES ($1, $2, $3)", title, genre, description)
	return err
}

func (repo *MovieRepository) GetMovieByID(ctx context.Context, id int) (*types.Movie, error) {
	var movie types.Movie
	err := repo.DB.QueryRowContext(ctx, "SELECT id, title, genre, description FROM movies WHERE id = $1", id).
		Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Description)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (repo *MovieRepository) GetAllMovies(ctx context.Context) ([]types.Movie, error) {
	rows, err := repo.DB.QueryContext(ctx, "SELECT id, title, genre, description FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []types.Movie
	for rows.Next() {
		var movie types.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Description); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (repo *MovieRepository) UpdateMovie(ctx context.Context, id int, title, genre, description string) error {
	_, err := repo.DB.ExecContext(ctx, "UPDATE movies SET title = $1, genre = $2, description = $3 WHERE id = $4", title, genre, description, id)
	return err
}

func (repo *MovieRepository) DeleteMovie(ctx context.Context, id int) error {
	_, err := repo.DB.ExecContext(ctx, "DELETE FROM movies WHERE id = $1", id)
	return err
}
