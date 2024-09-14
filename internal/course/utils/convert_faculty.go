package utils

import (
	"errors"
	"strings"
)

type convertFacultyIdErr string

func (e convertFacultyIdErr) Error() string {
	return string(e)
}

const (
	facultyIdErr convertFacultyIdErr = "invalid faculty"
)

type FacultyId string

const (
	All                    FacultyId = "all"
	Science                FacultyId = "10100"
	SocialTechnology       FacultyId = "10200"
	AgriculturalTechnology FacultyId = "10300"
	Medicine               FacultyId = "10600"
	Engineering            FacultyId = "10700"
	Nursing                FacultyId = "10800"
	Dentistry              FacultyId = "10900"
	PublicHealth           FacultyId = "11000"
	DigitalArtsAndScience  FacultyId = "11100"
)

var facultyIdMap = map[string]FacultyId{
	"ALL":                      All,
	"SCIENCE":                  Science,
	"SOCIAL_TECHNOLOGY":        SocialTechnology,
	"AGRICULTURAL_TECHNOLOGY":  AgriculturalTechnology,
	"MEDICINE":                 Medicine,
	"ENGINEERING":              Engineering,
	"NURSING":                  Nursing,
	"DENTISTRY":                Dentistry,
	"PUBLIC_HEALTH":            PublicHealth,
	"DIGITAL_ARTS_AND_SCIENCE": DigitalArtsAndScience,
}

func ConvertFacultyToFacultyId(faculty string) (*string, error) {
	facultyName := strings.ToUpper(strings.TrimSpace(faculty))

	facultyIdValue, found := facultyIdMap[facultyName]
	if !found {
		return nil, errors.New(string(facultyIdErr))
	}

	facultyIdStr := string(facultyIdValue)
	return &facultyIdStr, nil
}
