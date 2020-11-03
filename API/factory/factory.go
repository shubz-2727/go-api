package factory

import (
	"fmt"

	"../interfaces"
	"../repository"
)

func FactoryDao() interfaces.Repos {

	var i interfaces.Repos
	fmt.Println("in Factory")
	i = repository.CompanyDao{}
	return i
}
