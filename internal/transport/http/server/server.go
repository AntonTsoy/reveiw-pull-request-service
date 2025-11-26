package server

import (
	"context"
	"os"
	"time"

	"github.com/AntonTsoy/review-pull-request-service/internal/transport/http/api"
	"github.com/AntonTsoy/review-pull-request-service/internal/transport/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Server struct {
	app *fiber.App
}

func New(handlers api.ServerInterface) *Server {
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	})

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	app.Use(middleware.Timeout(3 * time.Second))

	api.RegisterHandlers(app, handlers)

	return &Server{
		app: app,
	}
}

func (s *Server) SetSwagger() {
	openAPISpec, err := os.ReadFile("openapi.yml")
	if err != nil {
		return
	}

	s.app.Get("/openapi.yml", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/yaml")
		return c.Send(openAPISpec)
	})

	s.app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Redirect("https://petstore.swagger.io/?url=http://localhost:8080/openapi.yml")
	})
}

func (s *Server) Start(addr string) error {
	return s.app.Listen(addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}
