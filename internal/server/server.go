package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"kolo_marvel_project/internal/server/handler"
	"kolo_marvel_project/pkg/cache"
	"kolo_marvel_project/pkg/marvel"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	ginlogrus "github.com/toorop/gin-logrus"
	"go.uber.org/fx"
)

// Module invokes mainserver
var Module = fx.Options(
	fx.Invoke(
		Run,
	),
)

const (
	addr = "0.0.0.0"
)

// Options is function arguments struct of `Run` function.
type Options struct {
	fx.In

	Config *viper.Viper
	Log    *logrus.Logger
	// Tracer opentracing.Tracer

	Redis         *redis.Pool `name:"redisWorker"`
	CacheService  *cache.Service
	MarvelService *marvel.Service
	MarvelHandler *handler.MarvelHandler
}

func inLambda(IsinLambda string) bool {
	return IsinLambda == "LAMBDA_TASK_ROOT"
}

// Run starts the mainserver REST API server
func Run(o Options) {
	router := SetupRouter(&o)
	IsinLambda := os.Getenv("LAMBDA_TASK_ROOT")
	if inLambda(IsinLambda) {
		fmt.Println("running aws lambda in aws")
		log.Fatal(gateway.ListenAndServe(addr, SetupRouter(&o)))
	} else {
		fmt.Println("running aws lambda in local")
		router.Run(fmt.Sprintf("%s:%s", addr, o.Config.GetString("port")))
	}
	return
}

// SetupRouter creates gin router and registers all deliveryRider routes to it
func SetupRouter(o *Options) (router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	router.Use(ginlogrus.Logger(o.Log), gin.Recovery())
	// Health routes. DO NOT MOVE IT FROM HERE!
	router.GET("/_healthz", HealthHandler(o))
	router.GET("/_readyz", HealthHandler(o))

	rootRouter := router.Group("/")
	rootRoutes(rootRouter, o)

	v1Routes(rootRouter, o)

	return
}

// HealthHandler
func HealthHandler(o *Options) func(*gin.Context) {
	return func(c *gin.Context) {
		var err error
		if err != nil {
			c.AbortWithError(http.StatusFailedDependency, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"ok": "ok"})
	}
}
