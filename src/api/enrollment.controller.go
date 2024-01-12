package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "JurusanKu/src/database/sqlc"
)

type enrollmentRequest struct {
	Email string `form:"email" binding:"required"`
	CourseID int64`form:"course_id" binding:"required"`
}

func (server *Server) enroll(ctx *gin.Context) {
	var req enrollmentRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// register user
	args := db.AddEnrollmentParams{
		UserEmail: req.Email,
		CourseID: req.CourseID,
	}

	enrollment, err := server.store.AddEnrollment(ctx, args)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, enrollment)
}

type getActiveCourse struct {
	Email string `form:"email" binding:"required"`
}

func (server *Server) GetActiveCourse(ctx *gin.Context) {
	var req getActiveCourse
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	courses, err := server.store.GetActiveCourse(ctx, req.Email)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}
