package server

import (
	"fmt"
	"log"

	dependency "github.com/JoaoGabrielManochio/webapi-go/config"
	"github.com/JoaoGabrielManochio/webapi-go/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {

	if err := dependency.Load(); err != nil {
		fmt.Println(err)
	}

	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() *gin.Engine {
	// gin.SetMode("debug")

	engine := gin.New()

	routes.ConfigRoutes(engine)

	log.Print("server is running at port:", s.port)

	log.Fatal(engine.Run(":" + s.port))

	return engine
}
