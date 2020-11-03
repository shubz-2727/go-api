package interfaces

import (
	"../models"
)

type Repos interface {
	GetAllCompanys() ([]models.Company, error)
	GetById(id int) (models.Company, error)
	CreateComp(comp models.Company) error
	Update(id int, headquater string) error
	Delete(id int) error
}

type ERepos interface {
	GetAllEmployees(id int) ([]models.Employee, error)
	//GetById(id int) (models.Company, error)
	//CreateComp(comp models.Company) error
	//Update(id int, headquater string) error
	//Delete(id int) error
}
