package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Abbygor/golang-microservices/mvc/services"
	"github.com/Abbygor/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	log.Print("About to process user_id: ", userIdParam)
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		// Just return the Bad Request to the client
		return
	}

	user, apiErr := services.GetUser(userId)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		// Handle the error and return to the client
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
