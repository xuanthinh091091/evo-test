package controllers

import (
	"candidate/api/src/dto/request"
	"candidate/api/src/dto/response"
	"candidate/api/src/service/user"
	"candidate/api/src/utils"
	"encoding/json"
	"net/http"
)

// Login handers Login http request and return reponse
func Login(service user.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestModel request.User
		err := json.NewDecoder(r.Body).Decode(&requestModel)
		if err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		user, errRes := service.GetUser(r.Context(), requestModel.UserName, requestModel.PassWord)
		if errRes != nil {
			utils.RespondJSON(w, http.StatusInternalServerError, errRes.Error())
			return
		}
		utils.RespondJSON(w, http.StatusOK, response.User{
			UserID:   user.ID,
			UserName: user.UserName,
			Role:     user.Role,
		})
	})
}
