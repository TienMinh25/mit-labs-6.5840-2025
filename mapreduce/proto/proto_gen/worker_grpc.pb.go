// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: worker.proto

package proto_gen

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
	Worker_AssignMapTask_FullMethodName    = "/Worker/AssignMapTask"
	Worker_AssignReduceTask_FullMethodName = "/Worker/AssignReduceTask"
	Worker_End_FullMethodName              = "/Worker/End"
	Worker_Health_FullMethodName           = "/Worker/Health"
	Worker_GetIMDFile_FullMethodName       = "/Worker/GetIMDFile"
)

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// service server for worker server
type WorkerClient interface {
	AssignMapTask(ctx context.Context, in *AssignMapTaskReq, opts ...grpc.CallOption) (*Result, error)
	AssignReduceTask(ctx context.Context, in *AssignReduceTaskReq, opts ...grpc.CallOption) (*Result, error)
	End(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Health(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthRes, error)
	GetIMDFile(ctx context.Context, in *GetIMDFileReq, opts ...grpc.CallOption) (*GetIMDFileRes, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) AssignMapTask(ctx context.Context, in *AssignMapTaskReq, opts ...grpc.CallOption) (*Result, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Result)
	err := c.cc.Invoke(ctx, Worker_AssignMapTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) AssignReduceTask(ctx context.Context, in *AssignReduceTaskReq, opts ...grpc.CallOption) (*Result, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Result)
	err := c.cc.Invoke(ctx, Worker_AssignReduceTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) End(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Worker_End_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Health(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthRes)
	err := c.cc.Invoke(ctx, Worker_Health_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetIMDFile(ctx context.Context, in *GetIMDFileReq, opts ...grpc.CallOption) (*GetIMDFileRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIMDFileRes)
	err := c.cc.Invoke(ctx, Worker_GetIMDFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
// All implementations must embed UnimplementedWorkerServer
// for forward compatibility.
//
// service server for worker server
type WorkerServer interface {
	AssignMapTask(context.Context, *AssignMapTaskReq) (*Result, error)
	AssignReduceTask(context.Context, *AssignReduceTaskReq) (*Result, error)
	End(context.Context, *Empty) (*Empty, error)
	Health(context.Context, *Empty) (*HealthRes, error)
	GetIMDFile(context.Context, *GetIMDFileReq) (*GetIMDFileRes, error)
	mustEmbedUnimplementedWorkerServer()
}

// UnimplementedWorkerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWorkerServer struct{}

func (UnimplementedWorkerServer) AssignMapTask(context.Context, *AssignMapTaskReq) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignMapTask not implemented")
}
func (UnimplementedWorkerServer) AssignReduceTask(context.Context, *AssignReduceTaskReq) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignReduceTask not implemented")
}
func (UnimplementedWorkerServer) End(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method End not implemented")
}
func (UnimplementedWorkerServer) Health(context.Context, *Empty) (*HealthRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedWorkerServer) GetIMDFile(context.Context, *GetIMDFileReq) (*GetIMDFileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIMDFile not implemented")
}
func (UnimplementedWorkerServer) mustEmbedUnimplementedWorkerServer() {}
func (UnimplementedWorkerServer) testEmbeddedByValue()                {}

// UnsafeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServer will
// result in compilation errors.
type UnsafeWorkerServer interface {
	mustEmbedUnimplementedWorkerServer()
}

func RegisterWorkerServer(s grpc.ServiceRegistrar, srv WorkerServer) {
	// If the following call pancis, it indicates UnimplementedWorkerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Worker_ServiceDesc, srv)
}

func _Worker_AssignMapTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignMapTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).AssignMapTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worker_AssignMapTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).AssignMapTask(ctx, req.(*AssignMapTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_AssignReduceTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignReduceTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).AssignReduceTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worker_AssignReduceTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).AssignReduceTask(ctx, req.(*AssignReduceTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_End_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).End(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worker_End_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).End(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worker_Health_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Health(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetIMDFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIMDFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetIMDFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worker_GetIMDFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetIMDFile(ctx, req.(*GetIMDFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Worker_ServiceDesc is the grpc.ServiceDesc for Worker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Worker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignMapTask",
			Handler:    _Worker_AssignMapTask_Handler,
		},
		{
			MethodName: "AssignReduceTask",
			Handler:    _Worker_AssignReduceTask_Handler,
		},
		{
			MethodName: "End",
			Handler:    _Worker_End_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Worker_Health_Handler,
		},
		{
			MethodName: "GetIMDFile",
			Handler:    _Worker_GetIMDFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "worker.proto",
}
