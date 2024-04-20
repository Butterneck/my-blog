package command

import (
	"context"
	"time"

	"github.com/butterneck/my-blog/blog-backend/domain/asset"
	"github.com/butterneck/my-blog/blog-backend/domain/post"
)

type PublishPostDraft struct {
	Slug string
}

type PublishPostDraftHandler struct {
	postRepository  post.Repository
	assetRepository asset.AssetRepository
}

func NewPublishPostDraftHandler(postRepository post.Repository, assetRepository asset.AssetRepository) PublishPostDraftHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	if assetRepository == nil {
		panic("assetRepository is nil")
	}

	return PublishPostDraftHandler{postRepository: postRepository, assetRepository: assetRepository}
}

func (c PublishPostDraftHandler) Handle(ctx context.Context, cmd PublishPostDraft) error {

	p, err := c.postRepository.GetAnyPost(ctx, cmd.Slug)
	if err != nil {
		return err
	}

	// Publish assets
	for _, asset := range p.Draft().Assets() {
		assetName, err := c.assetRepository.PublishAsset(ctx, asset)
		if err != nil {
			return err
		}

		p.RemoveAssets([]string{asset})
		p.AddAssets([]string{assetName})
	}

	return c.postRepository.UpdatePost(ctx, cmd.Slug, func(p *post.Post) (*post.Post, error) {
		p.Publish(time.Now())
		return p, nil
	})
}
