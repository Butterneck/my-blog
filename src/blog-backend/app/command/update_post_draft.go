package command

import (
	"context"
	"fmt"

	"github.com/butterneck/my-blog/src/blog-backend/domain/post"
)

type UpdatePostDraft struct {
	Slug          string
	Title         *string
	Body          *string
	NewAssets     []PostAsset
	DeletedAssets []string
}

type UpdatePostDraftHandler struct {
	postRepository post.Repository
	assetStore     post.AssetStore
}

func NewUpdatePostDraftHandler(postRepository post.Repository, assetStore post.AssetStore) UpdatePostDraftHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	if assetStore == nil {
		panic("assetStore is nil")
	}

	return UpdatePostDraftHandler{postRepository: postRepository, assetStore: assetStore}
}

func (c UpdatePostDraftHandler) Handle(ctx context.Context, cmd UpdatePostDraft) error {

	newPostAssets := []string{}
	for _, asset := range cmd.NewAssets {
		_, err := c.assetStore.UploadAsset(ctx, asset.File, asset.Name)
		if err != nil {
			return err
		}

		newPostAssets = append(newPostAssets, asset.Name)
	}

	// TODO: Replace on the body all the occurrencies of the old filename with the new one (the path on s3)

	deletedPostAssets := []string{}
	for _, asset := range cmd.DeletedAssets {
		err := c.assetStore.DeleteAsset(ctx, asset)
		if err != nil {
			fmt.Println("Error deleting asset", err)
		}

		deletedPostAssets = append(deletedPostAssets, asset)
	}

	return c.postRepository.UpdatePost(ctx, cmd.Slug, func(p *post.Post) (*post.Post, error) {
		if cmd.Title != nil {
			p.UpdateTitle(*cmd.Title)
		}

		if cmd.Body != nil {
			p.UpdateBody(*cmd.Body)
		}

		if len(newPostAssets) > 0 {
			p.AddAssets(newPostAssets)
		}

		if len(deletedPostAssets) > 0 {
			p.RemoveAssets(deletedPostAssets)
		}

		return p, nil
	})
}
