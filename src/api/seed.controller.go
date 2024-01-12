package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (server *Server) seed(ctx *gin.Context) {
	major, err := server.store.CreateDummyMajor(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	course, err := server.store.CreateCourse(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	projects, err := server.store.CreateDummyProject(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	materials, err := server.store.CreateDummyMaterial(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": map[string]interface{}{
		"courses": course,
		"majors":  major,
		"projects": projects,
		"materials": materials,
	}})
}
