package controllers

import (
	"candidate/api/src/dto/response"
	"candidate/api/src/models"
	"candidate/api/src/service/book"
	"candidate/api/src/utils"
	"fmt"
	"net/http"

	"github.com/dustin/go-humanize"
)

// GetBooksAvailableForSale handers GetBookAvailableForSale http request and return reponse
func GetBooksAvailableForSale(service book.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		books, err := service.GetBooksAvailableForSale(r.Context())
		if err != nil {
			utils.RespondJSON(w, err.Code, err.Error())
			return
		}
		response := make([]response.Book, 0)
		for _, book := range books {
			response = append(response, toBookResponse(book))
		}
		utils.RespondJSON(w, http.StatusOK, response)
	})
}

// GetAllBooks handers GetAllBook http request and return reponse
func GetAllBooks(service book.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		books, err := service.GetAllBooks(r.Context())
		if err != nil {
			utils.RespondJSON(w, err.Code, err.Error())
			return
		}
		response := make([]response.Book, 0)
		for _, book := range books {
			response = append(response, toBookResponse(book))
		}
		utils.RespondJSON(w, http.StatusOK, response)
	})
}

func toBookResponse(book models.Book) response.Book {
	return response.Book{
		ID:            book.ID,
		Name:          book.Name,
		Author:        book.Author,
		Publisher:     book.Publisher,
		TotalPage:     book.TotalPage,
		PublishedDate: book.PublishedDate.Format("2006-01-02 15:04:05"),
		Price:         fmt.Sprintf("%s VND", humanize.Comma(book.Price)),
		Stock:         book.Stock,
	}
}
