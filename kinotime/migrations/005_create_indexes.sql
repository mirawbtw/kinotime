-- OPTIONAL MIGRATION

CREATE INDEX idx_users_username ON users (username);
CREATE INDEX idx_movies_title ON movies (title);
CREATE INDEX idx_bookings_user_id ON bookings (user_id);
CREATE INDEX idx_bookings_movie_id ON bookings (movie_id);
CREATE INDEX idx_reviews_user_id ON reviews (user_id);
CREATE INDEX idx_reviews_movie_id ON reviews (movie_id);