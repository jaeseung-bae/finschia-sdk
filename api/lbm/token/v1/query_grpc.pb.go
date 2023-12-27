// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: lbm/token/v1/query.proto

package tokenv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_Balance_FullMethodName           = "/lbm.token.v1.Query/Balance"
	Query_Supply_FullMethodName            = "/lbm.token.v1.Query/Supply"
	Query_Minted_FullMethodName            = "/lbm.token.v1.Query/Minted"
	Query_Burnt_FullMethodName             = "/lbm.token.v1.Query/Burnt"
	Query_Contract_FullMethodName          = "/lbm.token.v1.Query/Contract"
	Query_GranteeGrants_FullMethodName     = "/lbm.token.v1.Query/GranteeGrants"
	Query_IsOperatorFor_FullMethodName     = "/lbm.token.v1.Query/IsOperatorFor"
	Query_HoldersByOperator_FullMethodName = "/lbm.token.v1.Query/HoldersByOperator"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Balance queries the number of tokens of a given contract owned by the address.
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
	// Supply queries the number of tokens from the given contract id.
	Supply(ctx context.Context, in *QuerySupplyRequest, opts ...grpc.CallOption) (*QuerySupplyResponse, error)
	// Minted queries the number of minted tokens from the given contract id.
	Minted(ctx context.Context, in *QueryMintedRequest, opts ...grpc.CallOption) (*QueryMintedResponse, error)
	// Burnt queries the number of burnt tokens from the given contract id.
	Burnt(ctx context.Context, in *QueryBurntRequest, opts ...grpc.CallOption) (*QueryBurntResponse, error)
	// Contract queries an token metadata based on its contract id.
	Contract(ctx context.Context, in *QueryContractRequest, opts ...grpc.CallOption) (*QueryContractResponse, error)
	// GranteeGrants queries permissions on a given grantee.
	GranteeGrants(ctx context.Context, in *QueryGranteeGrantsRequest, opts ...grpc.CallOption) (*QueryGranteeGrantsResponse, error)
	// IsOperatorFor queries authorization on a given operator holder pair.
	IsOperatorFor(ctx context.Context, in *QueryIsOperatorForRequest, opts ...grpc.CallOption) (*QueryIsOperatorForResponse, error)
	// HoldersByOperator queries holders on a given operator.
	HoldersByOperator(ctx context.Context, in *QueryHoldersByOperatorRequest, opts ...grpc.CallOption) (*QueryHoldersByOperatorResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, Query_Balance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Supply(ctx context.Context, in *QuerySupplyRequest, opts ...grpc.CallOption) (*QuerySupplyResponse, error) {
	out := new(QuerySupplyResponse)
	err := c.cc.Invoke(ctx, Query_Supply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Minted(ctx context.Context, in *QueryMintedRequest, opts ...grpc.CallOption) (*QueryMintedResponse, error) {
	out := new(QueryMintedResponse)
	err := c.cc.Invoke(ctx, Query_Minted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Burnt(ctx context.Context, in *QueryBurntRequest, opts ...grpc.CallOption) (*QueryBurntResponse, error) {
	out := new(QueryBurntResponse)
	err := c.cc.Invoke(ctx, Query_Burnt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Contract(ctx context.Context, in *QueryContractRequest, opts ...grpc.CallOption) (*QueryContractResponse, error) {
	out := new(QueryContractResponse)
	err := c.cc.Invoke(ctx, Query_Contract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GranteeGrants(ctx context.Context, in *QueryGranteeGrantsRequest, opts ...grpc.CallOption) (*QueryGranteeGrantsResponse, error) {
	out := new(QueryGranteeGrantsResponse)
	err := c.cc.Invoke(ctx, Query_GranteeGrants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IsOperatorFor(ctx context.Context, in *QueryIsOperatorForRequest, opts ...grpc.CallOption) (*QueryIsOperatorForResponse, error) {
	out := new(QueryIsOperatorForResponse)
	err := c.cc.Invoke(ctx, Query_IsOperatorFor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HoldersByOperator(ctx context.Context, in *QueryHoldersByOperatorRequest, opts ...grpc.CallOption) (*QueryHoldersByOperatorResponse, error) {
	out := new(QueryHoldersByOperatorResponse)
	err := c.cc.Invoke(ctx, Query_HoldersByOperator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Balance queries the number of tokens of a given contract owned by the address.
	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
	// Supply queries the number of tokens from the given contract id.
	Supply(context.Context, *QuerySupplyRequest) (*QuerySupplyResponse, error)
	// Minted queries the number of minted tokens from the given contract id.
	Minted(context.Context, *QueryMintedRequest) (*QueryMintedResponse, error)
	// Burnt queries the number of burnt tokens from the given contract id.
	Burnt(context.Context, *QueryBurntRequest) (*QueryBurntResponse, error)
	// Contract queries an token metadata based on its contract id.
	Contract(context.Context, *QueryContractRequest) (*QueryContractResponse, error)
	// GranteeGrants queries permissions on a given grantee.
	GranteeGrants(context.Context, *QueryGranteeGrantsRequest) (*QueryGranteeGrantsResponse, error)
	// IsOperatorFor queries authorization on a given operator holder pair.
	IsOperatorFor(context.Context, *QueryIsOperatorForRequest) (*QueryIsOperatorForResponse, error)
	// HoldersByOperator queries holders on a given operator.
	HoldersByOperator(context.Context, *QueryHoldersByOperatorRequest) (*QueryHoldersByOperatorResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedQueryServer) Supply(context.Context, *QuerySupplyRequest) (*QuerySupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Supply not implemented")
}
func (UnimplementedQueryServer) Minted(context.Context, *QueryMintedRequest) (*QueryMintedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Minted not implemented")
}
func (UnimplementedQueryServer) Burnt(context.Context, *QueryBurntRequest) (*QueryBurntResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Burnt not implemented")
}
func (UnimplementedQueryServer) Contract(context.Context, *QueryContractRequest) (*QueryContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Contract not implemented")
}
func (UnimplementedQueryServer) GranteeGrants(context.Context, *QueryGranteeGrantsRequest) (*QueryGranteeGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GranteeGrants not implemented")
}
func (UnimplementedQueryServer) IsOperatorFor(context.Context, *QueryIsOperatorForRequest) (*QueryIsOperatorForResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsOperatorFor not implemented")
}
func (UnimplementedQueryServer) HoldersByOperator(context.Context, *QueryHoldersByOperatorRequest) (*QueryHoldersByOperatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HoldersByOperator not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Balance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Supply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Supply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Supply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Supply(ctx, req.(*QuerySupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Minted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMintedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Minted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Minted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Minted(ctx, req.(*QueryMintedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Burnt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBurntRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Burnt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Burnt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Burnt(ctx, req.(*QueryBurntRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Contract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Contract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Contract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Contract(ctx, req.(*QueryContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GranteeGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGranteeGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GranteeGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GranteeGrants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GranteeGrants(ctx, req.(*QueryGranteeGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IsOperatorFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIsOperatorForRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IsOperatorFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_IsOperatorFor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IsOperatorFor(ctx, req.(*QueryIsOperatorForRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HoldersByOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHoldersByOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HoldersByOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_HoldersByOperator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HoldersByOperator(ctx, req.(*QueryHoldersByOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lbm.token.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
		{
			MethodName: "Supply",
			Handler:    _Query_Supply_Handler,
		},
		{
			MethodName: "Minted",
			Handler:    _Query_Minted_Handler,
		},
		{
			MethodName: "Burnt",
			Handler:    _Query_Burnt_Handler,
		},
		{
			MethodName: "Contract",
			Handler:    _Query_Contract_Handler,
		},
		{
			MethodName: "GranteeGrants",
			Handler:    _Query_GranteeGrants_Handler,
		},
		{
			MethodName: "IsOperatorFor",
			Handler:    _Query_IsOperatorFor_Handler,
		},
		{
			MethodName: "HoldersByOperator",
			Handler:    _Query_HoldersByOperator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lbm/token/v1/query.proto",
}