package book

import (
	"candidate/api/src/models"
	"context"
	"database/sql"
	"log"
)

// Repository is the interface spec expected of a book repository
type Repository interface {
	GetBooksAvailableForSale(ctx context.Context) ([]models.Book, error)
	GetAllBooks(ctx context.Context) ([]models.Book, error)
}

// PostgresClient is the implementation of book repository for Postgres
type PostgresClient struct {
	dbconn *sql.DB
}

// NewPostgresClient initializes postgres Client Repository
func NewPostgresClient(dbconn *sql.DB) *PostgresClient {
	return &PostgresClient{
		dbconn: dbconn,
	}
}

// GetBooksAvailableForSale get books available for sale
func (c *PostgresClient) GetBooksAvailableForSale(ctx context.Context) ([]models.Book, error) {
	var rs *sql.Rows
	var err error

	rs, err = c.dbconn.QueryContext(ctx, "SELECT * FROM books where stock > 0")

	if err != nil {
		log.Println("[GetBookAvailableForSale] query error: ", err.Error())
		return nil, err
	}
	defer rs.Close()

	books := make([]models.Book, 0)
	for rs.Next() {
		book := models.Book{}
		err := rs.Scan(&book.ID, &book.Name, &book.Author, &book.Publisher, &book.TotalPage, &book.PublishedDate, &book.Price, &book.Stock)
		if err != nil {
			log.Println("[GetBookAvailableForSale] scan error: ", err.Error())
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// GetAllBooks get all books
func (c *PostgresClient) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	var rs *sql.Rows
	var err error

	rs, err = c.dbconn.QueryContext(ctx, "SELECT * FROM books")

	if err != nil {
		log.Println("[GetAllBooks] query error: ", err.Error())
		return nil, err
	}
	defer rs.Close()

	books := make([]models.Book, 0)
	for rs.Next() {
		book := models.Book{}
		err := rs.Scan(&book.ID, &book.Name, &book.Author, &book.Publisher, &book.TotalPage, &book.PublishedDate, &book.Price, &book.Stock)
		if err != nil {
			log.Println("[GetAllBooks] scan error: ", err.Error())
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
