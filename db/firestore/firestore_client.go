package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/api/option"

	"github.com/abdulmoeid7112/inventory-backend/config"
	"github.com/abdulmoeid7112/inventory-backend/db"
	pb "github.com/abdulmoeid7112/inventory-backend/pb"
)

const (
	collectionName = "item"
)

func init() {
	db.Register("firestore", NewFirestoreClient)
}

type firestoreClient struct {
	client *firestore.Client
}

// NewFirestoreClient initializes a firestore database connection.
func NewFirestoreClient(conf db.Option) (db.DataStore, error) {
	log().Infof("initializing firestore: %s", viper.GetString(config.GCPProjectID))

	var options []option.ClientOption

	options = append(options, option.WithCredentialsFile(viper.GetString(config.GCPCredFile)))

	client, err := firestore.NewClient(context.Background(), viper.GetString(config.GCPProjectID), options...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}

	return &firestoreClient{client: client}, nil
}

func (f firestoreClient) AddInventoryItem(item *pb.Item) (string, error) {
	_, err := f.client.Collection(collectionName).Doc(item.GetId()).Set(context.TODO(), item)
	if err != nil {
		return "", errors.Wrap(err, "failed to make add item request")
	}

	return item.GetId(), nil
}

func (f firestoreClient) UpdateInventoryItem(id string, item *pb.Item) error {
	_, err := f.client.Collection(collectionName).Doc(id).Set(context.TODO(), item)
	if err != nil {
		return errors.Wrap(err, "failed to make update item request")
	}

	return nil
}

func (f firestoreClient) GetInventoryItem(id string) (*pb.Item, error) {
	var item *pb.Item

	doc, err := f.client.Collection(collectionName).Doc(id).Get(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "failed to make get item request")
	}

	if err = doc.DataTo(&item); err != nil {
		return nil, errors.Wrap(err, "failed to populate document to struct")
	}

	return item, nil
}

func (f firestoreClient) ListInventoryItems(filter map[string]interface{}, lim int64, off int64) ([]*pb.Item, error) {
	items := make([]*pb.Item, 0)

	documents, err := f.client.Collection(collectionName).Where("UserID", "==", filter["userid"]).Documents(context.TODO()).GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to make list items request")
	}

	for _, doc := range documents {
		var item *pb.Item

		if err = doc.DataTo(&item); err != nil {
			return nil, errors.Wrap(err, "failed to populate document to struct")
		}

		items = append(items, item)
	}

	return items, nil
}

func (f firestoreClient) DeleteInventoryItem(id string) error {
	_, err := f.client.Collection(collectionName).Doc(id).Delete(context.TODO())
	if err != nil {
		return errors.Wrap(err, "failed to make delete item request")
	}

	return nil
}
