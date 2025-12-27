package routes

import (
	"github.com/Joy-Phathomlezz/Kuchen/internals/handlers"
	"github.com/gorilla/mux"
)

func CakeStoreResult(router *mux.Router, handler *handlers.CakeHandler) {
	router.HandleFunc("/cake/", handler.CreateCake).Methods("POST")
	router.HandleFunc("/cake/", handler.GetCakes).Methods("GET")
	router.HandleFunc("/cake/{CakeId}", handler.GetCakeById).Methods("GET")
	router.HandleFunc("/cake/{CakeId}", handler.UpdateCake).Methods("PUT")
	router.HandleFunc("/cake/{CakeId}", handler.DeleteCake).Methods("DELETE")
}
