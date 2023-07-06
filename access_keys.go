package buxclient

import (
	"context"

	buxmodels "github.com/BuxOrg/bux-models"
)

// GetAccessKey gets the access key given by id
func (b *BuxClient) GetAccessKey(ctx context.Context, id string) (*buxmodels.AccessKey, error) {
	// TODO: Add GetAccessKey to buxmodels
	// return b.transport.GetAccessKey(ctx, id)
	return nil, nil
}

// GetAccessKeys gets all the access keys filtered by the metadata
func (b *BuxClient) GetAccessKeys(ctx context.Context, metadataConditions *buxmodels.Metadata) ([]*buxmodels.AccessKey, error) {
	return b.transport.GetAccessKeys(ctx, metadataConditions)
}

// CreateAccessKey creates new access key
func (b *BuxClient) CreateAccessKey(ctx context.Context, metadata *buxmodels.Metadata) (*buxmodels.AccessKey, error) {
	return b.transport.CreateAccessKey(ctx, metadata)
}

// RevokeAccessKey revoked the access key given by id
func (b *BuxClient) RevokeAccessKey(ctx context.Context, id string) (*buxmodels.AccessKey, error) {
	return b.transport.RevokeAccessKey(ctx, id)
}
