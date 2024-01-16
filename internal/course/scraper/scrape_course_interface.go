package scraper

import "github.com/google/uuid"

type IExamInfo struct {
	Date  string `json:"date"`
	Month string `json:"month"`
	Times string `json:"times"`
	Year  string `json:"year"`
	Room  string `json:"room"`
}

type IExam struct {
	Midterm *IExamInfo `json:"midterm"`
	Final   *IExamInfo `json:"final"`
}

type ICourseName struct {
	En string  `json:"en"`
	Th *string `json:"th"`
}

type IClassSchedule struct {
	Day   string  `json:"day"`
	Times string  `json:"times"`
	Room  *string `json:"room"`
}

type ISeat struct {
	TotalSeat  string `json:"totalSeat"`
	Registered string `json:"registered"`
	Remain     string `json:"remain"`
}

type ICourseDetails struct {
	CourseStatus     string     `json:"courseStatus"`
	CourseCondition  *[]string  `json:"courseCondition"`
	ContinueCourse   *[]string  `json:"continueCourse"`
	EquivalentCourse *[]string  `json:"equivalentCourse"`
	MidExam          *IExamInfo `json:"midExam"`
	FinalExam        *IExamInfo `json:"finalExam"`
}

type ICourse struct {
	ID            uuid.UUID        `json:"id"`
	URL           string           `json:"url"`
	CourseCode    string           `json:"courseCode"`
	Version       string           `json:"version"`
	CourseNameEN  string           `json:"courseNameEN"`
	CourseNameTH  *string          `json:"courseNameTH"`
	Faculty       string           `json:"faculty"`
	Department    string           `json:"department"`
	Note          *string          `json:"note"`
	Professors    []string         `json:"professors"`
	Credit        string           `json:"credit"`
	Section       string           `json:"section"`
	StatusSection string           `json:"statusSection"`
	Language      string           `json:"language"`
	Degree        string           `json:"degree"`
	ClassSchedule []IClassSchedule `json:"classSchedule"`
	Seat          ISeat            `json:"seat"`
	Details       ICourseDetails   `json:"details"`
}

type ISection struct {
	ID            uuid.UUID        `json:"id"`
	URL           string           `json:"url"`
	Section       string           `json:"section"`
	Status        string           `json:"status"`
	Note          *string          `json:"note"`
	Professors    []string         `json:"professors"`
	Language      string           `json:"language"`
	Seat          ISeat            `json:"seat"`
	ClassSchedule []IClassSchedule `json:"classSchedule"`
	Exams         IExam            `json:"exams"`
}

type IGroupedCourse struct {
	CourseCode       string      `json:"courseCode"`
	Version          string      `json:"version"`
	CourseName       ICourseName `json:"courseName"`
	Credit           string      `json:"credit"`
	Degree           string      `json:"degree"`
	Department       string      `json:"department"`
	Faculty          string      `json:"faculty"`
	CourseStatus     string      `json:"courseStatus"`
	CourseCondition  *[]string   `json:"courseCondition"`
	ContinueCourse   *[]string   `json:"continueCourse"`
	EquivalentCourse *[]string   `json:"equivalentCourse"`
	SectionsCount    int         `json:"sectionsCount"`
	Sections         []ISection  `json:"sections"`
}
