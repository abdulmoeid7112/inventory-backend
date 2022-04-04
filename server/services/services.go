package services

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/abdulmoeid7112/inventory-backend/pb"
)

// AddInventoryItem - add inventory item.
func (s *Service) AddInventoryItem(ctx context.Context, req *pb.AddInventoryItemRequest) (*pb.Item, error) {
	log().Infof("server received client add request: %+v", req)

	if _, err := s.db.AddInventoryItem(req.Item); err != nil {
		return nil, errors.Wrap(err, "failed to add")
	}

	return req.Item, nil
}

// GetInventoryItem - get inventory item.
func (s *Service) GetInventoryItem(ctx context.Context, req *pb.GetInventoryItemRequest) (*pb.Item, error) {
	log().Infof("server received client get request: %+v", req)

	item, err := s.db.GetInventoryItem(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}

	return item, nil
}

// UpdateInventoryItem - update inventory item details.
func (s *Service) UpdateInventoryItem(ctx context.Context, req *pb.UpdateInventoryItemRequest) (*pb.Item, error) {
	log().Infof("server received client update request: %+v", req)

	if err := s.db.UpdateInventoryItem(req.Item.Id, req.Item); err != nil {
		return nil, errors.Wrap(err, "failed to update")
	}

	return req.Item, nil
}

// DeleteInventoryItem - delete inventory item.
func (s *Service) DeleteInventoryItem(ctx context.Context, req *pb.DeleteInventoryItemRequest) (*emptypb.Empty, error) {
	log().Infof("server received client delete request: %+v", req)

	if err := s.db.DeleteInventoryItem(req.Id); err != nil {
		return &emptypb.Empty{}, errors.Wrap(err, "failed to delete")
	}

	return &emptypb.Empty{}, nil
}

// ListInventoryItems - list all inventory items.
func (s *Service) ListInventoryItems(ctx context.Context, req *pb.ListInventoryItemsRequest) (*pb.ListInventoryItemsResponse, error) {
	log().Infof("server received client list request: %+v", req)

	itemList, err := s.db.ListInventoryItems(map[string]interface{}{"userid": req.UserID}, req.Lim, req.Off)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list")
	}

	return &pb.ListInventoryItemsResponse{Items: itemList}, nil
}
