// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: generate.proto

package intrv1

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
	RPCGenerator_GenerateURL_FullMethodName      = "/intr.v1.RPCGenerator/GenerateURL"
	RPCGenerator_BatchGenerateURL_FullMethodName = "/intr.v1.RPCGenerator/BatchGenerateURL"
)

// RPCGeneratorClient is the client API for RPCGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCGeneratorClient interface {
	// 生成单条短链
	GenerateURL(ctx context.Context, in *URLRequest, opts ...grpc.CallOption) (*URLResponse, error)
	// 批量生成短链
	BatchGenerateURL(ctx context.Context, in *BatchURLRequest, opts ...grpc.CallOption) (*BatchURLResponse, error)
}

type rPCGeneratorClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCGeneratorClient(cc grpc.ClientConnInterface) RPCGeneratorClient {
	return &rPCGeneratorClient{cc}
}

func (c *rPCGeneratorClient) GenerateURL(ctx context.Context, in *URLRequest, opts ...grpc.CallOption) (*URLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(URLResponse)
	err := c.cc.Invoke(ctx, RPCGenerator_GenerateURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCGeneratorClient) BatchGenerateURL(ctx context.Context, in *BatchURLRequest, opts ...grpc.CallOption) (*BatchURLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchURLResponse)
	err := c.cc.Invoke(ctx, RPCGenerator_BatchGenerateURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCGeneratorServer is the server API for RPCGenerator service.
// All implementations must embed UnimplementedRPCGeneratorServer
// for forward compatibility.
type RPCGeneratorServer interface {
	// 生成单条短链
	GenerateURL(context.Context, *URLRequest) (*URLResponse, error)
	// 批量生成短链
	BatchGenerateURL(context.Context, *BatchURLRequest) (*BatchURLResponse, error)
	mustEmbedUnimplementedRPCGeneratorServer()
}

// UnimplementedRPCGeneratorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRPCGeneratorServer struct{}

func (UnimplementedRPCGeneratorServer) GenerateURL(context.Context, *URLRequest) (*URLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateURL not implemented")
}
func (UnimplementedRPCGeneratorServer) BatchGenerateURL(context.Context, *BatchURLRequest) (*BatchURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGenerateURL not implemented")
}
func (UnimplementedRPCGeneratorServer) mustEmbedUnimplementedRPCGeneratorServer() {}
func (UnimplementedRPCGeneratorServer) testEmbeddedByValue()                      {}

// UnsafeRPCGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCGeneratorServer will
// result in compilation errors.
type UnsafeRPCGeneratorServer interface {
	mustEmbedUnimplementedRPCGeneratorServer()
}

func RegisterRPCGeneratorServer(s grpc.ServiceRegistrar, srv RPCGeneratorServer) {
	// If the following call pancis, it indicates UnimplementedRPCGeneratorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RPCGenerator_ServiceDesc, srv)
}

func _RPCGenerator_GenerateURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCGeneratorServer).GenerateURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCGenerator_GenerateURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCGeneratorServer).GenerateURL(ctx, req.(*URLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCGenerator_BatchGenerateURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCGeneratorServer).BatchGenerateURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCGenerator_BatchGenerateURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCGeneratorServer).BatchGenerateURL(ctx, req.(*BatchURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCGenerator_ServiceDesc is the grpc.ServiceDesc for RPCGenerator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCGenerator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "intr.v1.RPCGenerator",
	HandlerType: (*RPCGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateURL",
			Handler:    _RPCGenerator_GenerateURL_Handler,
		},
		{
			MethodName: "BatchGenerateURL",
			Handler:    _RPCGenerator_BatchGenerateURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generate.proto",
}
