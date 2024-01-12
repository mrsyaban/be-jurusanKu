package api

import (
	"database/sql"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"JurusanKu/src/config"
)

func (server *Server) VerifyEnrollment(key string, ctx *gin.Context, courseId int64) bool {

	token, err := ctx.Cookie("token")
	if err != nil {
		return false
	}

	email, err := config.GetUserEmailByToken(token, key)
	if err != nil {
		return false
	}

	course, err := server.store.GetActiveCourse(ctx, email)
	if err != nil {
		return false
	}
	
	if !slices.Contains(course, courseId) {
		return false
	}

	return true
}	

type getCourseByMajorIdRequest struct {
	MajorId int64 `uri:"majorId" binding:"required"`
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

type getSyllabusRequest struct {
	CourseID int64 `uri:"courseId" binding:"required"`
}

func (server *Server) getSyllabusByCourseId(ctx *gin.Context) {
	var req getSyllabusRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !server.VerifyEnrollment(server.config.JWT_KEY, ctx, req.CourseID) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized to access this course"})
		return
	}

	course, err := server.store.GetSyllabusByCourseId(ctx, req.CourseID)
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

type getMaterialByCourseIdRequest  struct {
	CourseID int64 `uri:"courseId" binding:"required"`
}

func (server *Server) getMaterialByCourseId(ctx *gin.Context) {
	var req getMaterialByCourseIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !server.VerifyEnrollment(server.config.JWT_KEY, ctx, req.CourseID) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized to access this course"})
		return
	}

	materials, err := server.store.GetMaterialByCourseId(ctx, req.CourseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, materials)
}

type getProjectByCourseIdRequest struct {
	CourseID int64 `uri:"courseId" binding:"required"`
}

func (server *Server) getProjectByCourseId(ctx *gin.Context) {
	var req getProjectByCourseIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !server.VerifyEnrollment(server.config.JWT_KEY, ctx, req.CourseID) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized to access this course"})
		return
	}

	projects, err := server.store.GetProjectByCourseId(ctx, req.CourseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, projects)
}
