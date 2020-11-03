package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../factory"
	"../models"
)

var err error

func GetCompanys(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in services")
	var companys []models.Company
	i := factory.FactoryDao()

	companys, err = i.GetAll()
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(companys)
}

func AddCompany(w http.ResponseWriter, r *http.Request) {

	var company models.Company
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println("body= ", body)

	err = json.Unmarshal(body, &company)
	if err != nil {
		panic(err)
	}

	//	fmt.Println("Unmarshal= ", company)

	i := factory.FactoryDao()
	err = i.CreateComp(company)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "ID exist")
		//panic(err)
	}
	w.WriteHeader(http.StatusCreated)

}

func GetByID(w http.ResponseWriter, r *http.Request) {

	var company models.Company

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	i := factory.FactoryDao()
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Insert Valid company ID. You insert \"%s\"", params["id"])

		return

	}

	company, err = i.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, " Company ID \"%s\" not found", params["id"])
		return
	}

	json.NewEncoder(w).Encode(company)

}

func DeleteComp(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	i := factory.FactoryDao()
	err = i.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(w, "ID \"%s\" not exist!!", params["id"])
	}

	w.WriteHeader(http.StatusOK)

}

func Update(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	var company models.Company
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &company)
	if err != nil {
		panic(err)
	}

	if company.Headquarter == "" {
		fmt.Fprintf(w, "Field MISSING")
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	i := factory.FactoryDao()
	err = i.Update(id, company.Headquarter)

	if err != nil {
		fmt.Println("error: ", err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "ID \"%s\" not exist!!", params["id"])
	}

	w.WriteHeader(http.StatusOK)

}
