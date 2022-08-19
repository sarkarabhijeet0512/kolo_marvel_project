package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// Module is config module
var Module = fx.Options(
	fx.Provide(
		New,
	),
)

type argvMeta struct {
	desc       string
	defaultVal string
}

// New returns a viper object.
// This object is used to read environment variables or command line arguments.
func New() (config *viper.Viper) {
	config = viper.New()

	confList := map[string]argvMeta{
		"env": {
			defaultVal: "development",
			desc:       "Environment",
		},
		"postgresql_db": {
			defaultVal: "kolo_test_db",
			desc:       "Postgresql db name",
		},
		"postgresql_host": {
			defaultVal: "localhost",
			desc:       "Postgresql host",
		},
		"postgresql_port": {
			defaultVal: "5432",
			desc:       "Postgresql port",
		},
		"postgresql_user": {
			defaultVal: "postgres",
			desc:       "Postgresql username",
		},
		"postgresql_password": {
			defaultVal: "kolotest",
			desc:       "Postgresql password",
		},
		"port": {
			defaultVal: "8765",
			desc:       "Port number of delivery rider API server",
		},
		"mode": {
			defaultVal: "server",
			desc:       "App mode eg. consumer, server, worker",
		},
		"log_level": {
			defaultVal: "debug",
			desc:       "Log level to be printed. List of log level by Priority - debug, info, warn, error, dpanic, panic, fatal",
		},
		"timezone": {
			defaultVal: "Asia/Kolkata",
			desc:       "timezone of the user's country to whom this service is serving. eg. Asia/Kolkata",
		},
		"marvel_public_key": {
			defaultVal: "07216b5045a3252f244a86a0de131be3",
			desc:       "marvel public key",
		},
		"marvel_base_svc": {
			defaultVal: "https://gateway.marvel.com",
			desc:       "marvel main svc",
		},
		"marvel_private_key": {
			defaultVal: "4389006c4089ba422cb299f67e7516b96b87787c",
			desc:       "marvel private key",
		},
		"redis_worker": {
			defaultVal: "localhost:6379",
			desc:       "Postgresql username",
		},
	}

	for key, meta := range confList {
		// automatic conversion of environment var key to `UPPER_CASE` will happen.
		config.BindEnv(key)

		// read command-line arguments
		pflag.String(key, meta.defaultVal, meta.desc)
	}

	pflag.Parse()
	config.BindPFlags(pflag.CommandLine)
	return
}
