// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: pb/inventory.proto

package inventory_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryServiceClient interface {
	AddInventoryItem(ctx context.Context, in *AddInventoryItemRequest, opts ...grpc.CallOption) (*Item, error)
	GetInventoryItem(ctx context.Context, in *GetInventoryItemRequest, opts ...grpc.CallOption) (*Item, error)
	DeleteInventoryItem(ctx context.Context, in *DeleteInventoryItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateInventoryItem(ctx context.Context, in *UpdateInventoryItemRequest, opts ...grpc.CallOption) (*Item, error)
	ListInventoryItems(ctx context.Context, in *ListInventoryItemsRequest, opts ...grpc.CallOption) (*ListInventoryItemsResponse, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) AddInventoryItem(ctx context.Context, in *AddInventoryItemRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/inventoryPB.InventoryService/AddInventoryItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetInventoryItem(ctx context.Context, in *GetInventoryItemRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/inventoryPB.InventoryService/GetInventoryItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) DeleteInventoryItem(ctx context.Context, in *DeleteInventoryItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/inventoryPB.InventoryService/DeleteInventoryItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) UpdateInventoryItem(ctx context.Context, in *UpdateInventoryItemRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/inventoryPB.InventoryService/UpdateInventoryItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) ListInventoryItems(ctx context.Context, in *ListInventoryItemsRequest, opts ...grpc.CallOption) (*ListInventoryItemsResponse, error) {
	out := new(ListInventoryItemsResponse)
	err := c.cc.Invoke(ctx, "/inventoryPB.InventoryService/ListInventoryItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations should embed UnimplementedInventoryServiceServer
// for forward compatibility
type InventoryServiceServer interface {
	AddInventoryItem(context.Context, *AddInventoryItemRequest) (*Item, error)
	GetInventoryItem(context.Context, *GetInventoryItemRequest) (*Item, error)
	DeleteInventoryItem(context.Context, *DeleteInventoryItemRequest) (*emptypb.Empty, error)
	UpdateInventoryItem(context.Context, *UpdateInventoryItemRequest) (*Item, error)
	ListInventoryItems(context.Context, *ListInventoryItemsRequest) (*ListInventoryItemsResponse, error)
}

// UnimplementedInventoryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInventoryServiceServer struct {
}

func (UnimplementedInventoryServiceServer) AddInventoryItem(context.Context, *AddInventoryItemRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInventoryItem not implemented")
}
func (UnimplementedInventoryServiceServer) GetInventoryItem(context.Context, *GetInventoryItemRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInventoryItem not implemented")
}
func (UnimplementedInventoryServiceServer) DeleteInventoryItem(context.Context, *DeleteInventoryItemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInventoryItem not implemented")
}
func (UnimplementedInventoryServiceServer) UpdateInventoryItem(context.Context, *UpdateInventoryItemRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInventoryItem not implemented")
}
func (UnimplementedInventoryServiceServer) ListInventoryItems(context.Context, *ListInventoryItemsRequest) (*ListInventoryItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInventoryItems not implemented")
}

// UnsafeInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServiceServer will
// result in compilation errors.
type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_AddInventoryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInventoryItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).AddInventoryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventoryPB.InventoryService/AddInventoryItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).AddInventoryItem(ctx, req.(*AddInventoryItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetInventoryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInventoryItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetInventoryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventoryPB.InventoryService/GetInventoryItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetInventoryItem(ctx, req.(*GetInventoryItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_DeleteInventoryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInventoryItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).DeleteInventoryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventoryPB.InventoryService/DeleteInventoryItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).DeleteInventoryItem(ctx, req.(*DeleteInventoryItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_UpdateInventoryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInventoryItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).UpdateInventoryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventoryPB.InventoryService/UpdateInventoryItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).UpdateInventoryItem(ctx, req.(*UpdateInventoryItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_ListInventoryItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInventoryItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).ListInventoryItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventoryPB.InventoryService/ListInventoryItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).ListInventoryItems(ctx, req.(*ListInventoryItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventoryPB.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddInventoryItem",
			Handler:    _InventoryService_AddInventoryItem_Handler,
		},
		{
			MethodName: "GetInventoryItem",
			Handler:    _InventoryService_GetInventoryItem_Handler,
		},
		{
			MethodName: "DeleteInventoryItem",
			Handler:    _InventoryService_DeleteInventoryItem_Handler,
		},
		{
			MethodName: "UpdateInventoryItem",
			Handler:    _InventoryService_UpdateInventoryItem_Handler,
		},
		{
			MethodName: "ListInventoryItems",
			Handler:    _InventoryService_ListInventoryItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/inventory.proto",
}
