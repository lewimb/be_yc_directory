package pkg

import (
	"regexp"
	"strings"
)

func CreateSlug(title string) string {
	reg, err := regexp.Compile("[^a-zA-z0-9]+")

	if err != nil {
		panic(err)
	}

	processedString := reg.ReplaceAllString(title, " ")

	processedString = strings.TrimSpace(processedString)

	slug := strings.ReplaceAll(processedString, " ", " _")

	slug = strings.ToLower(slug)

	return slug
}
