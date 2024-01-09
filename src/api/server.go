package api

import (
	db "JurusanKu/src/database/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/dummy/create/major", server.createDummyMajor)
	router.GET("/majors", server.getAllMajors)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

