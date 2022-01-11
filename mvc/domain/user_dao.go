package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abbygor/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Luis", LastName: "Feregrino", Email: "djlaf2008@hotmail.com"},
		456: {Id: 1, FirstName: "Nanci", LastName: "Olvera", Email: "nanchy92@hotmail.com"},
	}
	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {
}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("Accesing the database...")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
