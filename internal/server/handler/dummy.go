package handler

import (
	"kolo_marvel_project/er"
	"kolo_marvel_project/pkg/dummy"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DummyHandler struct {
	log *logrus.Logger

	dummyService *dummy.Service
}

func newDummyHandler(
	log *logrus.Logger,
	dummyService *dummy.Service,
) *DummyHandler {
	return &DummyHandler{
		log:          log,
		dummyService: dummyService,
	}
}

type (
	// DummyRes is the response struct of dummy API
	DummyRes struct {
		// Dummy is the total available dummy.
		Timestamp time.Time `json:"timestamp"`
		Success   bool      `json:"dummy"`
		Error     string    `json:"error"`
	}
)

func (h *DummyHandler) Dummy(c *gin.Context) {
	var (
		err error
		now = time.Now()
		res = &DummyRes{}
	)
	defer func() {
		if err != nil {
			c.Error(err)
			return
		}
	}()

	res.Timestamp = now
	if err != nil {
		res.Error = err.Error()
		if er.From(err).Code != er.UncaughtException {
			return
		}
		err = nil
	}

	c.JSON(http.StatusOK, res)
	return
}
