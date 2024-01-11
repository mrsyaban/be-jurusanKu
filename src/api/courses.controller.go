package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getCourseByMajorIdRequest struct {
	MajorId int32 `uri:"majorId" binding:"required, min=1"`
}

func (server *Server) getCourseByMajorId(ctx *gin.Context) {
	var req getCourseByMajorIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course, err := server.store.GetCourseByMajorId(ctx, req.MajorId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, course)
}