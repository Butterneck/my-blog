package post

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func assertPostPointersDifferent(got, want *Post) (bool, error) {
	if got == want {
		return false, errors.New("pointers are equal")
	}

	return true, nil
}

func assertPostEqual(got, want *Post) (bool, error) {
	if got.title != want.title {
		return false, errors.New("title is not equal: got '" + string(got.title) + "' want '" + string(want.title) + "'")
	}

	if got.body != want.body {
		return false, errors.New("body is not equal: got '" + string(got.body) + "' want '" + string(want.body) + "'")
	}

	if got.slug != want.slug {
		return false, errors.New("slug is not equal: got '" + string(got.slug) + "' want '" + string(want.slug) + "'")
	}

	if len(got.assets) != len(want.assets) {
		return false, errors.New("assets length is not equal: got '" + fmt.Sprint(len(got.assets)) + "' want '" + fmt.Sprint(len(want.assets)) + "'")
	}

	for i, asset := range got.assets {
		if asset != want.assets[i] {
			return false, errors.New("assets is not equal: got '" + string(asset) + "' want '" + string(want.assets[i]) + "'")
		}
	}

	if got.creationDate != want.creationDate {
		return false, errors.New("creationDate is not equal: got '" + fmt.Sprint(got.creationDate) + "' want '" + fmt.Sprint(want.creationDate) + "'")
	}

	if ok, err := assertDraftEqual(&got.draft, &want.draft); !ok {
		return false, err
	}

	return true, nil
}

func newDefaultPost() (*Post, error) {
	return NewPost("Title", "Body", []string{"asset"})
}

func TestNewPost(t *testing.T) {
	want := &Post{
		title:        "",
		body:         "",
		slug:         Slug("title"),
		assets:       []string{},
		creationDate: 0,
		draft: Draft{
			title:  Title("Title"),
			body:   Body("Body"),
			slug:   Slug("title"),
			assets: []string{"asset"},
		},
	}
	got, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	if ok, err := assertPostEqual(got, want); !ok {
		t.Errorf("NewPost() error = '%v'", err)
	}
}

const mockTime int64 = 1577836800

var mockTimeNow = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

func TestPostIsNotPublished(t *testing.T) {
	want := false
	got, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	if got.isPublished() != want {
		t.Errorf("IsPublished() = '%v', want '%v'", got.isPublished(), want)
	}
}

func TestPublishPost(t *testing.T) {
	want := &Post{
		title:        Title("Title"),
		body:         Body("Body"),
		slug:         Slug("title"),
		assets:       []string{"asset"},
		creationDate: mockTime,
		draft: Draft{
			title:  "",
			body:   "",
			slug:   "",
			assets: []string{},
		},
	}
	got, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	got, err = got.Publish(mockTimeNow)
	if err != nil {
		t.Errorf("Publish() error = '%v'", err)
		return
	}

	if ok, err := assertPostEqual(got, want); !ok {
		t.Errorf("Publish() error = '%v'", err)
	}

	if ok, err := assertPostPointersDifferent(got, want); !ok {
		t.Errorf("Publish() error = '%v'", err)
	}
}

func TestPostIsPublished(t *testing.T) {
	want := true
	got, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	got, err = got.Publish(mockTimeNow)
	if err != nil {
		t.Errorf("Publish() error = '%v'", err)
		return
	}

	if got.isPublished() != want {
		t.Errorf("IsPublished() = '%v', want '%v'", got.isPublished(), want)
	}
}

func TestPostHasUnpublishedChanges(t *testing.T) {
	want := true
	got, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	if got.hasUnpublisedChanges() != want {
		t.Errorf("HasUnpublishedChanges() = '%v', want '%v'", got.hasUnpublisedChanges(), want)
	}
}

func TestPostDoesNotHaveUnpublishedChanges(t *testing.T) {
	want := false
	got, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	got, err = got.Publish(mockTimeNow)
	if err != nil {
		t.Errorf("Publish() error = '%v'", err)
		return
	}

	if got.hasUnpublisedChanges() != want {
		t.Errorf("HasUnpublishedChanges() = '%v', want '%v'", got.hasUnpublisedChanges(), want)
	}
}

