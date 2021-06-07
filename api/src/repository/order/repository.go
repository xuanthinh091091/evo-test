package order

import (
	"candidate/api/src/models"
	"context"
	"database/sql"
	"log"
	"time"
)

// Repository is the interface spec expected of a order repository
type Repository interface {
	OrderBook(ctx context.Context, userID, bookID int64, quantity int, shippingDetais string) error
	GetAllOrders(ctx context.Context) ([]models.Order, error)
}

// PostgresClient is the implementation of order repository for Postgres
type PostgresClient struct {
	dbconn *sql.DB
}

// NewPostgresClient initializes postgres Client Repository
func NewPostgresClient(dbconn *sql.DB) *PostgresClient {
	return &PostgresClient{
		dbconn: dbconn,
	}
}

// OrderBook make order for book
func (c *PostgresClient) OrderBook(ctx context.Context, userID, bookID int64, quantity int, shippingDetais string) error {
	_, err := c.dbconn.QueryContext(ctx, `INSERT INTO orders (users_id, books_id, quantity, shipping_details, created_at) 
	VALUES ($1, $2, $3, $4, $5)`, userID, bookID, quantity, shippingDetais, time.Now())
	if err != nil {
		log.Println("[OrderBook] error: ", err.Error())
		return err
	}
	return nil
}

// GetAllOrders get all books
func (c *PostgresClient) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	var rs *sql.Rows
	var err error

	rs, err = c.dbconn.QueryContext(ctx, `
	select orders.id,users.user_name, books.name, orders.quantity,orders.shipping_details,orders.created_at from orders 
	join users on  orders.users_id = users.id
	join books on  orders.books_id = books.id 
	`)

	if err != nil {
		log.Println("[GetAllOrders] query error: ", err.Error())
		return nil, err
	}
	defer rs.Close()

	orders := make([]models.Order, 0)
	for rs.Next() {
		order := models.Order{}
		err := rs.Scan(&order.ID, &order.UserName, &order.BookName, &order.Quantity, &order.ShippingDetails, &order.CreateAt)
		if err != nil {
			log.Println("[GetAllOrders] scan error: ", err.Error())
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
