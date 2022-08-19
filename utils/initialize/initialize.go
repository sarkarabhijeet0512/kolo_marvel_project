package initialize

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module provides common application dependencies
var Module = fx.Options(
	fx.Provide(
		newLogger,
		InitLogrus,
	),
	fx.Invoke(
		LivenessProbe,
	),
)

// LivenessProbe writes an empty at /tmp/_healthz at every 30 seconds
// K8s checks if the last modified of the file is <= 30 seconds, if not, pod is restarted.
func LivenessProbe(logger *logrus.Logger) {
	go func() {
		for {
			time.Sleep(30 * time.Second)
			err := ioutil.WriteFile("/tmp/_healthz", []byte{}, os.ModePerm)
			if err != nil {
				sentry.CaptureException(err)
				logger.Info("ERROR: livenessprobe /tmp/_healthz file creation failed",
					zap.Error(err),
				)
			}
		}
	}()
	return
}
