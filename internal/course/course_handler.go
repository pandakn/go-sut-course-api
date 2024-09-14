package course

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pandakn/go-sut-course-api/internal/cache"
	"github.com/pandakn/go-sut-course-api/internal/course/utils"
	cacheUtils "github.com/pandakn/go-sut-course-api/internal/utils"
)

type ICourseHandler interface {
	GetCourseData(c *fiber.Ctx) error
}

type courseHandler struct {
	courseService ICourseService
	cache         *cache.ISUTCache
}

func NewCourseHandler(courseService ICourseService, cache *cache.ISUTCache) ICourseHandler {
	return &courseHandler{
		courseService: courseService,
		cache:         cache,
	}
}

func buildUrl(baseURL string, req *Request) (string, error) {
	maxRow := 50

	if req.MaxRow != 0 {
		maxRow = req.MaxRow
	}

	query := url.Values{
		"coursestatus": []string{"O00"},
		"maxrow":       []string{strconv.Itoa(maxRow)},
		"acadyear":     []string{req.AcadYear},
		"semester":     []string{strconv.Itoa(req.Semester)},
		"coursecode":   []string{req.CourseCode},
		"coursename":   []string{req.CourseName},
	}

	facultyId, err := utils.ConvertFacultyToFacultyId(req.Faculty)
	if err != nil {
		return "", err
	}
	query.Set("facultyid", *facultyId)

	// cmd `2` is no filter
	cmd := "2"

	if req.IsFilter {
		cmd = "1"

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

	cachedData, err := h.cache.Get(req.Faculty, req.CourseCode, req.CourseName, strconv.Itoa(req.Semester),
		req.AcadYear, req.Day, req.TimeFrom, req.TimeTo, req.IsFilter)

	if err == nil && err != cache.ErrCacheMiss {
		var courseResp Response
		if err := cacheUtils.Deserialize([]byte(cachedData), &courseResp); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to deserialize cached data",
			})
		}
		return c.Status(fiber.StatusOK).JSON(courseResp)
	}

	baseURL := "http://reg.sut.ac.th/registrar/class_info_1.asp"

	url, err := buildUrl(baseURL, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	year := fmt.Sprintf("%d/%s", req.Semester, req.AcadYear)
	// shouldInsertData := req.CourseCode == "" && req.CourseName == ""
	shouldInsertData := false
	coursesData, err := h.courseService.GetCourseData(year, url, shouldInsertData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	// Prepare the response tructure
	courseResp := Response{
		Year:    year,
		Faculty: req.Faculty,
		Courses: coursesData,
	}

	serializedResp, err := cacheUtils.Serialize(courseResp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to serialize response data",
		})
	}

	if coursesData != nil {
		err = h.cache.Set(req.Faculty, req.CourseCode, req.CourseName, strconv.Itoa(req.Semester),
			req.AcadYear, req.Day, req.TimeFrom, req.TimeTo, req.IsFilter, serializedResp, time.Hour*3)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to set data in the cache",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(courseResp)
}
