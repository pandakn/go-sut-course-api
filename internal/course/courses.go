package course

import "github.com/pandakn/go-sut-course-api/internal/course/scraper"

type Request struct {
	AcadYear   string `json:"acadYear" form:"acadYear"`
	Semester   string `json:"semester" form:"semester"`
	CourseCode string `json:"courseCode" form:"courseCode"`
	CourseName string `json:"courseName" form:"courseName"`
	MaxRow     string `json:"maxRow" form:"maxRow"`
	IsFilter   bool   `json:"isFilter" form:"isFilter"`
	Day        string `json:"day,omitempty" form:"day,omitempty"`
	TimeFrom   string `json:"timeFrom,omitempty" form:"timeFrom,omitempty"`
	TimeTo     string `json:"timeTo,omitempty" form:"timeTo,omitempty"`
}

type Response struct {
	Year    string                    `json:"year"`
	Courses []*scraper.IGroupedCourse `json:"courses"`
}
