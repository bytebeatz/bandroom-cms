package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bytebeatz/bandroom-cms/api/handler"
	"github.com/bytebeatz/bandroom-cms/api/router"
	"github.com/bytebeatz/bandroom-cms/config"
	_interface "github.com/bytebeatz/bandroom-cms/core/interface"
	"github.com/bytebeatz/bandroom-cms/core/service"
	"github.com/gin-gonic/gin"
)

func Start() error {
	// Load configuration
	config.LoadConfig()

	// Initialize services
	config.InitPostgres()
	ctx := context.Background()
	config.InitGCS(ctx)

	// Init repositories & services
	courseRepo := _interface.NewCoursePG(config.DB)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	unitRepo := _interface.NewUnitPG(config.DB)
	unitService := service.NewUnitService(unitRepo)
	unitHandler := handler.NewUnitHandler(unitService)

	// Setup Gin router with both handlers
	r := router.SetupRouter(courseHandler, unitHandler)

	// Graceful shutdown setup
	srv := &httpServer{
		Engine: r,
		Port:   config.AppConfig.Port,
	}
	return srv.Run()
}

// httpServer encapsulates graceful HTTP server lifecycle
type httpServer struct {
	Engine *gin.Engine
	Port   string
}

func (s *httpServer) Run() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Port),
		Handler: s.Engine,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Bandroom CMS server running at http://localhost:%s", s.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

