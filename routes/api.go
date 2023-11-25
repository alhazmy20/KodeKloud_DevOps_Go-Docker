package api

import (
	"log"
	"net/http"

	productsController "github.com/alhazmy20/Go-WebAPI/controllers"
	database "github.com/alhazmy20/Go-WebAPI/database"
	"github.com/alhazmy20/Go-WebAPI/utils"
	"github.com/gorilla/mux"
)

func Initialize() *mux.Router {
	database.EnvFilePath = ".env"
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", defaultPage).Methods("GET")
	muxRouter.HandleFunc("/products", productsController.Index).Methods("GET")
	muxRouter.HandleFunc("/products/{id}", productsController.Show).Methods("GET")
	muxRouter.HandleFunc("/products", productsController.Store).Methods("POST")
	muxRouter.HandleFunc("/products/{id}", productsController.Destroy).Methods("DELETE")
	muxRouter.HandleFunc("/products/{id}", productsController.Update).Methods("PUT")
	return muxRouter
}

func defaultPage(w http.ResponseWriter, r *http.Request) {
	utils.SendJSON(w, http.StatusOK, []byte("Testing Page"))
}

func StartServer(router *mux.Router) error {
	serverAddress := ":5000"
	log.Printf("Server is running on localhost%s", serverAddress)
	return http.ListenAndServe(serverAddress, router)
}
