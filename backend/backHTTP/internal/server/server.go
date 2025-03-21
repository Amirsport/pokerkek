package server

import (
	"fmt"
	"log"
	"pokerok/internal/provider"

	"github.com/gin-gonic/gin"
)

type Server struct {
	r       *gin.Engine
	prov    *provider.Provider
	address string
}

func CreateNewServer(DBprov *provider.Provider, port int, ip string) *Server {
	serv := Server{prov: DBprov, r: gin.Default()}
	serv.r.POST("/registr", serv.Registration)
	serv.address = fmt.Sprintf("%s:%d", ip, port)
	return &serv
}

func (s *Server) Run() {
	if err := s.r.Run(s.address); err != nil {
		log.Fatal(err)
	}
}
