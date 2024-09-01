package main

import (
	"github.com/ghenoo/microservices-go/cmd/api/controllers"
	"github.com/ghenoo/microservices-go/internal/repositories"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

	categoryRoutes.POST("/", func (ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.GET("/", func (ctx *gin.Context) {
		controllers.ListCategories(ctx, inMemoryCategoryRepository)
	})
}