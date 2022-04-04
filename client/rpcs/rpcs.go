package rpcs

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/abdulmoeid7112/inventory-backend/pb"
)

// AddInventoryItem - add inventory item.
func (c *Client) AddInventoryItem(ctx context.Context, req *pb.AddInventoryItemRequest) (*pb.Item, error) {
	log().Infof("client initiated add request: %+v", req)

	res, err := c.client.AddInventoryItem(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make add item request")
	}

	return res, nil
}

// GetInventoryItem - get inventory item.
func (c *Client) GetInventoryItem(ctx context.Context, req *pb.GetInventoryItemRequest) (*pb.Item, error) {
	log().Infof("client initiated get request: %+v", req)

	res, err := c.client.GetInventoryItem(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make get item request")
	}

	return res, nil
}

// DeleteInventoryItem - delete inventory item.
func (c *Client) DeleteInventoryItem(ctx context.Context, req *pb.DeleteInventoryItemRequest) (*emptypb.Empty, error) {
	log().Infof("client initiated delete request: %+v", req)

	res, err := c.client.DeleteInventoryItem(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make delete item request")
	}

	return res, nil
}

// UpdateInventoryItem - update inventory item details.
func (c *Client) UpdateInventoryItem(ctx context.Context, req *pb.UpdateInventoryItemRequest) (*pb.Item, error) {
	log().Infof("client initiated update request: %+v", req)

	res, err := c.client.UpdateInventoryItem(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make update item request")
	}

	return res, nil
}

// ListInventoryItems - list all inventory items.
func (c *Client) ListInventoryItems(ctx context.Context, req *pb.ListInventoryItemsRequest) (*pb.ListInventoryItemsResponse, error) {
	log().Infof("client initiated list request: %+v", req)

	res, err := c.client.ListInventoryItems(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make list item request")
	}

	return res, nil
}
