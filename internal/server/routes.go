package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandakn/go-sut-course-api/internal/cache"
	"github.com/pandakn/go-sut-course-api/internal/course"
	"github.com/pandakn/go-sut-course-api/internal/health"
)

func (s *FiberServer) RegisterFiberRoutes() {
	// health check
	s.registerHealthCheckRoutes()

	api := s.App.Group("/api")
	v1 := api.Group("/v1")

	s.registerCourseRoutes(v1)
}

func (s *FiberServer) registerHealthCheckRoutes() {
	healthCheckHandler := health.NewHealthHandler(s.Config)
	s.Get("/", healthCheckHandler.HealthCheck)
}

func (s *FiberServer) registerCourseRoutes(v1 fiber.Router) {
	cache := cache.NewSUTCache()
	courseService := course.NewCourseService()
	courseHandler := course.NewCourseHandler(courseService, cache)

	v1.Post("/courses", courseHandler.GetCourseData)
}
