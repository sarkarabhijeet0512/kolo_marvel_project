package server

import (
	"context"
	"kolo_marvel_project/config"
	"kolo_marvel_project/internal/server/handler"
	"kolo_marvel_project/pkg/cache"
	"kolo_marvel_project/pkg/marvel"
	"kolo_marvel_project/utils/initialize"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

var router *gin.Engine
var params Options

func setupMockServer() (router *gin.Engine) {
	// params = server.Options{}
	gin.SetMode(gin.TestMode)
	app := fx.New(
		fx.Provide(
			//redis server
			initialize.NewRedisWorker,
		),
		config.Module,
		initialize.Module,
		handler.Module,
		Module,
		marvel.Module,
		cache.Module,
		// Run app forever

		fx.Populate(&params),
	)
	app.Start(context.TODO())
	defer app.Stop(context.TODO())
	router = SetupRouter(&params)
	return
}
func init() {
	router = setupMockServer()
}
func TestHealthz(t *testing.T) {
	// router = setupMockServer()
	assert.NotNil(t, router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_healthz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"ok":"ok"}`, w.Body.String())
}

func TestRun(t *testing.T) {
	type args struct {
		o Options
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run(tt.args.o)
		})
	}
}

func Test_inLambda(t *testing.T) {
	type args struct {
		IsinLambda string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "this_should_be_true",
			args: args{
				IsinLambda: "LAMBDA_TASK_ROOT",
			},
			want: true,
		},
		{
			name: "this_should_be_true",
			args: args{
				IsinLambda: "IS_NOT_LAMBDA_TASK_ROOT",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inLambda(tt.args.IsinLambda); got != tt.want {
				t.Errorf("inLambda() = %v, want %v", got, tt.want)
			}
		})
	}
}
