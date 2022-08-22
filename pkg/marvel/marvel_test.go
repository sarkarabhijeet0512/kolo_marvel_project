package marvel_test

import (
	"context"
	"fmt"
	"kolo_marvel_project/config"
	"kolo_marvel_project/internal/server"
	"kolo_marvel_project/internal/server/handler"
	"kolo_marvel_project/pkg/cache"
	"kolo_marvel_project/pkg/marvel"
	utils "kolo_marvel_project/utils/common"
	"kolo_marvel_project/utils/initialize"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

var router *gin.Engine
var params server.Options

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
		server.Module,
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
func TestService_FetchCharacterDetails(t *testing.T) {
	type fields struct {
		Conf         *viper.Viper
		Log          *logrus.Logger
		CacheService *cache.Service
	}
	type args struct {
		Payload *marvel.Payload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantMcd *marvel.MarvelCharacterDetails
		wantErr bool
	}{
		{
			name: "test_success",
			fields: fields{
				Conf:         params.Config,
				Log:          params.Log,
				CacheService: params.CacheService,
			},
			args: args{
				Payload: &marvel.Payload{
					Apikey: params.Config.GetString("MARVEL_PUBLIC_KEY"),
					Hash:   utils.GetMD5Hash(fmt.Sprint(1) + params.Config.GetString("MARVEL_PRIVATE_KEY") + params.Config.GetString("MARVEL_PUBLIC_KEY")),
					Page:   1,
					Ts:     1,
					Limit:  10,
					Offset: 0,
				},
			},
			wantMcd: &marvel.MarvelCharacterDetails{
				Code: 200,
			},
			wantErr: false,
		},
		{
			name: "test_success_page_2",
			fields: fields{
				Conf:         params.Config,
				Log:          params.Log,
				CacheService: params.CacheService,
			},
			args: args{
				Payload: &marvel.Payload{
					Apikey: params.Config.GetString("MARVEL_PUBLIC_KEY"),
					Hash:   utils.GetMD5Hash(fmt.Sprint(1) + params.Config.GetString("MARVEL_PRIVATE_KEY") + params.Config.GetString("MARVEL_PUBLIC_KEY")),
					Page:   2,
					Ts:     1,
					Limit:  10,
					Offset: 0,
				},
			},
			wantMcd: &marvel.MarvelCharacterDetails{
				Code: 200,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &marvel.Service{
				Conf:         tt.fields.Conf,
				Log:          tt.fields.Log,
				CacheService: tt.fields.CacheService,
			}
			gotMcd, err := s.FetchCharacterDetails(tt.args.Payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.MarvelCharacterList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantMcd.Code, gotMcd.Code)
			// t.Errorf("Service.MarvelCharacterList() = %v, want %v", gotMcd.Code, tt.wantMcd.Code)
		})
	}
}
func TestService_MarvelCharacterList(t *testing.T) {
	type fields struct {
		Conf         *viper.Viper
		Log          *logrus.Logger
		CacheService *cache.Service
	}
	type args struct {
		Payload *marvel.Payload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantMcd *marvel.MarvelCharacterDetails
		wantErr bool
	}{
		{
			name: "test_status_401",
			fields: fields{
				Conf:         params.Config,
				Log:          params.Log,
				CacheService: params.CacheService,
			},
			args: args{
				Payload: &marvel.Payload{},
			},
			wantMcd: &marvel.MarvelCharacterDetails{
				Code: 401,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &marvel.Service{
				Conf:         tt.fields.Conf,
				Log:          tt.fields.Log,
				CacheService: tt.fields.CacheService,
			}
			gotMcd, err := s.MarvelCharacterList(tt.args.Payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.MarvelCharacterList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantMcd.Code, gotMcd.Code)
		})
	}
}
func TestService_FetchCharacterDetailsNameStartsWith(t *testing.T) {
	type fields struct {
		Conf         *viper.Viper
		Log          *logrus.Logger
		CacheService *cache.Service
	}
	type args struct {
		Payload *marvel.Payload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantMcd *marvel.MarvelCharacterDetails
		wantErr bool
	}{
		{
			name: "test_status_401",
			fields: fields{
				Conf:         params.Config,
				Log:          params.Log,
				CacheService: params.CacheService,
			},
			args: args{
				Payload: &marvel.Payload{
					Page:           1,
					NameStartsWith: "thor",
				},
			},
			wantMcd: &marvel.MarvelCharacterDetails{
				Code: 200,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &marvel.Service{
				Conf:         tt.fields.Conf,
				Log:          tt.fields.Log,
				CacheService: tt.fields.CacheService,
			}
			gotMcd, err := s.FetchCharacterDetails(tt.args.Payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.MarvelCharacterList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantMcd.Code, gotMcd.Code)
		})
	}
}