func TestPublishAlreadyPublishedPost(t *testing.T) {
	want := PostAlreadyPublishedError{}

	newPost, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	PublishedPost, err := newPost.Publish(mockTimeNow)
	if err != nil {
		t.Errorf("Publish() error = '%v'", err)
		return
	}

	PublishedPostBackup := *PublishedPost
	got, err := PublishedPost.Publish(mockTimeNow)
	if err == nil {
		t.Errorf("Publish() error = '%v', want '%v'", err, want)
		return
	}

	if got != nil {
		t.Errorf("Publish() error = '%v', want '%v'", err, want)
		return
	}

	if ok, err := assertPostEqual(PublishedPost, &PublishedPostBackup); !ok {
		t.Errorf("Publish() error = '%v'", err)
	}

	if err != want {
		t.Errorf("Publish() error = '%v', want '%v'", err, want)
		return
	}
}

func TestPostUpdateTitle(t *testing.T) {
	want := &Post{
		title:        "",
		body:         "",
		slug:         Slug("title"),
		assets:       []string{},
		creationDate: 0,
		draft: Draft{
			title:  Title("New Title"),
			body:   Body("Body"),
			slug:   Slug("new-title"),
			assets: []string{"asset"},
		},
	}
	got, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	got, err = got.updateTitle("New Title")
	if err != nil {
		t.Errorf("UpdateTitle() error = '%v'", err)
		return
	}

	if ok, err := assertPostEqual(got, want); !ok {
		t.Errorf("UpdateTitle() error = '%v'", err)
	}

	if ok, err := assertPostPointersDifferent(got, want); !ok {
		t.Errorf("UpdateTitle() error = '%v'", err)
	}
}

func TestPostUpdateBody(t *testing.T) {
	want := &Post{
		title:        "",
		body:         "",
		slug:         Slug("title"),
		assets:       []string{},
		creationDate: 0,
		draft: Draft{
			title:  Title("Title"),
			body:   Body("New Body"),
			slug:   Slug("title"),
			assets: []string{"asset"},
		},
	}
	got, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	got, err = got.updateBody("New Body")
	if err != nil {
		t.Errorf("UpdateBody() error = '%v'", err)
		return
	}

	if ok, err := assertPostEqual(got, want); !ok {
		t.Errorf("UpdateBody() error = '%v'", err)
	}

	if ok, err := assertPostPointersDifferent(got, want); !ok {
		t.Errorf("UpdateBody() error = '%v'", err)
	}
}

func TestPostAddAssets(t *testing.T) {
	want := &Post{
		title:        "",
		body:         "",
		slug:         Slug("title"),
		assets:       []string{},
		creationDate: 0,
		draft: Draft{
			title:  Title("Title"),
			body:   Body("Body"),
			slug:   Slug("title"),
			assets: []string{"asset", "new-asset"},
		},
	}
	got, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}
	got, err = got.AddAssets([]string{"new-asset"})
	if err != nil {
		t.Errorf("AddAsset() error = '%v'", err)
		return
	}

	if ok, err := assertPostEqual(got, want); !ok {
		t.Errorf("AddAsset() error = '%v'", err)
	}

	if ok, err := assertPostPointersDifferent(got, want); !ok {
		t.Errorf("AddAsset() error = '%v'", err)
	}
}

func TestPostRemoveAssets(t *testing.T) {
	want := &Post{
		title:        "",
		body:         "",
		slug:         Slug("title"),
		assets:       []string{},
		creationDate: 0,
		draft: Draft{
			title:  Title("Title"),
			body:   Body("Body"),
			slug:   Slug("title"),
			assets: []string{},
		},
	}
	got, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	got, err = got.RemoveAssets([]string{"asset"})
	if err != nil {
		t.Errorf("RemoveAssets() error = '%v'", err)
		return
	}

	if ok, err := assertPostEqual(got, want); !ok {
		t.Errorf("RemoveAssets() error = '%v'", err)
	}

	if ok, err := assertPostPointersDifferent(got, want); !ok {
		t.Errorf("RemoveAssets() error = '%v'", err)
	}
}

func TestPostUnpublish(t *testing.T) {
	want := &Post{
		title:        "",
		body:         "",
		slug:         Slug("title"),
		assets:       []string{},
		creationDate: 0,
		draft: Draft{
			title:  Title("Title"),
			body:   Body("Body"),
			slug:   Slug("title"),
			assets: []string{"asset"},
		},
	}
	newPost, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	PublishedPost, err := newPost.Publish(mockTimeNow)
	if err != nil {
		t.Errorf("Publish() error = '%v'", err)
		return
	}

	got, err := PublishedPost.Unpublish()
	if err != nil {
		t.Errorf("Unpublish() error = '%v'", err)
		return
	}

	if ok, err := assertPostEqual(got, want); !ok {
		t.Errorf("Unpublish() error = '%v'", err)
	}

	if ok, err := assertPostPointersDifferent(got, want); !ok {
		t.Errorf("Unpublish() error = '%v'", err)
	}
}

