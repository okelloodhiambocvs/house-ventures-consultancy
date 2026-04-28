package ats

import (
	"strings"

	"github.com/okelloodhiambocvs/house-ventures-consultancy/backend/internal/models"
)

func ParseCV(input string) models.CV {
	lines := strings.Split(input, "\n")

	cv := models.CV{
		Skills:     []string{},
		Experience: []string{},
		Education:  []string{},
	}

	for _, line := range lines {
		l := strings.ToLower(line)

		if strings.Contains(l, "name") {
			cv.Name = extract(line)
		}

		if strings.Contains(l, "email") {
			cv.Email = extract(line)
		}

		if strings.Contains(l, "skill") {
			cv.Skills = append(cv.Skills, extract(line))
		}

		if strings.Contains(l, "experience") {
			cv.Experience = append(cv.Experience, extract(line))
		}

		if strings.Contains(l, "education") {
			cv.Education = append(cv.Education, extract(line))
		}
	}

	return cv
}

func extract(line string) string {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[1])
	}
	return ""
}