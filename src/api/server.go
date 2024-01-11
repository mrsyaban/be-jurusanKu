package api

import (
	"JurusanKu/src/config"
	db "JurusanKu/src/database/sqlc"
	"github.com/gin-gonic/gin"
	"JurusanKu/src/middleware"
)

type Server struct {
	store      db.Store
	config     config.Config
	router     *gin.Engine
}

func NewServer(store db.Store, runtime_config config.Config) (*Server, error) {
	server := &Server {
		store: store,
		config: runtime_config,
	}
	router := gin.Default()

	router.POST("/register", server.Register)
	router.POST("/login", server.Login)

	authRoutes := router.Group("/").Use(middleware.Auth(runtime_config.JWT_KEY))

	authRoutes.POST("/dummy/create/major", server.createDummyMajor)
	authRoutes.GET("/majors", server.getAllMajors)

	authRoutes.GET("/course/:majorId", server.getCourseByMajorId)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

