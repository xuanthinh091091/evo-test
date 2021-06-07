package controllers

import (
	"candidate/api/src/dto/request"
	"candidate/api/src/dto/response"
	"candidate/api/src/models"
	"candidate/api/src/service/order"
	"candidate/api/src/utils"
	"encoding/json"
	"net/http"
)

// OrderBook handers OrderBook  http request and return reponse
func OrderBook(service order.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestModel request.BookOrder
		err := json.NewDecoder(r.Body).Decode(&requestModel)
		if err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		errRes := service.OrderBook(r.Context(), requestModel.UserID, requestModel.BookID, requestModel.Quantity, requestModel.ShippingDetails)
		if errRes != nil {
			utils.RespondJSON(w, http.StatusInternalServerError, errRes.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

// GetAllOrders handers GetAllOrders http request and return reponse
func GetAllOrders(service order.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orders, err := service.GetAllOrders(r.Context())
		if err != nil {
			utils.RespondJSON(w, err.Code, err.Error())
			return
		}
		response := make([]response.Order, 0)
		for _, order := range orders {
			response = append(response, toOrderResponse(order))
		}
		utils.RespondJSON(w, http.StatusOK, response)
	})
}

func toOrderResponse(order models.Order) response.Order {
	return response.Order{
		ID:              order.ID,
		UserName:        order.UserName,
		BookName:        order.BookName,
		Quantity:        order.Quantity,
		ShippingDetails: order.ShippingDetails,
		CreateAt:        order.CreateAt.Format("2006-01-02 15:04:05"),
	}
}
