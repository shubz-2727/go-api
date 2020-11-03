package main

import (
	"net/http"

	serv "./services"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/company", serv.GetCompanys).Methods("GET")
	router.HandleFunc("/company", serv.AddCompany).Methods("POST")
	router.HandleFunc("/company/{id}", serv.GetByID).Methods("GET")
	router.HandleFunc("/company/{id}", serv.Update).Methods("PUT") // Update company's headquater
	router.HandleFunc("/company/{id}", serv.DeleteComp).Methods("DELETE")

	sub := router.PathPrefix("/company").Subrouter()
	sub.HandleFunc("/{id}/employees", GetEmployees).Methods("GET")
	http.ListenAndServe(":8000", router)
}
