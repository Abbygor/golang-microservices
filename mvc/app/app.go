package app

import (
	"net/http"

	"github.com/Abbygor/golang-microservices/mvc/controllers"
)

func StartArpp() {
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
