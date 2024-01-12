package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getAllMajors(ctx *gin.Context) {
	majors, err := server.store.GetMajors(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, majors)
}