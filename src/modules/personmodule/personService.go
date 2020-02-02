package personmodule

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// PersonService : Busines logic to fetch persons
type PersonService struct {
	db     *gorm.DB
	router *mux.Router
}

// NewPersonService : Create a PersonService
func NewPersonService(db *gorm.DB, router *mux.Router) *PersonService {
	return &PersonService{db, router}
}

// Get : Returns a Person
func (s PersonService) Get() (*Person, error) {
	return nil, nil
}

// GetList : Returns a list of Persons
func (s PersonService) GetList() (*[]Person, error) {
	persons := []Person{}
	err := s.db.Find(&persons).Error

	return &persons, err
}
