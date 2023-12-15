package post

import (
	"fmt"
	"time"
)

type Post struct {
	title        Title
	body         Body
	slug         Slug
	creationDate int64
	draft        Draft
}

type PostAdapter struct {
	Title        Title
	Body         Body
	Slug         Slug
	CreationDate int64
	Draft        DraftAdapter
}

func NewPost(title, body string) (*Post, error) {

	draft, err := newDraft(title, body)
	if err != nil {
		return nil, fmt.Errorf("NewPost - error: %v", err)
	}

	post := &Post{
		draft: *draft,
		slug:  draft.slug,
	}

	return post, nil
}

func (p *Post) IsPublished() bool {
	return p.title != "" && p.body != "" && p.creationDate != 0
}

func (p *Post) HasUnpublishedChanges() bool {
	return p.draft.title != "" && p.draft.body != ""
}

func (p *Post) Publish() error {
	p.title = p.draft.title
	p.body = p.draft.body

	// Set only on first publish
	if !p.IsPublished() {
		p.slug = p.draft.slug
	}

	// Set only on first publish
	if !p.IsPublished() {
		p.creationDate = time.Now().Unix()
	}

	p.draft = Draft{}

	return nil
}

func (p *Post) UpdateTitle(title string) error {
	_, err := p.draft.updateTitle(title)
	if err != nil {
		return fmt.Errorf("UpdateTitle - error: %v", err)
	}
	return nil
}

func (p *Post) UpdateBody(body string) error {
	_, err := p.draft.updateBody(body)
	if err != nil {
		return fmt.Errorf("UpdateBody - error: %v", err)
	}
	return nil
}

func (p *Post) Unpublish() error {
	draft, err := newDraft(string(p.title), string(p.body))
	if err != nil {
		return fmt.Errorf("Unpublish - error: %v", err)
	}

	p.draft = *draft

	p.title = ""
	p.body = ""
	p.creationDate = 0

	return nil
}

func (p *Post) Title() string {
	return string(p.title)
}

func (p *Post) Body() string {
	return string(p.body)
}

func (p *Post) Slug() string {
	return string(p.slug)
}

func (p *Post) CreationDate() int64 {
	return p.creationDate
}

func (p *Post) Draft() *Draft {
	return &p.draft
}

func (p *Post) String() string {
	return fmt.Sprintf("Post{title: %s, body: %s, slug: %s, creationDate: %d, draft: %v}", p.title, p.body, p.slug, p.creationDate, p.draft)
}

// This method is used to load objects from various adapters who need to provide a PostAdapter object
// DO not use this method to create a new Post object, it may not be in a valid state, use NewPost instead
func LoadPostFromAdapter(post PostAdapter) (*Post, error) {
	return &Post{
		title:        post.Title,
		body:         post.Body,
		slug:         post.Slug,
		creationDate: post.CreationDate,
		draft: Draft{
			title: post.Draft.Title,
			body:  post.Draft.Body,
			slug:  post.Draft.Slug,
		},
	}, nil
}
