package server

import (
	"kolo_marvel_project/internal/server/mw"

	"github.com/gin-gonic/gin"
)

func rootRoutes(router *gin.RouterGroup, o *Options) {
	r := router.Group("/")

	// middlewares
	r.Use(mw.ErrorHandler(o.Log))
}

func v1Routes(router *gin.RouterGroup, o *Options) {
	r := router.Group("/v1/")

	// middlewares
	r.Use(mw.ErrorHandlerX(o.Log))

	r.GET("/marvel/character/search", o.MarvelHandler.SearchMarvelCharacters)
	//add new routes here

}
