package handler

import (
	"kolo_marvel_project/er"
	"kolo_marvel_project/pkg/marvel"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MarvelHandler struct {
	log           *logrus.Logger
	marvelservice *marvel.Service
}

func newMarvelHandler(
	log *logrus.Logger,
	marvelservice *marvel.Service,

) *MarvelHandler {
	return &MarvelHandler{
		log:           log,
		marvelservice: marvelservice,
	}
}

type GeneticRes struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

func (h *MarvelHandler) SearchMarvelCharacters(c *gin.Context) {
	var (
		err error
		req = marvel.Payload{}
		res = &GeneticRes{}
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", err).Warn(err.Error())
			return
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		h.log.WithField("span", err).Warn(err.Error())
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		return
	}

	charcaterObj, err := h.marvelservice.FetchCharacterDetails(&req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
		return
	}
	res.Success = true
	res.Message = "charcaterObj"
	res.Data = charcaterObj
	c.JSON(http.StatusOK, res)

}
func (h *MarvelHandler) DeferBlock(c *gin.Context, err error) {

}
