package rpcs

import (
	pb "github.com/abdulmoeid7112/inventory-backend/pb"
)

// Client - client struct.
type Client struct {
	client pb.InventoryServiceClient
}

// NewClient creates a connection to our client.
func NewClient(cli pb.InventoryServiceClient) *Client {
	return &Client{client: cli}
}
