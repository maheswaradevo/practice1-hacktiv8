package main

import (
	"pertemuan6/config"
	controller "pertemuan6/controller"
	db "pertemuan6/database"
	"pertemuan6/repository"
	"pertemuan6/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := db.InitializeDatabase(cfg.Database.Username, cfg.Database.Password, cfg.Database.Port, cfg.Database.Name, cfg.Database.Host)
	personRepository := repository.ProvideRepository(*db)
	personService := service.ProvideService(personRepository)
	personController := controller.ProvidePersonController(personService)
	router := gin.Default()
	router.GET("/person/:id", personController.GetPersonController)
	// router.GET("/person/:id", InDB.GetPerson)
	// router.GET("/persons", InDB.GetPersons)
	// router.POST("/person", InDB.CreatePerson)
	// router.PUT("/person", InDB.UpdatePerson)
	// router.DELETE("/person/:id", InDB.DeletePerson)

	router.Run(cfg.ServerPort)

}
