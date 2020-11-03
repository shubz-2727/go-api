package repository

import (
	"database/sql"

	"../models"
	_ "github.com/go-sql-driver/mysql"
)

type EmployeeDao struct {
}

func dbConnection() *sql.DB {

	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/organization")
	if err != nil {
		panic(err.Error())
	}
	//	defer db.Close()

	return db
}

func (cDao CompanyDao) GetAllEmployee(id int) ([]models.Employee, error) {

	db := dbConnection()
	defer db.Close()

	var companys []models.Company

	result, err := db.Query("SELECT * from companys")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var company models.Company
		err := result.Scan(&company.ID, &company.Name, &company.Headquarter)
		if err != nil {
			panic(err.Error())
		}
		companys = append(companys, company)
	}
	return companys, nil
}
