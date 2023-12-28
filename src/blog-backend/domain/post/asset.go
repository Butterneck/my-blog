package post

import (
	"regexp"
)

type Asset string

func NewAsset(postSlug, name string) (Asset, error) {
	return Asset(postSlug + "/draft" + name), nil
}

func PublishAsset(asset Asset) Asset {
	replacement := "published"

	// Create a regular expression to match "draft"
	regex := regexp.MustCompile(`draft`)

	// Use the ReplaceAllString method to replace "draft" with the replacement string
	return Asset(regex.ReplaceAllString(string(asset), replacement))

}

func UnpublishAsset(asset Asset) Asset {
	replacement := "draft"

	// Create a regular expression to match "published"
	regex := regexp.MustCompile(`published`)

	// Use the ReplaceAllString method to replace "published" with the replacement string
	return Asset(regex.ReplaceAllString(string(asset), replacement))
}
