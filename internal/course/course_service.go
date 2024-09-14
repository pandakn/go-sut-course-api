package course

import (
	"fmt"

	"github.com/pandakn/go-sut-course-api/internal/course/scraper"
)

type ICourseService interface {
	GetCourseData(acadYear, url string, shouldInsert bool) ([]*scraper.IGroupedCourse, error)
}

type courseService struct{}

func NewCourseService() ICourseService {
	return &courseService{}
}

func (s *courseService) GetCourseData(acadYear, url string, shouldInsert bool) ([]*scraper.IGroupedCourse, error) {
	coursesData, err := scraper.ScrapeCourseData(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get course data: %w", err)
	}

	return coursesData, nil
}
