# Go Bookstore API

This is a simple RESTful API for managing books, built with Go, GORM, Gorilla Mux, and MySQL.

## Features
- Add, view, update, and delete books
- MySQL database integration
- RESTful endpoints

## Endpoints
- `GET /books` - List all books
- `GET /books/{bookId}` - Get a book by ID
- `POST /books` - Add a new book
- `PUT /books/{bookId}` - Update a book
- `DELETE /books/{bookId}` - Delete a book

## Setup
1. Clone the repository
2. Configure your MySQL credentials in `pkg/config/app.go`
3. Run `go mod tidy` to install dependencies
4. Build and run the project:
   ```sh
   go run cmd/main/main.go
   ```

## Example Book JSON
```
{
  "name": "The Go Programming Language",
  "author": "Alan A. A. Donovan",
  "publication": "Addison-Wesley"
}
```

## License
MIT
