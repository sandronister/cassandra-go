package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sandronister/cassandra-go/internal/infra/web/handler"
)

type Server struct {
	engine *gin.Engine
	port   string
}

func NewServer(port string) *Server {
	return &Server{
		engine: gin.Default(),
		port:   port,
	}
}

func (s *Server) Run() error {
	return s.engine.Run(":" + s.port)
}

func (s *Server) AddHDriverTruckHandler(h *handler.DriverTruckHandler) {
	public := s.engine.Group("/api/v1")

	public.GET("/drivers", h.GetDrivers)
	public.GET("/drivers/:id", h.GetDriver)
	public.POST("/drivers", h.CreateDriver)
	public.PUT("/drivers/:id", h.UpdateDriver)
	public.DELETE("/drivers/:id", h.DeleteDriver)
}
