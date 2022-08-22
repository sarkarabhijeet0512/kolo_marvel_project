package handler_test

import (
	"context"
	"kolo_marvel_project/config"
	"kolo_marvel_project/internal/server"
	"kolo_marvel_project/pkg/cache"
	"kolo_marvel_project/pkg/dummy"
	"kolo_marvel_project/pkg/marvel"
	"kolo_marvel_project/utils/initialize"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var params server.Options

func setupMockServer() (router *gin.Engine) {
	// params = server.Options{}
	gin.SetMode(gin.TestMode)
	app := fx.New(
		fx.Provide(
			// postgresql
			initialize.NewRedisWorker,
		),
		config.Module,
		initialize.Module,
		// Module,
		server.Module,
		dummy.Module,
		marvel.Module,
		cache.Module,
		// Run app forever

		fx.Populate(&params),
	)
	app.Start(context.TODO())
	defer app.Stop(context.TODO())
	router = server.SetupRouter(&params)
	return
}
func init() {
	router = setupMockServer()
}
func TestHealthz(t *testing.T) {
	// router = setupMockServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_healthz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"ok":"ok"}`, w.Body.String())
}

func TestFetchCharacterDetails(t *testing.T) {
	// router = setupMockServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/marvel/character/search", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"success":true,"message":"Marvel Character Details Successful","data":{"offset":0,"limit":0,"total":0,"count":0,"results":null},"meta":null}`, w.Body.String())
}
