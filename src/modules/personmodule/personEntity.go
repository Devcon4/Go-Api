package personmodule

import "github.com/jinzhu/gorm"

// Person : Person Entity
type Person struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
