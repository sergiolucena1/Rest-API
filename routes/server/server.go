package server

import (
	"github.com/sergiolucena1/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		"5000",
		gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Println("server esta rodando na porta", s.port)
	log.Fatal(router.Run(":" + s.port))
}
