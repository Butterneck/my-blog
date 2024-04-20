package post

import (
	"errors"
	"fmt"
	"testing"
)

func assertDraftEqual(got, want *Draft) (bool, error) {
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

	return true, nil
}

func assertDraftPointersDifferent(got, want *Draft) (bool, error) {
	if got == want {
		return false, errors.New("pointers are equal")
	}

	return true, nil
}

func TestNewDraft(t *testing.T) {
	want := &Draft{
		title:  Title("Title"),
		body:   Body("Body"),
		slug:   Slug("title"),
		assets: []string{},
	}
	got, err := newDraft("Title", "Body", []string{})
	if err != nil {
		t.Errorf("NewDraft() error = '%v'", err)
		return
	}

	if ok, err := assertDraftEqual(got, want); !ok {
		t.Errorf("NewDraft() error = '%v'", err)
	}
}

func TestDraftUpdateTitle(t *testing.T) {
	want := &Draft{
		title:  Title("New title"),
		body:   Body("Body"),
		slug:   Slug("new-title"),
		assets: []string{},
	}
	d, err := newDraft("Title", "Body", []string{})
	if err != nil {
		t.Errorf("NewDraft() error = '%v'", err)
		return
	}

	got, err := d.updateTitle("New title")
	if err != nil {
		t.Errorf("UpdateTitle() error = '%v'", err)
		return
	}

	if ok, err := assertDraftEqual(got, want); !ok {
		t.Errorf("UpdateTitle() error = '%v'", err)
	}

	if ok, err := assertDraftPointersDifferent(got, d); !ok {
		t.Errorf("UpdateTitle() error = '%v'", err)
	}
}

func TestDraftUpdateBody(t *testing.T) {
	want := &Draft{
		title:  Title("Title"),
		body:   Body("New body"),
		slug:   Slug("title"),
		assets: []string{},
	}
	d, err := newDraft("Title", "Body", []string{})
	if err != nil {
		t.Errorf("NewDraft() error = '%v'", err)
		return
	}

	got, err := d.updateBody("New body")
	if err != nil {
		t.Errorf("UpdateBody() error = '%v'", err)
		return
	}

	if ok, err := assertDraftEqual(got, want); !ok {
		t.Errorf("UpdateBody() error = '%v'", err)
	}

	if ok, err := assertDraftPointersDifferent(got, d); !ok {
		t.Errorf("UpdateBody() error = '%v'", err)
	}
}

func TestDraftAddAssets(t *testing.T) {
	want := &Draft{
		title:  Title("Title"),
		body:   Body("Body"),
		slug:   Slug("title"),
		assets: []string{"asset"},
	}
	d, err := newDraft("Title", "Body", []string{})
	if err != nil {
		t.Errorf("NewDraft() error = '%v'", err)
		return
	}

	got, err := d.addAssets([]string{"asset"})
	if err != nil {
		t.Errorf("AddAssets() error = '%v'", err)
		return
	}

	if ok, err := assertDraftEqual(got, want); !ok {
		t.Errorf("AddAssets() error = '%v'", err)
	}

	if ok, err := assertDraftPointersDifferent(got, d); !ok {
		t.Errorf("AddAssets() error = '%v'", err)
	}
}

func TestDraftRemoveAssets(t *testing.T) {
	want := &Draft{
		title:  Title("Title"),
		body:   Body("Body"),
		slug:   Slug("title"),
		assets: []string{},
	}

	d, err := newDraft("Title", "Body", []string{"asset"})
	if err != nil {
		t.Errorf("NewDraft() error = '%v'", err)
		return
	}

	got, err := d.removeAssets([]string{"asset"})
	if err != nil {
		t.Errorf("RemoveAssets() error = '%v'", err)
		return
	}

	if ok, err := assertDraftEqual(got, want); !ok {
		t.Errorf("RemoveAssets() error = '%v'", err)
	}

	if ok, err := assertDraftPointersDifferent(got, d); !ok {
		t.Errorf("RemoveAssets() error = '%v'", err)
	}
}

func TestDraftTitle(t *testing.T) {
	want := Title("Title")
	d, err := newDraft("Title", "Body", []string{})
	if err != nil {
		t.Errorf("NewDraft() error = '%v'", err)
		return
	}

	got := d.Title()

	if got != want {
		t.Errorf("Title() = '%v', want '%v'", got, want)
	}
}

func TestDraftBody(t *testing.T) {
	want := Body("Body")
	d, err := newDraft("Title", "Body", []string{})
	if err != nil {
		t.Errorf("NewDraft() error = '%v'", err)
		return
	}

	got := d.Body()

	if got != want {
		t.Errorf("Body() = '%v', want '%v'", got, want)
	}
}

func TestDraftSlug(t *testing.T) {
	want := Slug("title")
	d, err := newDraft("Title", "Body", []string{})
	if err != nil {
		t.Errorf("NewDraft() error = '%v'", err)
		return
	}

	got := d.Slug()

	if got != want {
		t.Errorf("Slug() = '%v', want '%v'", got, want)
	}
}

func TestDraftAssets(t *testing.T) {
	want := []string{"asset"}

	d, err := newDraft("Title", "Body", []string{"asset"})
	if err != nil {
		t.Errorf("NewDraft() error = '%v'", err)
		return
	}

	got := d.Assets()

	if len(got) != len(want) {
		t.Errorf("Assets() length = '%v', want '%v'", len(got), len(want))
		return
	}

	for i, asset := range got {
		if asset != want[i] {
			t.Errorf("Assets() = '%v', want '%v'", got, want)
		}
	}
}
