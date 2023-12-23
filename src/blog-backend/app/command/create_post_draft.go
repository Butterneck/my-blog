package command

import (
	"context"

	"github.com/butterneck/my-blog/src/blog-backend/domain/post"
)

type PostAsset struct {
	Name string
	File []byte
}

type CreatePostDraft struct {
	Title  string
	Body   string
	Assets []PostAsset
}

type CreatePostDraftHandler struct {
	postRepository post.Repository
	postAssetStore post.AssetStore
}

func NewCreatePostDraftHandler(postRepository post.Repository, assetStore post.AssetStore) CreatePostDraftHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	if assetStore == nil {
		panic("assetStore is nil")
	}

	return CreatePostDraftHandler{postRepository: postRepository, postAssetStore: assetStore}
}

func (c CreatePostDraftHandler) Handle(ctx context.Context, cmd CreatePostDraft) error {

	postAssets := []string{}

	// Upload assets to S3
	for _, asset := range cmd.Assets {
		_, err := c.postAssetStore.UploadAsset(ctx, asset.File, asset.Name)
		if err != nil {
			return err
		}

		postAssets = append(postAssets, asset.Name)
	}

	// TODO: Replace on the body all the occurrencies of the old filename with the new one (the path on s3)

	p, err := post.NewPost(cmd.Title, cmd.Body, postAssets)
	if err != nil {
		return err
	}

	return c.postRepository.CreatePost(ctx, p)
}
