package server_test

import (
	"context"
	"kolo_marvel_project/config"
	"kolo_marvel_project/internal/server"
	"kolo_marvel_project/internal/server/handler"
	"kolo_marvel_project/pkg/cache"
	"kolo_marvel_project/pkg/dummy"
	"kolo_marvel_project/pkg/marvel"
	"kolo_marvel_project/utils/initialize"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func setupMockServer() (router *gin.Engine) {
	params := server.Options{}

	app := fx.New(
		fx.Provide(
			// postgresql
			initialize.NewRedisWorker,
		),
		config.Module,
		initialize.Module,
		server.Module,
		handler.Module,
		dummy.Module,
		marvel.Module,
		cache.Module,
		// Run app forever

		// fx.Populate(&params),
	)
	go app.Start(context.TODO())
	time.Sleep(time.Second * 5)
	// ;
	//  err != nil {
	// 	panic(err)
	// }
	// defer app.Stop(context.TODO())

	Router = server.SetupRouter(&params)
	Router.Run()
	return
}

var Router *gin.Engine

func intz(t *testing.T) {
	Router = setupMockServer()
}
func init() {
	go intz(&testing.T{})
	time.Sleep(time.Second * 5)
}
func TestHealthz(t *testing.T) {
	// router := setupMockServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_healthz", nil)
	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"ok":"ok"}`, w.Body.String())
}

// func TestFetchCharacterDetails(t *testing.T) {
// 	serverRun()
// 	o := server.Options{}
// 	r := server.SetupRouter(&server.Options{})
// 	// r.Run()
// 	r.GET("/marvel/character/search", o.MarvelHandler.SearchMarvelCharacters)
// 	req, _ := http.NewRequest("GET", "/marvel/character/search", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// }
