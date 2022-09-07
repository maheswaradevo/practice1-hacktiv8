package service

import (
	"pertemuan6/repository"

	"github.com/gin-gonic/gin"
)

type ServiceImpl struct {
	rr repository.PersonApi
}

type PersonServiceApi interface {
	GetPersonService(c *gin.Context) gin.H
}

func (s ServiceImpl) GetPersonService(c *gin.Context) gin.H {
	id := c.Param("id")
	res, err := s.rr.GetPerson(c, id)
	if err != nil {
		c.JSON(500, "internal server error")
	}
	result := gin.H{
		"result": res,
	}
	return result
}

func ProvideService(rr repository.PersonApi) *ServiceImpl {
	return &ServiceImpl{rr: rr}
}
