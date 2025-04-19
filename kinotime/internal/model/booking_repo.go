package model

import (
	"context"
	"database/sql"

	"kinotime/internal/types"
)

type BookingRepository struct {
	DB *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{DB: db}
}

func (repo *BookingRepository) CreateBooking(ctx context.Context, userID, movieID, seatsBooked int, totalPrice float64, status string) error {
	_, err := repo.DB.ExecContext(ctx, "INSERT INTO bookings (user_id, movie_id, seats_booked, total_price, status) VALUES ($1, $2, $3, $4, $5)", userID, movieID, seatsBooked, totalPrice, status)
	return err
}

func (repo *BookingRepository) GetBookingByID(ctx context.Context, id int) (*types.Booking, error) {
	var booking types.Booking
	err := repo.DB.QueryRowContext(ctx, "SELECT id, user_id, movie_id, seats_booked, total_price, status FROM bookings WHERE id = $1", id).
		Scan(&booking.ID, &booking.UserID, &booking.MovieID, &booking.SeatsBooked, &booking.TotalPrice, &booking.Status)
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (repo *BookingRepository) GetBookingsByUserID(ctx context.Context, userID int) ([]types.Booking, error) {
	rows, err := repo.DB.QueryContext(ctx, "SELECT id, user_id, movie_id, seats_booked, total_price, status FROM bookings WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []types.Booking
	for rows.Next() {
		var booking types.Booking
		if err := rows.Scan(&booking.ID, &booking.UserID, &booking.MovieID, &booking.SeatsBooked, &booking.TotalPrice, &booking.Status); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}

func (repo *BookingRepository) GetAllBookings(ctx context.Context) ([]types.Booking, error) {
	rows, err := repo.DB.QueryContext(ctx, "SELECT id, user_id, movie_id, seats_booked, total_price, status FROM bookings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []types.Booking
	for rows.Next() {
		var booking types.Booking
		if err := rows.Scan(&booking.ID, &booking.UserID, &booking.MovieID, &booking.SeatsBooked, &booking.TotalPrice, &booking.Status); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}

func (repo *BookingRepository) UpdateBooking(ctx context.Context, id int, seatsBooked int, totalPrice float64, status string) error {
	_, err := repo.DB.ExecContext(ctx, "UPDATE bookings SET seats_booked = $1, total_price = $2, status = $3 WHERE id = $4", seatsBooked, totalPrice, status, id)
	return err
}

func (repo *BookingRepository) DeleteBooking(ctx context.Context, id int) error {
	_, err := repo.DB.ExecContext(ctx, "DELETE FROM bookings WHERE id = $1", id)
	return err
}