func TestPostUnpublishNotPublishedPost(t *testing.T) {
	want := PostNotPublishedError{}

	newPost, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	got, err := newPost.Unpublish()
	if err == nil {
		t.Errorf("Unpublish() error = '%v', want '%v'", err, want)
		return
	}

	if got != nil {
		t.Errorf("Unpublish() error = '%v', want '%v'", err, want)
		return
	}

	if err != want {
		t.Errorf("Unpublish() error = '%v', want '%v'", err, want)
		return
	}
}

func TestPostTitle(t *testing.T) {
	want := Title("Title")
	newPost, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	PublishedPost, err := newPost.Publish(mockTimeNow)
	if err != nil {
		t.Errorf("Publish() error = '%v'", err)
		return
	}

	got := PublishedPost.Title()

	if got != want {
		t.Errorf("Title() = '%v', want '%v'", got, want)
	}
}

func TestPostBody(t *testing.T) {
	want := Body("Body")
	newPost, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	PublishedPost, err := newPost.Publish(mockTimeNow)
	if err != nil {
		t.Errorf("Publish() error = '%v'", err)
		return
	}

	got := PublishedPost.Body()

	if got != want {
		t.Errorf("Body() = '%v', want '%v'", got, want)
	}
}

func TestPostSlug(t *testing.T) {
	want := Slug("new-title")
	newPost, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	newPost, err = newPost.updateTitle("New Title")
	if err != nil {
		t.Errorf("UpdateTitle() error = '%v'", err)
		return
	}

	PublishedPost, err := newPost.Publish(mockTimeNow)
	if err != nil {
		t.Errorf("Publish() error = '%v'", err)
		return
	}

	got := PublishedPost.Slug()

	if got != want {
		t.Errorf("Slug() = '%v', want '%v'", got, want)
	}
}

func TestPostAssets(t *testing.T) {
	want := []string{"asset"}
	newPost, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	PublishedPost, err := newPost.Publish(mockTimeNow)
	if err != nil {
		t.Errorf("Publish() error = '%v'", err)
		return
	}

	got := PublishedPost.Assets()

	if len(got) != len(want) {
		t.Errorf("Assets() length = '%v', want '%v'", len(got), len(want))
	}

	for i, asset := range got {
		if asset != want[i] {
			t.Errorf("Assets() = '%v', want '%v'", got, want)
		}
	}
}

func TestPostCreationDate(t *testing.T) {
	want := mockTime
	newPost, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	PublishedPost, err := newPost.Publish(mockTimeNow)
	if err != nil {
		t.Errorf("Publish() error = '%v'", err)
		return
	}

	got := PublishedPost.CreationDate()

	if got != want {
		t.Errorf("CreationDate() = '%v', want '%v'", got, want)
	}
}

func TestPostDraft(t *testing.T) {
	want := Draft{
		title:  Title("Title"),
		body:   Body("Body"),
		slug:   Slug("title"),
		assets: []string{"asset"},
	}
	newPost, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	got := newPost.Draft()

	if ok, err := assertDraftEqual(&got, &want); !ok {
		t.Errorf("Draft() error = '%v'", err)
	}
}

func TestPostCanBeDeleted(t *testing.T) {
	want := true
	newPost, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}
	got := newPost.CanBeDeleted()

	if got != want {
		t.Errorf("CanBeDeleted() = '%v', want '%v'", got, want)
	}
}

func TestPostCanNotBeDeleted(t *testing.T) {
	want := false
	newPost, err := newDefaultPost()
	if err != nil {
		t.Errorf("NewPost() error = '%v'", err)
		return
	}

	PublishedPost, err := newPost.Publish(mockTimeNow)
	if err != nil {
		t.Errorf("Publish() error = '%v'", err)
		return
	}
	got := PublishedPost.CanBeDeleted()

	if got != want {
		t.Errorf("CanBeDeleted() = '%v', want '%v'", got, want)
	}
}
