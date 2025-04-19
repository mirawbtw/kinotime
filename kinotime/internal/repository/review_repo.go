package repository

import (
	"context"
	"database/sql"
	"kinotime/internal/models"
)

type ReviewRepository struct {
	DB *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{DB: db}
}

func (repo *ReviewRepository) CreateReview(ctx context.Context, userID, movieID, rating int, comment string) error {
	_, err := repo.DB.ExecContext(ctx, "INSERT INTO reviews (user_id, movie_id, rating, comment) VALUES ($1, $2, $3, $4)", userID, movieID, rating, comment)
	return err
}

func (repo *ReviewRepository) GetReviewByID(ctx context.Context, id int) (*models.Review, error) {
	var review models.Review
	err := repo.DB.QueryRowContext(ctx, "SELECT id, user_id, movie_id, rating, comment FROM reviews WHERE id = $1", id).
		Scan(&review.ID, &review.UserID, &review.MovieID, &review.Rating, &review.Comment)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (repo *ReviewRepository) GetReviewsByMovieID(ctx context.Context, movieID int) ([]models.Review, error) {
	rows, err := repo.DB.QueryContext(ctx, "SELECT id, user_id, movie_id, rating, comment FROM reviews WHERE movie_id = $1", movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.UserID, &review.MovieID, &review.Rating, &review.Comment); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func (repo *ReviewRepository) UpdateReview(ctx context.Context, id, rating int, comment string) error {
	_, err := repo.DB.ExecContext(ctx, "UPDATE reviews SET rating = $1, comment = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3", rating, comment, id)
	return err
}

func (repo *ReviewRepository) DeleteReview(ctx context.Context, id int) error {
	_, err := repo.DB.ExecContext(ctx, "DELETE FROM reviews WHERE id = $1", id)
	return err
}
