package handlers

import "github.com/gin-gonic/gin"

type SongsHandlersInterface interface {
	GetAllSongs(*gin.Context)
	DeleteSongs(*gin.Context)
}
