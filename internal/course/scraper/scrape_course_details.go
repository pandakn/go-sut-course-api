package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type courseDetail struct {
	CourseNameTH     string
	CourseNameEN     string
	Faculty          string
	Department       string
	CourseStatus     string
	CourseCondition  *[]string
	ContinueCourse   *[]string
	EquivalentCourse *[]string
	MidExam          *IExamInfo
	FinalExam        *IExamInfo
}

func extractCourse(e *colly.HTMLElement, selector string) *[]string {
	var data []string

	e.ForEach(selector, func(i int, h *colly.HTMLElement) {
		result := h.DOM.NextFiltered("td").Find("a").Map(func(i int, s *goquery.Selection) string {
			return s.Text()
		})
		data = append(data, result...)
	})

	return &data
}

func extractExamAndRemark(e *colly.HTMLElement, selector string) string {
	data := strings.TrimSpace(e.DOM.Find(selector).Next().Text())
	return data
}

func cleanText(text string) string {
	return strings.Join(strings.Fields(text), " ")
}

func scrapeCourseDetails(courseCodeUrl string) (*courseDetail, error) {
	courseDetail := &courseDetail{}
	c := colly.NewCollector(
		colly.DetectCharset(),
		colly.Async(true),
	)

	c.OnRequest(func(r *colly.Request) {
		r.ResponseCharacterEncoding = "windows-874"
	})

	c.OnHTML("td:nth-child(3) > table:nth-child(2) > tbody > tr > td:nth-child(2) > table > tbody", func(e *colly.HTMLElement) {
		courseDetail.CourseNameEN = cleanText(e.ChildText("tr:nth-child(1) > td:nth-child(2) > b > font"))
		courseDetail.CourseNameTH = cleanText(e.ChildText("tr:nth-child(2) > td:nth-child(2) > font"))

		facultyRaw := strings.Split(e.ChildText("tr:nth-child(3) > td:nth-child(3) > font"), ", ")
		courseDetail.Faculty = strings.TrimSpace(facultyRaw[0])
		courseDetail.Department = strings.TrimSpace(facultyRaw[1])

		courseDetail.CourseStatus = cleanText(e.ChildText("tr:nth-child(5) > td:nth-child(3) > font"))

		courseDetail.CourseCondition = extractCourse(e, `td:contains("เงื่อนไขรายวิชา")`)
		courseDetail.ContinueCourse = extractCourse(e, `td:contains("รายวิชาต่อเนื่อง")`)
		courseDetail.EquivalentCourse = extractCourse(e, `td:contains("รายวิชาเทียบเท่า")`)
	})

	c.OnHTML("td:nth-child(3) > table:nth-child(5)", func(e *colly.HTMLElement) {
		midExam := extractExamAndRemark(e, `td:contains("สอบกลางภาค")`)
		finalExam := extractExamAndRemark(e, `td:contains("สอบประจำภาค")`)

		courseDetail.MidExam = extractExamInfo(midExam)
		courseDetail.FinalExam = extractExamInfo(finalExam)
	})

	err := c.Visit(courseCodeUrl)
	if err != nil {
		return nil, err
	}

	c.Wait()
	return courseDetail, nil
}
