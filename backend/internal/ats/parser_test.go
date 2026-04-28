package ats

import "testing"

func TestParseCV(t *testing.T) {

	input := `
Name: John Doe
Email: john@test.com
Skills: Go, Docker
Experience: Backend Developer
Education: Computer Science
`

	cv := ParseCV(input)

	if cv.Name != "John Doe" {
		t.Errorf("expected John Doe, got %s", cv.Name)
	}

	if len(cv.Skills) == 0 {
		t.Error("skills not parsed")
	}
}