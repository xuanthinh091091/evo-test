package controllers

import (
	"candidate/api/src/service/book"
	"candidate/api/src/service/order"
	"candidate/api/src/service/user"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

// PublicRouter registers handlers to the router provided in the argument
func PublicRouter(userService user.Service) func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/login", Login(userService))
	}
}

// UserRouter registers handlers to the router provided in the argument
func UserRouter(userService user.Service, bookService book.Service, orderService order.Service) func(chi.Router) {
	return func(r chi.Router) {
		r.Use(authenUser(userService, "USER"))
		r.Route("/user", func(r chi.Router) {
			r.Get("/booksavailableforsale", GetBooksAvailableForSale(bookService))
			r.Post("/book/order", OrderBook(orderService))
		})
	}
}

// AdminRouter registers handlers to the router provided in the argument
func AdminRouter(userService user.Service, bookService book.Service, orderService order.Service) func(chi.Router) {
	return func(r chi.Router) {
		r.Use(authenUser(userService, "ADMIN"))
		r.Route("/admin", func(r chi.Router) {
			r.Get("/books", GetAllBooks(bookService))
			r.Get("/orders", GetAllOrders(orderService))
		})
	}
}

func authenUser(userService user.Service, role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			userID := r.Header.Get("userID")
			if strings.TrimSpace(userID) == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			id, err := strconv.ParseInt(userID, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			user, err := userService.GetUserByID(ctx, id)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if user.Role != role {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
