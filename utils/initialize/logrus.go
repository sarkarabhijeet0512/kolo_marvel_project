package initialize

import (
	"github.com/makasim/sentryhook"
	"github.com/sebest/logrusly"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogrus(conf *viper.Viper) *logrus.Logger {
	logglyToken := conf.GetString("LOGGLY_TOKEN")
	log := logrus.New()
	if !conf.GetBool("sentry_enable") {
		return log
	}
	hook := logrusly.NewLogglyHook(logglyToken, "http://logs-01.loggly.com/inputs", logrus.InfoLevel, "http")
	log.Hooks.Add(hook)
	log.AddHook(sentryhook.New([]logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel}))
	hook.Flush()
	return log
}
