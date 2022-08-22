package initialize

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogrus(conf *viper.Viper) *logrus.Logger {
	return logrus.New()
}
