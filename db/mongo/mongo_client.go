package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/abdulmoeid7112/inventory-backend/config"
	"github.com/abdulmoeid7112/inventory-backend/db"
	pb "github.com/abdulmoeid7112/inventory-backend/pb"
)

const (
	collectionName = "item"
)

func init() {
	db.Register("mongo", NewMongoClient)
}

type client struct {
	conn *mongo.Client
}

// NewMongoClient initializes a mysql database connection.
func NewMongoClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)

	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}

	return &client{conn: cli}, nil
}

func (m client) AddInventoryItem(item *pb.Item) (string, error) {
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if _, err := collection.InsertOne(context.TODO(), item); err != nil {
		return "", errors.Wrap(err, "failed to add item in db")
	}

	return item.Id, nil
}

func (m client) UpdateInventoryItem(id string, item *pb.Item) error {
	item.Id = id

	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"id": bson.M{"$eq": item.Id}}, bson.M{"$set": item}); err != nil {
		return errors.Wrap(err, "failed to update item in db")
	}

	return nil
}

func (m client) GetInventoryItem(id string) (*pb.Item, error) {
	var item *pb.Item

	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&item); err != nil {
		if mongo.ErrNoDocuments != nil {
			return nil, errors.Wrap(err, "failed to fetch item....not found")
		}

		return nil, errors.Wrap(err, "failed to fetch item from db")
	}

	return item, nil
}

func (m client) ListInventoryItems(filter map[string]interface{}, lim int64, off int64) ([]*pb.Item, error) {
	var items []*pb.Item
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)

	cursor, err := collection.Find(context.TODO(), filter, &options.FindOptions{
		Skip:  &off,
		Limit: &lim,
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve items from db")
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var em *pb.Item
		if err = cursor.Decode(&em); err != nil {
			return nil, errors.Wrap(err, "failed to scan retrieved rows from db")
		}
		items = append(items, em)
	}

	return items, nil
}

func (m client) DeleteInventoryItem(id string) error {
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"id": id}); err != nil {
		return errors.Wrap(err, "failed to delete items from db")
	}

	return nil
}
