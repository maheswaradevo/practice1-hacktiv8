package contrroller

import (
	"net/http"
	"pertemuan6/service"
	"pertemuan6/structs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB gorm.DB
}

type personController struct {
	ps service.PersonServiceApi
}

func ProvidePersonController(ps service.PersonServiceApi) *personController {
	return &personController{ps: ps}
}

func (pc *personController) GetPersonController(c *gin.Context) {
	res := pc.ps.GetPersonService(c)
	c.JSON(200, res)
}

// func (pc *personController) InitializeController() {
// 	router.GET("/person/:id", pc.GetPersonController)
// }

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	requestBody := structs.PersonRequestBody{}

	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	person.FirstName = requestBody.FirstName
	person.LastName = requestBody.LastName

	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")

	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newPerson.FirstName = firstName
	newPerson.LastName = lastName

	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "Delete failed",
		}
	} else {
		result = gin.H{
			"result": "data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
