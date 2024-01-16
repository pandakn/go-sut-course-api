package scraper

import (
	"regexp"
	"strings"
)

var monthAbbreviations = map[string]string{
	"ม.ค.":  "Jan",
	"ก.พ.":  "Feb",
	"มี.ค.": "Mar",
	"เม.ย.": "Apr",
	"พ.ค.":  "May",
	"มิ.ย.": "Jun",
	"ก.ค.":  "Jul",
	"ส.ค.":  "Aug",
	"ก.ย.":  "Sep",
	"ต.ค.":  "Oct",
	"พ.ย.":  "Nov",
	"ธ.ค.":  "Dec",
}

func convertMonthAbbreviation(thaiMonth string) string {
	return monthAbbreviations[thaiMonth]
}

func cleanString(text string) string {
	return strings.ReplaceAll(text, "\\s+", "")
}

func extractExamInfo(rawText string) *IExamInfo {
	if rawText == "" {
		return nil
	}

	parts := strings.Split(rawText, "เวลา")
	datePart := strings.TrimSpace(parts[0])
	timePart := strings.TrimSpace(parts[1])

	timeRangePattern := regexp.MustCompile(`^\d{2}:\d{2}\s*-\s*\d{2}:\d{2}`)
	timeReplaceWhiteSpace := strings.ReplaceAll(timePart, " ", "")
	extractedTimeRange := timeRangePattern.FindString(timeReplaceWhiteSpace)
	extractedRoom := strings.TrimSpace(strings.Split(timeReplaceWhiteSpace, extractedTimeRange)[1])

	cleanedTimeRange := cleanString(extractedTimeRange)

	dateParts := strings.Fields(datePart)
	date := dateParts[0]
	month := convertMonthAbbreviation(dateParts[1])
	year := dateParts[2]

	return &IExamInfo{
		Date:  date,
		Month: month,
		Times: cleanedTimeRange,
		Year:  year,
		Room:  extractedRoom,
	}
}
