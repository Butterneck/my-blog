package asset

// import "strings"

// type Asset string

// func newAsset(postSlug Slug, name string) (Asset, error) {
// 	if postSlug == "" {
// 		return "", EmptyPostSlugError{}
// 	}

// 	if name == "" {
// 		return "", EmptyNameError{}
// 	}

// 	if strings.Contains(name, "/") {
// 		return "", InvalidNameError{}
// 	}

// 	return Asset(strings.Join([]string{string(postSlug), "draft", name}, "/")), nil
// }

// type EmptyPostSlugError struct{}

// func (e EmptyPostSlugError) Error() string {
// 	return "empty post slug"
// }

// type EmptyNameError struct{}

// func (e EmptyNameError) Error() string {
// 	return "empty name"
// }

// type InvalidNameError struct{}

// func (e InvalidNameError) Error() string {
// 	return "invalid name: name cannot contain '/'"
// }

// func (a Asset) publish() (*Asset, error) {
// 	asset := Asset(strings.Replace(string(a), "draft", "published", 1))
// 	return &asset, nil
// }

// func (a Asset) unpublish() (*Asset, error) {
// 	asset := Asset(strings.Replace(string(a), "published", "draft", 1))
// 	return &asset, nil
// }
