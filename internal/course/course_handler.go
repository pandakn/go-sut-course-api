package course

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pandakn/go-sut-course-api/internal/course/utils"
)

type ICourseHandler interface {
	GetCourseData(c *fiber.Ctx) error
}

type courseHandler struct {
	courseService ICourseService
}

func NewCourseHandler(courseService ICourseService) ICourseHandler {
	return &courseHandler{
		courseService: courseService,
	}
}

func buildUrl(baseURL string, req *Request) (string, error) {
	query := url.Values{
		"coursestatus": []string{"O00"},
		"facultyid":    []string{"all"},
		"maxrow":       []string{req.MaxRow},
		"acadyear":     []string{req.AcadYear},
		"semester":     []string{req.Semester},
		"coursecode":   []string{req.CourseCode},
		"coursename":   []string{req.CourseName},
	}

	day, err := utils.ConvertDayToInteger(req.Day)
	if err != nil {
		return "", err
	}

	timeFrom, err := utils.ConvertTimeToInteger(req.TimeFrom)
	if err != nil {
		return "", err
	}

	timeTo, err := utils.ConvertTimeToInteger(req.TimeTo)
	if err != nil {
		return "", err
	}

	// cmd `2` is no filter
	cmd := "2"

	if req.IsFilter {
		cmd = "1"
		query.Set("cmd", cmd)
		query.Set("weekdays", strconv.Itoa(day))
		query.Set("timefrom", strconv.Itoa(timeFrom))
		query.Set("timeto", strconv.Itoa(timeTo))
	}

	url := baseURL + "?" + query.Encode()
	return url, nil
}

func (h *courseHandler) GetCourseData(c *fiber.Ctx) error {
	req := new(Request)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	baseURL := "http://reg.sut.ac.th/registrar/class_info_1.asp"

	url, err := buildUrl(baseURL, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	coursesData, err := h.courseService.GetCourseData(url)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	year := fmt.Sprintf("%s/%s", req.Semester, req.AcadYear)

	resp := Response{
		Year:    year,
		Courses: coursesData,
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
