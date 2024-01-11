package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createDummyMajor(ctx *gin.Context) {
	major, err := server.store.CreateMajor(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, major)
}


func (server *Server) getAllMajors(ctx *gin.Context) {
	majors, err := server.store.GetMajors(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, majors)
}