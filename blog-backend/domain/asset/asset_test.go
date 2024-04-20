package asset

// import (
// 	"errors"
// 	"testing"
// )

// func assertAssetPointersDifferent(got, want *Asset) (bool, error) {
// 	if got == want {
// 		return false, errors.New("pointers are equal")
// 	}

// 	return true, nil
// }

// func TestNewAsset(t *testing.T) {
// 	want := Asset("post-slug/draft/name")
// 	got, err := newAsset("post-slug", "name")
// 	if err != nil {
// 		t.Errorf("newAsset() error = '%v'", err)
// 		return
// 	}

// 	if got != want {
// 		t.Errorf("newAsset() = '%v', want '%v'", got, want)
// 	}
// }

// func TestNewAssetEmptyPostSlug(t *testing.T) {
// 	want := EmptyPostSlugError{}
// 	_, err := newAsset("", "name")
// 	if err == nil {
// 		t.Errorf("newAsset() error = '%v', want '%v'", err, want)
// 		return
// 	}

// 	if err != want {
// 		t.Errorf("newAsset() error = '%v', want '%v'", err, want)
// 	}
// }

// func TestNewAssetEmptyName(t *testing.T) {
// 	want := EmptyNameError{}
// 	_, err := newAsset("post-slug", "")
// 	if err == nil {
// 		t.Errorf("newAsset() error = '%v', want '%v'", err, want)
// 		return
// 	}

// 	if err != want {
// 		t.Errorf("newAsset() error = '%v', want '%v'", err, want)
// 	}
// }

// func TestNewAssetWithInvalidName(t *testing.T) {
// 	want := InvalidNameError{}
// 	_, err := newAsset("post-slug", "name/name")
// 	if err == nil {
// 		t.Errorf("newAsset() error = '%v', want '%v'", err, want)
// 		return
// 	}

// 	if err != want {
// 		t.Errorf("newAsset() error = '%v', want '%v'", err, want)
// 	}
// }

// func TestPublishAsset(t *testing.T) {
// 	want := Asset("post-slug/published/name")

// 	asset := Asset("post-slug/draft/name")
// 	got, err := asset.publish()
// 	if err != nil {
// 		t.Errorf("PublishAsset() error = '%v'", err)
// 		return
// 	}

// 	if *got != want {
// 		t.Errorf("PublishAsset() = '%v', want '%v'", got, want)
// 	}

// 	if ok, err := assertAssetPointersDifferent(got, &asset); !ok {
// 		t.Errorf("PublishAsset() error = '%v'", err)
// 	}
// }

// func TestPublishPublishedAsset(t *testing.T) {
// 	want := Asset("post-slug/published/name")

// 	asset := Asset("post-slug/published/name")
// 	got, err := asset.publish()
// 	if err != nil {
// 		t.Errorf("PublishAsset() error = '%v'", err)
// 		return
// 	}

// 	if *got != want {
// 		t.Errorf("PublishAsset() = '%v', want '%v'", got, want)
// 	}

// 	if ok, err := assertAssetPointersDifferent(got, &asset); !ok {
// 		t.Errorf("PublishAsset() error = '%v'", err)
// 	}
// }

// func TestUnpublishAsset(t *testing.T) {
// 	want := Asset("post-slug/draft/name")

// 	asset := Asset("post-slug/published/name")
// 	got, err := asset.unpublish()
// 	if err != nil {
// 		t.Errorf("UnpublishAsset() error = '%v'", err)
// 		return
// 	}

// 	if *got != want {
// 		t.Errorf("UnpublishAsset() = '%v', want '%v'", got, want)
// 	}

// 	if ok, err := assertAssetPointersDifferent(got, &asset); !ok {
// 		t.Errorf("UnpublishAsset() error = '%v'", err)
// 	}
// }

// func TestUnpublishDraftAsset(t *testing.T) {
// 	want := Asset("post-slug/draft/name")

// 	asset := Asset("post-slug/draft/name")
// 	got, err := asset.unpublish()
// 	if err != nil {
// 		t.Errorf("UnpublishAsset() error = '%v'", err)
// 		return
// 	}

// 	if *got != want {
// 		t.Errorf("UnpublishAsset() = '%v', want '%v'", got, want)
// 	}

// 	if ok, err := assertAssetPointersDifferent(got, &asset); !ok {
// 		t.Errorf("UnpublishAsset() error = '%v'", err)
// 	}
// }
