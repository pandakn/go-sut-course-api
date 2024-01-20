package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandakn/go-sut-course-api/internal/cache"
	"github.com/pandakn/go-sut-course-api/internal/course"
)

func (s *FiberServer) RegisterFiberRoutes() {
	api := s.App.Group("/api")
	v1 := api.Group("/v1")

	s.registerCourseRoutes(v1)
}

func (s *FiberServer) registerCourseRoutes(v1 fiber.Router) {
	cache := cache.NewSUTCache()
	courseService := course.NewCourseService()
	courseHandler := course.NewCourseHandler(courseService, cache)

	v1.Post("/courses", courseHandler.GetCourseData)
}
