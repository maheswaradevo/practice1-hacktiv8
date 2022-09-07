package structs

import "gorm.io/gorm"

type PersonGorm struct {
	gorm.Model
	FirstName string
	LastName  string
}
