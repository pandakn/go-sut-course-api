package scraper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/google/uuid"
)

var statusObj = map[string]string{
	"A": "เพิ่มผ่าน WEB ได้เท่านั้น",
	"C": "ปิดไม่รับลง",
	"D": "ถอนผ่าน WEB ได้เท่านั้น",
	"N": "เปิดลงปกติ ทำการโดยเจ้าหน้าที่เท่านั้น",
	"W": "เปิดลงปกติ สามารถลงทะเบียนผ่าน WEB ได้",
	"X": "เปลี่ยนกลุ่มผ่าน WEB ได้เท่านั้น",
}

func groupedCourses(courses []ICourse) []*IGroupedCourse {
	groupedCoursesMap := make(map[string]*IGroupedCourse)

	for _, course := range courses {
		key := course.CourseCode + "-" + course.Version

		if _, ok := groupedCoursesMap[key]; !ok {
			groupedCoursesMap[key] = &IGroupedCourse{
				CourseCode: course.CourseCode,
				Version:    course.Version,
				CourseName: ICourseName{
					En: course.CourseNameEN,
					Th: course.CourseNameTH,
				},
				Credit:           course.Credit,
				Degree:           course.Degree,
				Department:       course.Department,
				Faculty:          course.Faculty,
				CourseStatus:     course.Details.CourseStatus,
				CourseCondition:  course.Details.CourseCondition,
				ContinueCourse:   course.Details.ContinueCourse,
				EquivalentCourse: course.Details.EquivalentCourse,
			}
		}

		sectionExists := false
		for _, existingSection := range groupedCoursesMap[key].Sections {
			if existingSection.Section == course.Section {
				sectionExists = true
				break
			}
		}

		if !sectionExists {
			section := ISection{
				ID:            course.ID,
				URL:           course.URL,
				Section:       course.Section,
				Status:        course.StatusSection,
				Note:          course.Note,
				Professors:    course.Professors,
				Language:      course.Language,
				Seat:          course.Seat,
				ClassSchedule: course.ClassSchedule,
				Exams: IExam{
					Midterm: course.Details.MidExam,
					Final:   course.Details.FinalExam,
				},
			}
			groupedCoursesMap[key].Sections = append(groupedCoursesMap[key].Sections, section)
		}

		groupedCoursesMap[key].SectionsCount = len(groupedCoursesMap[key].Sections)
	}

	var groupedCourses []*IGroupedCourse
	for _, value := range groupedCoursesMap {
		groupedCourses = append(groupedCourses, value)
	}

	return groupedCourses
}

func parseCourseCode(courseCodeStr string) (string, string) {
	parts := strings.Split(courseCodeStr, "-")
	courseCode := strings.TrimSpace(parts[0])
	version := ""
	if len(parts) > 1 {
		version = strings.TrimSpace(parts[1])
	}
	return courseCode, version
}

func parseSchedule(schedule string) []IClassSchedule {
	timeRegex := regexp.MustCompile(`\d{2}:\d{2}-\d{2}:\d{2}`)
	dayRegex := regexp.MustCompile(`Mo|Tu|We|Th|Fr|Sa|Su`)

	times := timeRegex.FindAllString(schedule, -1)
	days := dayRegex.FindAllString(schedule, -1)

	var classSchedules []IClassSchedule

	if times == nil && days == nil {
		return classSchedules
	}

	for i, time := range times {
		classSchedules = append(classSchedules, IClassSchedule{
			Day:   days[i],
			Times: time,
		})
	}

	return classSchedules
}

func mergeSchedulesWithRooms(schedules []IClassSchedule, rooms []string) []IClassSchedule {
	var classSchedules []IClassSchedule

	for idx, schedule := range schedules {
		classSchedule := IClassSchedule{
			Day:   schedule.Day,
			Times: schedule.Times,
			Room:  &rooms[idx],
		}

		classSchedules = append(classSchedules, classSchedule)
	}

	return classSchedules
}

func ScrapeCourseData(url string) ([]*IGroupedCourse, error) {
	var courseData []ICourse

	c := colly.NewCollector(
		colly.DetectCharset(),
	)

	c.OnRequest(func(r *colly.Request) {
		r.ResponseCharacterEncoding = "windows-874"
	})

	c.OnHTML("table:nth-child(2) tr[valign]", func(e *colly.HTMLElement) {
		courseCodeStr := e.ChildText("td:nth-child(2)")
		courseCode, version := parseCourseCode(courseCodeStr)

		corseCodeUrl := e.ChildAttr("td:nth-child(2) a", "href")
		urlCourseDetails := fmt.Sprintf("http://reg.sut.ac.th/registrar/%s", corseCodeUrl)

		courseDetails, _ := scrapeCourseDetails(urlCourseDetails)

		professors := e.ChildTexts("td:nth-child(3) font[color='#407060'] li")

		noteEl := e.DOM.Find("td:nth-child(3) font[color='#660000']").Contents().First()
		noteStr := noteEl.Text()
		note := strings.TrimSpace(strings.Trim(noteStr, "( )"))

		credit := strings.TrimSpace(e.ChildText("td:nth-child(4)"))
		language := strings.TrimSpace(strings.Split(e.ChildText("td:nth-child(5)"), ":")[0])
		degree := strings.TrimSpace(e.ChildText("td:nth-child(6)"))

		section := strings.TrimSpace(e.ChildText("td:nth-child(8)"))

		statusSectionCode := e.ChildText("td:nth-child(12)")
		statusSection := statusObj[statusSectionCode]

		totalSeat := strings.TrimSpace(e.ChildText("td:nth-child(9)"))
		registered := strings.TrimSpace(e.ChildText("td:nth-child(10)"))
		remain := strings.TrimSpace(e.ChildText("td:nth-child(11)"))
		seat := ISeat{
			TotalSeat:  totalSeat,
			Registered: registered,
			Remain:     remain,
		}

		var rooms []string
		e.DOM.Find("td:nth-child(7) u").Each(func(i int, s *goquery.Selection) {
			room := strings.TrimSpace(s.Text())
			rooms = append(rooms, room)
		})

		scheduleStr := e.ChildText("td:nth-child(7) > font")
		schedules := parseSchedule(scheduleStr)
		classSchedule := mergeSchedulesWithRooms(schedules, rooms)

		details := ICourseDetails{
			CourseStatus:     courseDetails.CourseStatus,
			CourseCondition:  courseDetails.CourseCondition,
			ContinueCourse:   courseDetails.ContinueCourse,
			EquivalentCourse: courseDetails.EquivalentCourse,
			MidExam:          courseDetails.MidExam,
			FinalExam:        courseDetails.FinalExam,
		}

		uniqueId := uuid.New()

		dataObj := ICourse{
			ID:            uniqueId,
			URL:           urlCourseDetails,
			CourseCode:    courseCode,
			Version:       version,
			CourseNameEN:  courseDetails.CourseNameEN,
			CourseNameTH:  &courseDetails.CourseNameTH,
			Faculty:       courseDetails.Faculty,
			Department:    courseDetails.Department,
			Note:          &note,
			Professors:    professors,
			Credit:        credit,
			Section:       section,
			StatusSection: statusSection,
			Language:      language,
			Degree:        degree,
			ClassSchedule: classSchedule,
			Seat:          seat,
			Details:       details,
		}

		courseData = append(courseData, dataObj)
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	c.Wait()
	result := groupedCourses(courseData)

	return result, err
}
