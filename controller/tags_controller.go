package controller

import (
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagController {
	return &TagController{
		tagsService: service,
	}
}

// Create controller
func (controller *TagController) Create(ctx *gin.Context) {
	createTagRequest := request.CreateTagRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)
	controller.tagsService.Create(createTagRequest)
	webResponse := response.Respose{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   nil,
	}

	ctx.Header("Conten-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)
}

// Update controller
func (controller *TagController) Update(ctx *gin.Context) {
	updateTagRequest := request.UpdateTagRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagRequest.Id = id
	controller.tagsService.Update(updateTagRequest)
	webResponse := response.Respose{
		Code:   http.StatusOK,
		Status: "Updated",
		Data:   nil,
	}

	ctx.Header("Conten-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete controller
func (controller *TagController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagsService.Delete(id)
	webResponse := response.Respose{
		Code:   http.StatusNoContent,
		Status: "Deleted",
		Data:   nil,
	}

	ctx.Header("Conten-Type", "application/json")
	ctx.JSON(http.StatusNoContent, webResponse)
}

// Findby controller
func (controller *TagController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagsService.FindById(id)

	webResponse := response.Respose{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}

	ctx.Header("Conten-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindAll controller
func (controller *TagController) FindAll(ctx *gin.Context) {

	tagResponse := controller.tagsService.FindAll()

	webResponse := response.Respose{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}
	ctx.Header("Conten-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
