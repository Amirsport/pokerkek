package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	r    *gin.Engine
	prov Provider
}

func (s *Server) StartServer() {
	s.r = gin.Default()
	s.r.POST("/registr", s.Registration)
	if err := s.r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	return

}
