package models

type BookingRequest struct {
	MovieID     int     `json:"movie_id"`
	SeatsBooked int     `json:"seats_booked"`
	TotalPrice  float64 `json:"total_price"`
	Status      string  `json:"status"`
	BookingTime string  `json:"booking_time"`
}
