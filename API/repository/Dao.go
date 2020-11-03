package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"../models"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var row *sql.Row
var err error

type CompanyDao struct {
}

func dbConnection() *sql.DB {

	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/organization")
	if err != nil {
		panic(err.Error())
	}
	//	defer db.Close()

	return db
}

func (cDao CompanyDao) GetAllCompanys() ([]models.Company, error) {

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

func (cDao CompanyDao) CreateComp(comp models.Company) error {

	query := "insert into companys (ID,name,headquater) values(?,?,?)"

	db := dbConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(comp.ID, comp.Name, comp.Headquarter)
	if err != nil {
		return err
	}

	return nil

}

func (cDao CompanyDao) GetById(id int) (models.Company, error) {

	//query := "select ID,name,headquater from comapnys where id = ?"
	var company models.Company
	db := dbConnection()
	defer db.Close()

	result := db.QueryRow("select ID,name,headquater from companys where id = ?", id).Scan(&company.ID, &company.Name, &company.Headquarter)

	if result == sql.ErrNoRows {

		return company, errors.New("no record found")
	}

	//for result.Next() {
	//_ = result.Scan(&company.ID, &company.Name, &company.Headquarter)
	//	}
	return company, nil
}

func (cDao CompanyDao) Delete(id int) error {

	db := dbConnection()
	defer db.Close()

	stmt, err := db.Prepare("delete from companys where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, _ := stmt.Exec(id)
	if count, _ := result.RowsAffected(); count == 0 {
		return errors.New("ID Not Found")
	}

	return nil
}

func (cDao CompanyDao) Update(id int, headquater string) error {

	db := dbConnection()
	defer db.Close()

	stmt, err := db.Prepare("update companys set headquater= ? where ID = ?")
	fmt.Println("id", id)
	if err != nil {

		return err
	}
	defer stmt.Close()

	result, _ := stmt.Exec(headquater, id)
	if count, _ := result.RowsAffected(); count == 0 {
		return errors.New("ID Not Found")
	}

	return nil
}
