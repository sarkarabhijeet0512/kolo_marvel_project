package main

import (
	"kolo_marvel_project/internal/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

// func SetUpRouter() *gin.Engine {
// 	router := gin.Default()
// 	router.Run()
// 	return router
// }
func TestFetchCharacterDetails(t *testing.T) {
	serverRun()
	o := server.Options{}
	r := server.SetupRouter(&server.Options{})
	// r.Run()
	r.GET("/marvel/character/search", o.MarvelHandler.SearchMarvelCharacters)
	req, _ := http.NewRequest("GET", "/marvel/character/search", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

}
