// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.7
// source: card/card.proto

package cardv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Card_BalanceByPan_FullMethodName = "/card.Card/BalanceByPan"
)

// CardClient is the client API for Card service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardClient interface {
	BalanceByPan(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
}

type cardClient struct {
	cc grpc.ClientConnInterface
}

func NewCardClient(cc grpc.ClientConnInterface) CardClient {
	return &cardClient{cc}
}

func (c *cardClient) BalanceByPan(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, Card_BalanceByPan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardServer is the server API for Card service.
// All implementations must embed UnimplementedCardServer
// for forward compatibility.
type CardServer interface {
	BalanceByPan(context.Context, *BalanceRequest) (*BalanceResponse, error)
	mustEmbedUnimplementedCardServer()
}

// UnimplementedCardServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCardServer struct{}

func (UnimplementedCardServer) BalanceByPan(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalanceByPan not implemented")
}
func (UnimplementedCardServer) mustEmbedUnimplementedCardServer() {}
func (UnimplementedCardServer) testEmbeddedByValue()              {}

// UnsafeCardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardServer will
// result in compilation errors.
type UnsafeCardServer interface {
	mustEmbedUnimplementedCardServer()
}

func RegisterCardServer(s grpc.ServiceRegistrar, srv CardServer) {
	// If the following call pancis, it indicates UnimplementedCardServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Card_ServiceDesc, srv)
}

func _Card_BalanceByPan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServer).BalanceByPan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Card_BalanceByPan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServer).BalanceByPan(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Card_ServiceDesc is the grpc.ServiceDesc for Card service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Card_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "card.Card",
	HandlerType: (*CardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BalanceByPan",
			Handler:    _Card_BalanceByPan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "card/card.proto",
}
