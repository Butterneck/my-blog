package post

import (
	"fmt"
)

type Draft struct {
	title Title
	body  Body
	slug  Slug
}

type DraftAdapter struct {
	Title Title `json:"title"`
	Body  Body  `json:"body"`
	Slug  Slug  `json:"slug"`
}

func newDraft(title, body string) (*Draft, error) {
	var err error
	var d Draft

	d.title, err = newTitle(title)
	if err != nil {
		return nil, fmt.Errorf("NewDraft - error: %v", err)
	}

	d.slug, err = newSlug(d.title)
	if err != nil {
		return nil, fmt.Errorf("NewDraft - error: %v", err)
	}

	d.body, err = newBody(string(body))
	if err != nil {
		return nil, fmt.Errorf("NewDraft - error: %v", err)
	}

	return &d, nil
}

func (d *Draft) updateTitle(title string) (*Draft, error) {
	var err error
	d.title, err = newTitle(title)
	if err != nil {
		return d, fmt.Errorf("UpdateTitle - error: %v", err)
	}

	d.slug, err = newSlug(d.title)
	if err != nil {
		return d, fmt.Errorf("UpdateTitle - error: %v", err)
	}

	return d, nil
}

func (d *Draft) updateBody(body string) (*Draft, error) {
	var err error
	d.body, err = newBody(body)
	if err != nil {
		return d, fmt.Errorf("UpdateBody - error: %v", err)
	}

	return d, nil
}

func (d *Draft) Title() string {
	return string(d.title)
}

func (d *Draft) Body() string {
	return string(d.body)
}

func (d *Draft) Slug() string {
	return string(d.slug)
}
