package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MarvelHandler struct {
	log *logrus.Logger
}

func newMarvelHandler(
	log *logrus.Logger,
) *MarvelHandler {
	return &MarvelHandler{
		log: log,
	}
}
func (h *MarvelHandler) SearchMarvelCharacters(c *gin.Context) {

}
