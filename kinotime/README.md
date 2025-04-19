### Prerequisites
- **Go**: Ensure Go 1.23 or later is installed.
- **PostgreSQL**: Install PostgreSQL and set up a database (`users`).
- **.env**: Configure the `.env` file with necessary variables.
- **Migrations**: Run SQL migration files in the `migrations` directory to set up tables. Connect to the database using the following command:

```bash
PGPASSWORD=999000 psql -U postgres -d users -h localhost -p 5432
```

Once connected, run the migration scripts to create tables like `movies`, `users`, `bookings`, etc., to match your application's needs. You can run SQL scripts in the `migrations` directory to set up the necessary tables.

### Setup
1. **Clone the Repository**: `git clone <repo-url>`
2. **Install Dependencies**: Run `go mod tidy` to download necessary packages.
3. **Run Migrations**: Use `psql` or a migration tool to apply the SQL files.
4. **Run the Application**: Use `make run` to start the server.

### `go mod tidy`
- **Purpose**: Cleans up `go.mod` by removing unused dependencies and adding missing ones.
- **Usage**: Run `go mod tidy` to keep the module dependencies up to date.

### Running the Application
1. Start the app with `make run`.
2. The server listens on the port specified in `.env` (`8080` by default).