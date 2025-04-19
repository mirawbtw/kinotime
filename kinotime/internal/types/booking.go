package types

type Booking struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	MovieID     int     `json:"movie_id"`
	SeatsBooked int     `json:"seats_booked"`
	TotalPrice  float64 `json:"total_price"`
	Status      string  `json:"status"`
	BookingTime string  `json:"booking_time"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
