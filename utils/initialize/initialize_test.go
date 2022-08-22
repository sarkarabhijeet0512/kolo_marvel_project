package initialize

import (
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func TestNewRedisWorker(t *testing.T) {
	type args struct {
		conf *viper.Viper
		log  *logrus.Logger
	}
	tests := []struct {
		name    string
		args    args
		wantO   RedisWorkerOut
		wantErr bool
	}{
		{
			name: "redis_server_test",
			args: args{
				conf: &viper.Viper{},
				log:  &logrus.Logger{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotO, err := NewRedisWorker(tt.args.conf, tt.args.log)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRedisWorker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(gotO, tt.wantO) {
				t.Errorf("NewRedisWorker() = %v, want %v", gotO, tt.wantO)
			}
		})
	}
}
func TestInitLogrus(t *testing.T) {
	type args struct {
		conf *viper.Viper
	}
	tests := []struct {
		name string
		args args
		want *logrus.Logger
	}{
		{
			name: "logrus_setup_test",
			args: args{
				conf: &viper.Viper{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitLogrus(tt.args.conf); reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitLogrus() = %v, want %v", got, tt.want)
			}
		})
	}
}
