package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/taepunphu/go-sekai-shop-api/config"
	"github.com/taepunphu/go-sekai-shop-api/modules/middleware/middlewareHandler"
	"github.com/taepunphu/go-sekai-shop-api/modules/middleware/middlewareRepository"
	"github.com/taepunphu/go-sekai-shop-api/modules/middleware/middlewareUsecase"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	server struct {
		app        *echo.Echo
		db         *mongo.Client
		cfg        *config.Config
		middleware middlewareHandler.MiddlewareHandlerService
		pctx       context.Context
	}
)

func NewMiddleware(cfg *config.Config) middlewareHandler.MiddlewareHandlerService {
	repo := middlewareRepository.NewMiddlewareRepositoryService()
	usecase := middlewareUsecase.NewMiddlewareUsecaseService(repo)
	return middlewareHandler.NewMiddlewareHandlerService(cfg, usecase)
}

func (s *server) gracefulShutdown(pctx context.Context, quit <-chan os.Signal) {
	log.Printf("Start services: %s", s.cfg.App.Name)
	<-quit
	log.Printf("shutting down services: %s", s.cfg.App.Name)

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
}

func (s *server) httpListening() {
	if err := s.app.Start(s.cfg.App.Url); err != nil && err != http.ErrServerClosed{
		log.Fatal("Error: %v", err)
	}
}

func Start(pctx context.Context, cfg *config.Config, db *mongo.Client) {
	s := &server{
		app:        echo.New(),
		db:         db,
		cfg:        cfg,
		middleware: NewMiddleware(cfg),
		pctx:       pctx,
	}

	// Request timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error Request timeout",
		Timeout:      30 * time.Second,
	}))

	// Cors
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		Skipper:      middleware.DefaultSkipper,
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// body limit
	s.app.Use(middleware.BodyLimit("10M"))

	// custom middleware
	switch s.cfg.App.Name {
	case "auth":
	case "player":
	case "item":
	case "inventory":
	case "payment":
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	
	s.app.Use(middleware.Logger())

	go s.gracefulShutdown(pctx, quit)

	// Listen and serve on
	s.httpListening()

}
