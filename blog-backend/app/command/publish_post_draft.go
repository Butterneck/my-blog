package command

import (
	"context"

	"github.com/butterneck/my-blog/blog-backend/domain/post"
)

type PublishPostDraft struct {
	Slug string
}

type PublishPostDraftHandler struct {
	postRepository post.Repository
	postAssetStore post.AssetStore
}

func NewPublishPostDraftHandler(postRepository post.Repository, assetStore post.AssetStore) PublishPostDraftHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	if assetStore == nil {
		panic("assetStore is nil")
	}

	return PublishPostDraftHandler{postRepository: postRepository, postAssetStore: assetStore}
}

func (c PublishPostDraftHandler) Handle(ctx context.Context, cmd PublishPostDraft) error {

	p, err := c.postRepository.GetAnyPost(ctx, cmd.Slug)
	if err != nil {
		return err
	}

	// Move draft assets to published assets
	for _, asset := range p.Draft().Assets() {
		err := c.postAssetStore.MoveAsset(ctx, p.Slug()+"/draft/"+asset, p.Slug()+"/published/"+asset)
		if err != nil {
			return err
		}
	}

	return c.postRepository.UpdatePost(ctx, cmd.Slug, func(p *post.Post) (*post.Post, error) {
		p.Publish()
		return p, nil
	})
}
