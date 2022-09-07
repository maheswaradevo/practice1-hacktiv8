package repository

import (
	"log"
	"pertemuan6/structs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB gorm.DB
}

type PersonApi interface {
	GetPerson(c *gin.Context, id string) (structs.Person, error)
}

func (idb *InDB) GetPerson(c *gin.Context, id string) (structs.Person, error) {
	var (
		person structs.Person
	)
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		log.Println("ERROR -> Invalid SQL Syntax")
	}
	return person, err
}

func ProvideRepository(DB gorm.DB) *InDB {
	return &InDB{DB: DB}
}
