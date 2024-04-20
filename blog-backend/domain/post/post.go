package post

import "time"

type Post struct {
	title        Title
	body         Body
	slug         Slug
	creationDate int64
	draft        Draft
	assets       []string
}

type PostAlreadyPublishedError struct{}

func (e PostAlreadyPublishedError) Error() string {
	return "post already published"
}

func NewPost(title, body string, assets []string) (*Post, error) {

	draft, err := newDraft(title, body, []string{})
	if err != nil {
		return nil, err
	}

	draft, err = draft.addAssets(assets)
	if err != nil {
		return nil, err
	}

	post := &Post{
		draft: *draft,
		slug:  draft.slug,
	}

	return post, nil
}

func (p Post) isPublished() bool {
	return p.title != "" && p.body != "" && p.creationDate != 0
}

func (p Post) hasUnpublisedChanges() bool {
	return p.draft.title != "" || p.draft.body != "" || p.draft.assets != nil
}

func (p Post) Publish(time time.Time) (*Post, error) {
	if !p.hasUnpublisedChanges() {
		return nil, PostAlreadyPublishedError{}
	}

	p.title = p.draft.title
	p.body = p.draft.body
	p.assets = p.draft.assets

	// Set only on first publish
	if !p.isPublished() {
		p.slug = p.draft.slug
	}

	// Set only on first publish
	if !p.isPublished() {
		p.creationDate = time.Unix()
	}

	p.draft = Draft{}

	return &p, nil
}

func (p Post) updateTitle(title string) (*Post, error) {
	updatedDraft, err := p.draft.updateTitle(title)
	if err != nil {
		return nil, err
	}

	p.draft = *updatedDraft
	return &p, nil
}

func (p Post) updateBody(body string) (*Post, error) {
	updatedDraft, err := p.draft.updateBody(body)
	if err != nil {
		return nil, err
	}

	p.draft = *updatedDraft
	return &p, nil
}

func (p Post) AddAssets(assets []string) (*Post, error) {

	updatedDraft, err := p.draft.addAssets(assets)
	if err != nil {
		return nil, err
	}

	p.draft = *updatedDraft
	return &p, nil
}

func (p Post) RemoveAssets(assets []string) (*Post, error) {

	updatedDraft, err := p.draft.removeAssets(assets)
	if err != nil {
		return nil, err
	}

	p.draft = *updatedDraft
	return &p, nil
}

func (p Post) Unpublish() (*Post, error) {
	if !p.isPublished() {
		return nil, PostNotPublishedError{}
	}

	updatedDraft, err := newDraft(string(p.title), string(p.body), p.assets)
	if err != nil {
		return nil, err
	}

	p.draft = *updatedDraft
	p.title = ""
	p.body = ""
	p.creationDate = 0
	p.assets = nil

	return &p, nil
}

type PostNotPublishedError struct{}

func (e PostNotPublishedError) Error() string {
	return "post not published"
}

func (p Post) CanBeDeleted() bool {
	return !p.isPublished()
}

func (p Post) Title() Title {
	return p.title
}

func (p Post) Body() Body {
	return p.body
}

func (p Post) Slug() Slug {
	return p.slug
}

func (p Post) Assets() []string {
	return p.assets
}

func (p Post) CreationDate() int64 {
	return p.creationDate
}

func (p Post) Draft() Draft {
	return p.draft
}
