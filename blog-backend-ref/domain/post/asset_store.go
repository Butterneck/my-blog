package post

import "context"

type AssetStore interface {
	UploadAsset(ctx context.Context, asset []byte, name string) error
	MoveAsset(ctx context.Context, oldName, newName string) error
	DeleteAsset(ctx context.Context, name string) error
}
