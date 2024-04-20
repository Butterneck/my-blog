package command

import (
	"context"

	"github.com/butterneck/my-blog/blog-backend-back/domain/post"
)

type UnpublishPost struct {
	Slug string
}

type UnpublishPostHandler struct {
	postRepository post.Repository
	postAssetStore post.AssetStore
}

func NewUnpublishPostHandler(postRepository post.Repository, assetStore post.AssetStore) UnpublishPostHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	if assetStore == nil {
		panic("assetStore is nil")
	}

	return UnpublishPostHandler{postRepository: postRepository, postAssetStore: assetStore}
}

func (c UnpublishPostHandler) Handle(ctx context.Context, cmd UnpublishPost) error {

	p, err := c.postRepository.GetAnyPost(ctx, cmd.Slug)
	if err != nil {
		return err
	}

	// Move draft assets to published assets
	for _, asset := range p.Draft().Assets() {
		err := c.postAssetStore.MoveAsset(ctx, p.Slug()+"/published/"+asset, p.Slug()+"/draft/"+asset)
		if err != nil {
			return err
		}
	}

	return c.postRepository.UpdatePost(ctx, cmd.Slug, func(p *post.Post) (*post.Post, error) {
		p.Unpublish()
		return p, nil
	})
}
