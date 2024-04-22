package router

import (
	"golang-crud/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagController) *gin.Engine {

	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome home")
	})

	baseRouter := router.Group("/api")
	tagRouter := baseRouter.Group("/tags")
	tagRouter.GET("", tagsController.FindAll)
	tagRouter.GET("/:tagId", tagsController.FindById)
	tagRouter.POST("", tagsController.Create)
	tagRouter.PUT("/:tagId", tagsController.Update)
	tagRouter.DELETE("/:tagId", tagsController.Delete)

	return router
}
