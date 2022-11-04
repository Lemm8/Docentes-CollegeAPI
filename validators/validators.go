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
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d)")
	return re.MatchString(date)
}
