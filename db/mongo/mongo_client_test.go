package mongo

import (
	"os"
	"reflect"
	"testing"

	"github.com/abdulmoeid7112/inventory-backend/db"
	pb "github.com/abdulmoeid7112/inventory-backend/pb"
)

func Test_client_AddInventoryItem(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "inventory-backend-mongo-db")

	type args struct {
		item *pb.Item
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success - add item in db",
			args: args{item: &pb.Item{
				Id:          "1234",
				ItemName:    "Item 1",
				UserID:      "abc123",
				Description: "My Item Description",
				ImageUrl:    "https://example.com/img/test.jpg",
				Quantity:    4,
				Price:       1200.0,
			}},
			want:    "1234",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewMongoClient(db.Option{})
			got, err := m.AddInventoryItem(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddInventoryItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddInventoryItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_UpdateInventoryItem(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "inventory-backend-mongo-db")

	type args struct {
		id   string
		item *pb.Item
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - update item in db",
			args: args{
				id: "1234",
				item: &pb.Item{
					UserID:      "abc123",
					ItemName:    "Item Updated",
					Description: "My Item Description Updated",
					ImageUrl:    "https://example.com/img/test.jpg",
					Quantity:    4,
					Price:       1200.0,
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewMongoClient(db.Option{})
			if err := m.UpdateInventoryItem(tt.args.id, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("UpdateInventoryItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_DeleteInventoryItem(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "inventory-backend-mongo-db")

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - add item in db",
			args:    args{id: "1234"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewMongoClient(db.Option{})
			if err := m.DeleteInventoryItem(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteInventoryItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetInventoryItem(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "inventory-backend-mongo-db")

	m, _ := NewMongoClient(db.Option{})

	item1 := &pb.Item{
		Id:          "123",
		UserID:      "abc123",
		ItemName:    "Item 1",
		Description: "My Item Description",
		ImageUrl:    "https://example.com/img/test.jpg",
		Quantity:    4,
		Price:       1200.0,
	}
	_, _ = m.AddInventoryItem(item1)

	type args struct {
		id string
	}

	tests := []struct {
		name    string
		args    args
		want    *pb.Item
		wantErr bool
	}{
		{
			name:    "success - get item by id",
			args:    args{id: "123"},
			want:    item1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GetInventoryItem(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInventoryItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInventoryItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_ListInventoryItems(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "inventory-backend-mongo-db")

	m, _ := NewMongoClient(db.Option{})

	item1 := &pb.Item{
		Id:          "1234",
		UserID:      "abc123",
		ItemName:    "Item1",
		Description: "My Item Description",
		ImageUrl:    "https://example.com/img/test.jpg",
		Quantity:    4,
		Price:       1200.0,
	}
	_, _ = m.AddInventoryItem(item1)

	item2 := &pb.Item{
		Id:          "123456",
		UserID:      "abc123",
		ItemName:    "Item2",
		Description: "My Item 2 Description",
		ImageUrl:    "https://example.com/img/test.jpg",
		Quantity:    4,
		Price:       1200.0,
	}
	_, _ = m.AddInventoryItem(item2)

	type args struct {
		filter map[string]interface{}
		lim    int64
		off    int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*pb.Item
		wantErr bool
	}{
		{
			name: "success - list all items",
			args: args{
				filter: map[string]interface{}{
					"itemname": "Item1",
				},
				lim: 1,
				off: 0,
			},
			want:    []*pb.Item{item1},
			wantErr: false,
		},
		{
			name: "success - list all items",
			args: args{
				filter: map[string]interface{}{
					"itemname": "Item2",
				},
				lim: 1,
				off: 0,
			},
			want:    []*pb.Item{item2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.ListInventoryItems(tt.args.filter, tt.args.lim, tt.args.off)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListInventoryItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListInventoryItems() got = %v, want %v", got, tt.want)
			}
		})
	}
}
