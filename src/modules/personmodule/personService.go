package personmodule

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// PersonService : Busines logic to fetch persons
type PersonService struct {
	db     *sqlx.DB
	router *mux.Router
}

// NewPersonService : Create a PersonService
func NewPersonService(db *sqlx.DB, router *mux.Router) *PersonService {
	return &PersonService{db, router}
}

// Get : Returns a Person
func (s PersonService) Get() (*Person, error) {
	return &Person{1, "bob", "ross"}, nil
}

// GetList : Returns a list of Persons
func (s PersonService) GetList() (*[]Person, error) {
	persons := []Person{}
	err := s.db.Select(&persons, "select * from person")

	return &persons, err
}
