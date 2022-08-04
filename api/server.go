package api

import (
	"github.com/TamerB/products-import-service/api/controller/healthcheck"
	"github.com/TamerB/products-import-service/api/controller/importcsvcontroller"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for ecommerce-products-uploader service
type Server struct {
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer() *Server {
	router := gin.Default()

	router.GET("/healthz", healthcheck.Healthz)
	router.GET("/readyz", healthcheck.Readyz)
	router.POST("/import", importcsvcontroller.ImportCSV)

	return &Server{router: router}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
