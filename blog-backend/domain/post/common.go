package post

import (
	"fmt"
	"regexp"
	"strings"
)

type Title string

func newTitle(title string) (Title, error) {

	if title == "" {
		return "", EmptyTitleError{}
	}

	if len(title) > 100 {
		return "", TooLongTitleError{
			MaxTitleLength:      100,
			ProvidedTitleLength: len(title),
		}
	}

	return Title(title), nil
}

func (t Title) ToString() string {
	return string(t)
}

type EmptyTitleError struct{}

func (e EmptyTitleError) Error() string {
	return "empty title not allowed"
}

type TooLongTitleError struct {
	MaxTitleLength      int
	ProvidedTitleLength int
}

func (e TooLongTitleError) Error() string {
	return fmt.Sprintf(
		"too long title, max title length: %d, provided title length: %d",
		e.MaxTitleLength,
		e.ProvidedTitleLength,
	)
}

type Body string

func newBody(body string) (Body, error) {
	if body == "" {
		return "", EmptyBodyError{}
	}

	return Body(body), nil
}

func (b Body) ToString() string {
	return string(b)
}

type EmptyBodyError struct{}

func (e EmptyBodyError) Error() string {
	return "empty body not allowed"
}

type Slug string

func newSlug(title Title) (Slug, error) {
	// Convert to lowercase
	slug := strings.ToLower(string(title))

	// Replace spaces with hyphens
	re, err := regexp.Compile(`\s+`)
	if err != nil {
		return "", fmt.Errorf("NewSlug - regexp.Compile - error: %v", err)
	}
	slug = re.ReplaceAllString(slug, "-")

	// Remove special characters, leaving only letters, numbers, hyphens, and underscores
	re, err = regexp.Compile(`[^\w-]+`)
	if err != nil {
		return "", fmt.Errorf("NewSlug - regexp.Compile - error: %v", err)
	}
	slug = re.ReplaceAllString(slug, "")

	// Remove trailing hyphens
	re, err = regexp.Compile(`-+$`)
	if err != nil {
		return "", fmt.Errorf("NewSlug - regexp.Compile - error: %v", err)
	}
	slug = re.ReplaceAllString(slug, "")

	return Slug(slug), nil
}

func (s Slug) ToString() string {
	return string(s)
}
