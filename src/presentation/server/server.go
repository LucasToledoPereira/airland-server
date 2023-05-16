package server

import (
	"fmt"

	"airland-server/src/cross_cutting/config"

	"airland-server/src/presentation/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(r *router.Routes) (server *Server) {
	return &Server{
		router: r.Router,
	}
}

func (s *Server) Run() {
	fmt.Printf("collection manager listening on port '%s'", config.C.Server.Port)
	fmt.Println()

	panic(s.router.Run(":" + config.C.Server.Port))
}
