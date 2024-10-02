package router

import (
	"postgres_RestApi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/get", controller.GetAllItem).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/get/{id}", controller.GetOneItem).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tambah", controller.AddItem).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/update/{id}", controller.UpdateItem).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete/{id}", controller.DeleteOneItem).Methods("DELETE", "OPTIONS")

	return router
}
