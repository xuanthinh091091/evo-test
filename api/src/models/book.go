package models

import "time"

// Book struct type
type Book struct {
	ID            int64
	Name          string
	Author        string
	Publisher     string
	TotalPage     int
	PublishedDate time.Time
	Price         int64
	Stock         int
}
