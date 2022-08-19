// Package mw is Delivery-rider Middleware package
package mw

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"

	"kolo_marvel_project/er"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := c.Errors.Last()
			if err == nil {
				// no errors, abort with success
				return
			}

			log.Error(err.Err.Error())

			e := er.From(err.Err)

			if !e.NOP {
				sentry.CaptureException(e)
			}
		}()

		c.Next()
	}
}

func ErrorHandlerX(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := c.Errors.Last()
			if err == nil {
				// no errors, abort with success
				return
			}

			log.Error(err.Err.Error())

			e := er.From(err.Err)

			if !e.NOP {
				sentry.CaptureException(e)
			}

			httpStatus := http.StatusInternalServerError
			if e.Status > 0 {
				httpStatus = e.Status
			}

			c.JSON(httpStatus, e)
		}()

		c.Next()
	}
}
