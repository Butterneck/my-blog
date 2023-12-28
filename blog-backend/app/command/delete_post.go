package command

import (
	"context"

	"github.com/butterneck/my-blog/blog-backend/domain/post"
)

type DeletePost struct {
	Slug string
}

type DeletePostHandler struct {
	postRepository post.Repository
	postAssetStore post.AssetStore
}

func NewDeletePostHandler(postRepository post.Repository, assetStore post.AssetStore) DeletePostHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	if assetStore == nil {
		panic("assetStore is nil")
	}

	return DeletePostHandler{postRepository: postRepository, postAssetStore: assetStore}
}

func (c DeletePostHandler) Handle(ctx context.Context, cmd DeletePost) error {

	p, err := c.postRepository.GetAnyPost(ctx, cmd.Slug)
	if err != nil {
		return err
	}

	// Delete draft assets from S3
	for _, asset := range p.Draft().Assets() {
		err := c.postAssetStore.DeleteAsset(ctx, asset)
		if err != nil {
			return err
		}
	}

	return c.postRepository.DeletePost(ctx, cmd.Slug)
}
