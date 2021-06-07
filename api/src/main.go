package main

import (
	"candidate/api/src/controllers"
	bookRepository "candidate/api/src/repository/book"
	orderRepository "candidate/api/src/repository/order"
	userRepository "candidate/api/src/repository/user"
	"candidate/api/src/service/book"
	"candidate/api/src/service/order"
	"candidate/api/src/service/user"
	"candidate/api/src/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	dbconn := utils.CreateConnection()
	defer dbconn.Close()
	log.Println("Successfully connected!")
	// Init user services
	userRepository := userRepository.NewPostgresClient(dbconn)
	userService := user.NewManager(userRepository)
	// Init book services
	bookRepository := bookRepository.NewPostgresClient(dbconn)
	bookService := book.NewManager(bookRepository)
	// Init order services
	orderRepository := orderRepository.NewPostgresClient(dbconn)
	orderService := order.NewManager(orderRepository)

	r := chi.NewRouter()
	r.Group(controllers.PublicRouter(userService))
	r.Group(controllers.AdminRouter(userService, bookService, orderService))
	r.Group(controllers.UserRouter(userService, bookService, orderService))
	log.Println("Server started on: http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
