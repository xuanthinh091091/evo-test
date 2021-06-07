package user

import (
	"candidate/api/src/models"
	"context"
	"database/sql"
	"errors"
	"log"
)

// Repository is the interface spec expected of a user repository
type Repository interface {
	GetUser(ctx context.Context, userName, passWord string) (models.User, error)
	GetUserByID(ctx context.Context, userID int64) (models.User, error)
}

// PostgresClient is the implementation of user repository for Postgres
type PostgresClient struct {
	dbconn *sql.DB
}

// NewPostgresClient initializes postgres Client Repository
func NewPostgresClient(dbconn *sql.DB) *PostgresClient {
	return &PostgresClient{
		dbconn: dbconn,
	}
}

// GetUser get User by user_name and pass_word
func (c *PostgresClient) GetUser(ctx context.Context, userName, passWord string) (models.User, error) {
	var rs *sql.Rows
	var err error

	rs, err = c.dbconn.QueryContext(ctx, "SELECT * FROM users where user_name = $1 and pass_word = $2 order by id", userName, passWord)

	if err != nil {
		log.Println("[GetUser] query error: ", err.Error())
		return models.User{}, err
	}
	defer rs.Close()

	if rs.Next() {
		user := models.User{}
		err := rs.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Role)
		if err != nil {
			log.Println("[GetUser] scan error: ", err.Error())
			return models.User{}, err
		}

		return user, nil
	}

	return models.User{}, errors.New("user name or password is incorrect")
}

// GetUserByID get User by user_name and pass_word
func (c *PostgresClient) GetUserByID(ctx context.Context, userID int64) (models.User, error) {
	var rs *sql.Rows
	var err error

	rs, err = c.dbconn.QueryContext(ctx, "SELECT * FROM users where id = $1", userID)

	if err != nil {
		log.Println("[GetUserByID] query error: ", err.Error())
		return models.User{}, err
	}
	defer rs.Close()

	if rs.Next() {
		user := models.User{}
		err := rs.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Role)
		if err != nil {
			log.Println("[GetUserByID] scan error: ", err.Error())
			return models.User{}, err
		}

		return user, nil
	}

	return models.User{}, errors.New("could not find user")
}
