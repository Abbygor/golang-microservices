package services

import (
	"net/http"

	"github.com/Abbygor/golang-microservices/mvc/domain"
	"github.com/Abbygor/golang-microservices/mvc/utils"
)

type itemsService struct {
}

var (
	IntemsService itemsService
)

func (i *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
