package command

import (
	"context"
	"fmt"

	"github.com/butterneck/my-blog/blog-backend/domain/asset"
	"github.com/butterneck/my-blog/blog-backend/domain/post"
)

type DeletePost struct {
	Slug string
}

type DeletePostHandler struct {
	postRepository  post.Repository
	assetRepository asset.AssetRepository
}

func NewDeletePostHandler(postRepository post.Repository, assetRepository asset.AssetRepository) DeletePostHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	if assetRepository == nil {
		panic("assetRepository is nil")
	}

	return DeletePostHandler{postRepository: postRepository, assetRepository: assetRepository}
}

func (c DeletePostHandler) Handle(ctx context.Context, cmd DeletePost) error {

	p, err := c.postRepository.GetAnyPost(ctx, cmd.Slug)
	if err != nil {
		return err
	}

	// Check if post can be deleted
	if !p.CanBeDeleted() {
		return fmt.Errorf("post %s cannot be deleted", cmd.Slug)
	}

	// Delete draft assets from S3
	for _, asset := range p.Draft().Assets() {
		err := c.assetRepository.DeleteAsset(ctx, asset)
		if err != nil {
			return err
		}
	}

	return c.postRepository.DeletePost(ctx, cmd.Slug)
}
