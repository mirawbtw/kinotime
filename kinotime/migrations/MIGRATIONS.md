# MIGRATIONS

## RUNNING

    ``` bash
    psql -U username -d database_name -a -f migrations/001_create_users_table.sql
    psql -U username -d database_name -a -f migrations/002_create_movies_table.sql
    psql -U username -d database_name -a -f migrations/003_create_bookings_table.sql
    psql -U username -d database_name -a -f migrations/004_create_reviews_table.sql
    psql -U username -d database_name -a -f migrations/005_create_indexes.sql
    psql -U username -d database_name -a -f migrations/006_create_triggers.sql

    ```

## INFO

Here’s a short description for each migration file:

---

### 1. **`001_create_users_table.sql`**

**Description**:  
This migration creates the `users` table with the following columns:
- `id`: Primary key, auto-incremented.
- `username`: Unique, required string.
- `password`: Required string.
- `created_at`: Timestamp when the user is created.
- `updated_at`: Timestamp when the user record is last updated.

---

### 2. **`002_create_movies_table.sql`**

**Description**:  
This migration creates the `movies` table with the following columns:
- `id`: Primary key, auto-incremented.
- `title`: Required string representing the movie title.
- `genre`: Required string representing the movie genre.
- `description`: Optional text field for additional movie details.
- `release_date`: Optional date field for the movie release date.
- `created_at`: Timestamp when the movie is created.
- `updated_at`: Timestamp when the movie record is last updated.

---

### 3. **`003_create_bookings_table.sql`**

**Description**:  
This migration creates the `bookings` table with the following columns:
- `id`: Primary key, auto-incremented.
- `user_id`: Foreign key to the `users` table.
- `movie_id`: Foreign key to the `movies` table.
- `seats_booked`: Integer representing the number of booked seats.
- `total_price`: Decimal representing the total price for the booking.
- `status`: Default value 'pending', representing the booking status.
- `booking_time`: Timestamp when the booking is made.
- `created_at`: Timestamp when the booking is created.
- `updated_at`: Timestamp when the booking record is last updated.

---

### 4. **`004_create_reviews_table.sql`**

**Description**:  
This migration creates the `reviews` table with the following columns:
- `id`: Primary key, auto-incremented.
- `user_id`: Foreign key to the `users` table.
- `movie_id`: Foreign key to the `movies` table.
- `rating`: Integer between 1 and 5 representing the user's rating.
- `comment`: Optional text field for additional review comments.
- `created_at`: Timestamp when the review is created.
- `updated_at`: Timestamp when the review record is last updated.

---

### 5. **`005_create_indexes.sql`**

**Description**:  
This migration creates indexes on the frequently queried fields in each table:
- `users.username`: Index to speed up lookups based on the username.
- `movies.title`: Index to speed up lookups based on the movie title.
- `bookings.user_id`: Index to speed up lookups for bookings by a user.
- `bookings.movie_id`: Index to speed up lookups for bookings by movie.
- `reviews.user_id`: Index to speed up lookups for reviews by user.
- `reviews.movie_id`: Index to speed up lookups for reviews by movie.

---

### 6. **`006_create_triggers.sql`**

**Description**:  
This migration creates triggers to automatically update the `updated_at` timestamp whenever a row in any table (`users`, `movies`, `bookings`, `reviews`) is updated. The `update_timestamp` function is called before an update, ensuring the `updated_at` field is always accurate.
