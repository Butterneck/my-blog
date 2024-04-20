package asset

import "context"

type AssetRepository interface {
	PublishAsset(ctx context.Context, name string) (string, error)
	// MoveAsset(ctx context.Context, oldName, newName string) error
	DeleteAsset(ctx context.Context, name string) error
	GeneratePresignedURL(ctx context.Context, name string) (string, error)
}
