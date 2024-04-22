package main

import (
	"golang-crud/config"
	"golang-crud/controller"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
	"golang-crud/router"
	"golang-crud/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Logger.Info().Msg("Started server")

	//Database setup
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	//Repository
	tagsRepository := repository.NewTagsRepositoryImpl(*db)

	//Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	//Controller
	tagsController := controller.NewTagsController(tagsService)

	//Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
