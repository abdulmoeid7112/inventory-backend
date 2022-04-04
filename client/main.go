package main

import (
	"context"
	"fmt"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/abdulmoeid7112/inventory-backend/client/rpcs"
	"github.com/abdulmoeid7112/inventory-backend/config"
	pb "github.com/abdulmoeid7112/inventory-backend/pb"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", viper.GetString(config.ServerHost), viper.GetString(config.ServerPort)), grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("Could not connect %v", err))
	}

	defer conn.Close()

	cli := rpcs.NewClient(pb.NewInventoryServiceClient(conn))

	requests(ctx, cli)
}

func requests(ctx context.Context, cli *rpcs.Client) {
	item, err := cli.AddInventoryItem(ctx, &pb.AddInventoryItemRequest{
		Item: &pb.Item{
			Id:          "12345",
			UserID:      "abc123",
			ItemName:    "Lorem",
			Description: "Lorem Ipsum lorem lorem ipsum",
			ImageUrl:    "https://example.com/test.jpg",
			Quantity:    5,
			Price:       1200,
		}})
	if err != nil {
		panic(err)
	}

	log().Printf("Item Added: %v\n", item)

	item, err = cli.GetInventoryItem(ctx, &pb.GetInventoryItemRequest{
		Id: "12345",
	})
	if err != nil {
		panic(err)
	}

	log().Printf("Item Retrieved: %v\n", item)

	item, err = cli.UpdateInventoryItem(ctx, &pb.UpdateInventoryItemRequest{
		Item: &pb.Item{
			Id:          "12345",
			UserID:      "abc123",
			ItemName:    "Lorem Updated",
			Description: "Lorem Ipsum lorem lorem ipsum updated",
			ImageUrl:    "https://example.com/test-updated.jpg",
			Quantity:    10,
			Price:       1500,
		}})
	if err != nil {
		panic(err)
	}

	log().Printf("Item Updated: %v\n", item)

	if _, err = cli.DeleteInventoryItem(ctx, &pb.DeleteInventoryItemRequest{Id: "12345"}); err != nil {
		panic(err)
	}

	log().Println("Item Deleted")

	itemList, err := cli.ListInventoryItems(ctx, &pb.ListInventoryItemsRequest{
		UserID: "abc123",
		Lim:    0,
		Off:    0,
	})
	if err != nil {
		panic(err)
	}

	log().Printf("Item List: %v\n", itemList)
}

func log() *logger.Entry {
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp: true,
	})

	return logger.WithFields(logger.Fields{
		"package": "client",
	})
}
