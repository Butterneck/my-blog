package post

import (
	"strings"
	"testing"
	"testing/quick"
)

func TestNewTitle(t *testing.T) {
	want := Title("title")
	got, err := newTitle("title")
	if err != nil {
		t.Errorf("newTitle() error = '%v'", err)
		return
	}

	if got != want {
		t.Errorf("newTitle() = '%v', want '%v'", got, want)
	}
}

func TestNewTitleEmpty(t *testing.T) {
	want := EmptyTitleError{}
	_, err := newTitle("")
	if err == nil {
		t.Errorf("newTitle() error = '%v', want '%v'", err, want)
		return
	}

	if err != want {
		t.Errorf("newTitle() error = '%v', want '%v'", err, want)
	}
}

func TestNewTitleTooLong(t *testing.T) {
	want := TooLongTitleError{
		MaxTitleLength:      100,
		ProvidedTitleLength: 101,
	}
	_, err := newTitle(strings.Repeat("a", 101))
	if err == nil {
		t.Errorf("newTitle() error = '%v', want '%v'", err, want)
		return
	}

	if err != want {
		t.Errorf("newTitle() error = '%v', want '%v'", err, want)
	}
}

func TestNewBody(t *testing.T) {
	want := Body("body")
	got, err := newBody("body")
	if err != nil {
		t.Errorf("newBody() error = '%v'", err)
		return
	}

	if got != want {
		t.Errorf("newBody() = '%v', want '%v'", got, want)
	}
}

func TestNewBodyEmpty(t *testing.T) {
	want := EmptyBodyError{}
	_, err := newBody("")
	if err == nil {
		t.Errorf("newBody() error = '%v', want '%v'", err, want)
		return
	}

	if err != want {
		t.Errorf("newBody() error = '%v', want '%v'", err, want)
	}
}

func TestNewSlug(t *testing.T) {

	cases := []struct {
		Description string
		Title       Title
		Want        Slug
	}{
		{"'Title' gets converted to 'title'", "Title", "title"},
		{"'Example title with spaces' gets converted to 'example-title-with-spaces'", "Example title with spaces", "example-title-with-spaces"},
		{"'Example title with special characters*#&' gets converted to 'example-title-with-special-characters'", "Example title with special characters *#&", "example-title-with-special-characters"},
		{"'Example title with special characters *#&' gets converted to 'example-title-with-special-characters'", "Example title with special characters *#&", "example-title-with-special-characters"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := newSlug(test.Title)
			if err != nil {
				t.Errorf("newSlug() error = '%v'", err)
				return
			}
			if got != test.Want {
				t.Errorf("got '%q', want '%q'", got, test.Want)
			}
		})
	}
}

func TestPropertiesOfSlug(t *testing.T) {
	cases := []struct {
		Description string
		Assertion   func(title Title) bool
	}{
		{
			"slug should be lowercase",
			func(title Title) bool {
				slug, _ := newSlug(title)
				return slug == Slug(strings.ToLower(string(slug)))
			},
		},
		{
			"slug should not contain spaces",
			func(title Title) bool {
				slug, _ := newSlug(title)
				return !strings.Contains(string(slug), " ")
			},
		},
		{
			"slug should not contain special characters",
			func(title Title) bool {
				slug, _ := newSlug(title)
				return !strings.ContainsAny(string(slug), "*#&!@^%$?\\t/|,.")
			},
		},
		{
			"slug should not have trailing hyphens",
			func(title Title) bool {
				slug, _ := newSlug(title)
				return !strings.HasSuffix(string(slug), "-")
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			if err := quick.Check(test.Assertion, nil); err != nil {
				t.Error("failed checks", err)
			}
		})
	}
}
