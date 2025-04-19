package repository

import (
	"context"
	"database/sql"
	"strings"

	"kinotime/internal/models"
)

type MovieRepository struct {
	DB *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (repo *MovieRepository) CreateMovie(ctx context.Context, title, posterUrl, genre, description string, year int, actors []string) error {
	_, err := repo.DB.ExecContext(ctx, `
		INSERT INTO movies (title, poster_url, genre, description, year, actors) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		title, posterUrl, genre, description, year, strings.Join(actors, ","),
	)
	return err
}

func (repo *MovieRepository) GetMovieByID(ctx context.Context, id int) (*models.Movie, error) {
	var movie models.Movie
	var actorsStr string

	err := repo.DB.QueryRowContext(ctx, `
		SELECT id, title, poster_url, genre, description, year, actors 
		FROM movies WHERE id = $1`, id).
		Scan(&movie.ID, &movie.Title, &movie.PosterUrl, &movie.Genre, &movie.Description, &movie.Year, &actorsStr)

	if err != nil {
		return nil, err
	}

	movie.Actors = strings.Split(actorsStr, ",")
	return &movie, nil
}

func (repo *MovieRepository) GetAllMovies(ctx context.Context) ([]models.Movie, error) {
	rows, err := repo.DB.QueryContext(ctx, `SELECT id, title, poster_url, genre, description, year, actors FROM movies`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		var actorsStr string

		if err := rows.Scan(&movie.ID, &movie.Title, &movie.PosterUrl, &movie.Genre, &movie.Description, &movie.Year, &actorsStr); err != nil {
			return nil, err
		}

		movie.Actors = strings.Split(actorsStr, ",")
		movies = append(movies, movie)
	}
	return movies, nil
}

func (repo *MovieRepository) UpdateMovie(ctx context.Context, id int, title, posterUrl, genre, description string, year int, actors []string) error {
	_, err := repo.DB.ExecContext(ctx, `
		UPDATE movies SET title = $1, poster_url = $2, genre = $3, description = $4, year = $5, actors = $6 
		WHERE id = $7`,
		title, posterUrl, genre, description, year, strings.Join(actors, ","), id,
	)
	return err
}

func (repo *MovieRepository) DeleteMovie(ctx context.Context, id int) error {
	_, err := repo.DB.ExecContext(ctx, "DELETE FROM movies WHERE id = $1", id)
	return err
}
