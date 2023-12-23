package post

import "context"

type AssetStore interface {
	UploadAsset(ctx context.Context, asset []byte, name string) (string, error)
	DeleteAsset(ctx context.Context, name string) error
}
