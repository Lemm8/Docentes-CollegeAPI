package validators

import (
	"regexp"
)

func IsValidPath(path string) bool {

	if path != "/Docentes" {
		return false
	}

	return true
}

func IsValidDate(date string) bool {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	return re.MatchString(date)
}
