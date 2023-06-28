package viewer

import (
	"fmt"
	"strconv"
)

type ViewTemplate struct {
	Cover string
	Title string
}

const ViewTemplatePath = "view/tmpl.html"

func ValidateYyyy(yyyy string) string {
	i, err := strconv.Atoi(yyyy)
	if err != nil {
		return ""
	}
	if i >= 1970 {
		return yyyy
	}
	return ""
}

func ValidateMm(mm string) string {
	i, err := strconv.Atoi(mm)
	if err != nil {
		return ""
	}
	if 1 <= i && i <= 12 {
		return fmt.Sprintf("%02d", int64(i))
	}

	return ""
}
