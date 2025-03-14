// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: manager.proto

package pb

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
	ManagerService_UpsertManifest_FullMethodName  = "/claudie.ManagerService/UpsertManifest"
	ManagerService_MarkForDeletion_FullMethodName = "/claudie.ManagerService/MarkForDeletion"
	ManagerService_ListConfigs_FullMethodName     = "/claudie.ManagerService/ListConfigs"
	ManagerService_GetConfig_FullMethodName       = "/claudie.ManagerService/GetConfig"
	ManagerService_NextTask_FullMethodName        = "/claudie.ManagerService/NextTask"
	ManagerService_TaskUpdate_FullMethodName      = "/claudie.ManagerService/TaskUpdate"
	ManagerService_TaskComplete_FullMethodName    = "/claudie.ManagerService/TaskComplete"
	ManagerService_UpdateNodePool_FullMethodName  = "/claudie.ManagerService/UpdateNodePool"
)

// ManagerServiceClient is the client API for ManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerServiceClient interface {
	// UpsertManifest will process the request by either created a new configuration for the
	// given input manifest or updating an existing one.
	UpsertManifest(ctx context.Context, in *UpsertManifestRequest, opts ...grpc.CallOption) (*UpsertManifestResponse, error)
	// MarkForDeletion will mark the requested configuration to be deleted. Once the
	// manager determines the configuration can be deleted it will be deleted.
	MarkForDeletion(ctx context.Context, in *MarkForDeletionRequest, opts ...grpc.CallOption) (*MarkForDeletionResponse, error)
	// ListConfigs will list all stored configuration that the manager manages.
	ListConfigs(ctx context.Context, in *ListConfigRequest, opts ...grpc.CallOption) (*ListConfigResponse, error)
	// GetConfig will retrieve the requested configuration by name.
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
	// NextTask will return the next available task to be worked on or a nil response if no task is available.
	NextTask(ctx context.Context, in *NextTaskRequest, opts ...grpc.CallOption) (*NextTaskResponse, error)
	// TaskUpdate will update the state of the requested task. This should be periodically called as the task
	// that was picked up by the NextTask RPC enters different stages of the build process.
	TaskUpdate(ctx context.Context, in *TaskUpdateRequest, opts ...grpc.CallOption) (*TaskUpdateResponse, error)
	// TaskComplete will update the state of the requested task to either Done or Error. Further, it will
	// update the current state of the clusters from the passed in value, so that subsequent tasks will
	// work with an up-to-date current state that reflects the actual state of the infrastructure.
	// This RPC  should be called when a task that has been previously picked up by the NextTask RPC
	// finished processing, either in error or successfully.
	TaskComplete(ctx context.Context, in *TaskCompleteRequest, opts ...grpc.CallOption) (*TaskCompleteResponse, error)
	// UpdateNodePool updates a single nodepool within a cluster, and should only be called by
	// the autoscaler-adapter service. This RPC bypasses the main loop of how changes are applied
	// to the configuration, and directly changes the nodepool to the state specified in the request
	// to initiate the build process.
	UpdateNodePool(ctx context.Context, in *UpdateNodePoolRequest, opts ...grpc.CallOption) (*UpdateNodePoolResponse, error)
}

type managerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerServiceClient(cc grpc.ClientConnInterface) ManagerServiceClient {
	return &managerServiceClient{cc}
}

func (c *managerServiceClient) UpsertManifest(ctx context.Context, in *UpsertManifestRequest, opts ...grpc.CallOption) (*UpsertManifestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpsertManifestResponse)
	err := c.cc.Invoke(ctx, ManagerService_UpsertManifest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServiceClient) MarkForDeletion(ctx context.Context, in *MarkForDeletionRequest, opts ...grpc.CallOption) (*MarkForDeletionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarkForDeletionResponse)
	err := c.cc.Invoke(ctx, ManagerService_MarkForDeletion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServiceClient) ListConfigs(ctx context.Context, in *ListConfigRequest, opts ...grpc.CallOption) (*ListConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListConfigResponse)
	err := c.cc.Invoke(ctx, ManagerService_ListConfigs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServiceClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, ManagerService_GetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServiceClient) NextTask(ctx context.Context, in *NextTaskRequest, opts ...grpc.CallOption) (*NextTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NextTaskResponse)
	err := c.cc.Invoke(ctx, ManagerService_NextTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServiceClient) TaskUpdate(ctx context.Context, in *TaskUpdateRequest, opts ...grpc.CallOption) (*TaskUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskUpdateResponse)
	err := c.cc.Invoke(ctx, ManagerService_TaskUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServiceClient) TaskComplete(ctx context.Context, in *TaskCompleteRequest, opts ...grpc.CallOption) (*TaskCompleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskCompleteResponse)
	err := c.cc.Invoke(ctx, ManagerService_TaskComplete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServiceClient) UpdateNodePool(ctx context.Context, in *UpdateNodePoolRequest, opts ...grpc.CallOption) (*UpdateNodePoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNodePoolResponse)
	err := c.cc.Invoke(ctx, ManagerService_UpdateNodePool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServiceServer is the server API for ManagerService service.
// All implementations must embed UnimplementedManagerServiceServer
// for forward compatibility.
type ManagerServiceServer interface {
	// UpsertManifest will process the request by either created a new configuration for the
	// given input manifest or updating an existing one.
	UpsertManifest(context.Context, *UpsertManifestRequest) (*UpsertManifestResponse, error)
	// MarkForDeletion will mark the requested configuration to be deleted. Once the
	// manager determines the configuration can be deleted it will be deleted.
	MarkForDeletion(context.Context, *MarkForDeletionRequest) (*MarkForDeletionResponse, error)
	// ListConfigs will list all stored configuration that the manager manages.
	ListConfigs(context.Context, *ListConfigRequest) (*ListConfigResponse, error)
	// GetConfig will retrieve the requested configuration by name.
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	// NextTask will return the next available task to be worked on or a nil response if no task is available.
	NextTask(context.Context, *NextTaskRequest) (*NextTaskResponse, error)
	// TaskUpdate will update the state of the requested task. This should be periodically called as the task
	// that was picked up by the NextTask RPC enters different stages of the build process.
	TaskUpdate(context.Context, *TaskUpdateRequest) (*TaskUpdateResponse, error)
	// TaskComplete will update the state of the requested task to either Done or Error. Further, it will
	// update the current state of the clusters from the passed in value, so that subsequent tasks will
	// work with an up-to-date current state that reflects the actual state of the infrastructure.
	// This RPC  should be called when a task that has been previously picked up by the NextTask RPC
	// finished processing, either in error or successfully.
	TaskComplete(context.Context, *TaskCompleteRequest) (*TaskCompleteResponse, error)
	// UpdateNodePool updates a single nodepool within a cluster, and should only be called by
	// the autoscaler-adapter service. This RPC bypasses the main loop of how changes are applied
	// to the configuration, and directly changes the nodepool to the state specified in the request
	// to initiate the build process.
	UpdateNodePool(context.Context, *UpdateNodePoolRequest) (*UpdateNodePoolResponse, error)
	mustEmbedUnimplementedManagerServiceServer()
}

// UnimplementedManagerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedManagerServiceServer struct{}

func (UnimplementedManagerServiceServer) UpsertManifest(context.Context, *UpsertManifestRequest) (*UpsertManifestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertManifest not implemented")
}
func (UnimplementedManagerServiceServer) MarkForDeletion(context.Context, *MarkForDeletionRequest) (*MarkForDeletionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkForDeletion not implemented")
}
func (UnimplementedManagerServiceServer) ListConfigs(context.Context, *ListConfigRequest) (*ListConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigs not implemented")
}
func (UnimplementedManagerServiceServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedManagerServiceServer) NextTask(context.Context, *NextTaskRequest) (*NextTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextTask not implemented")
}
func (UnimplementedManagerServiceServer) TaskUpdate(context.Context, *TaskUpdateRequest) (*TaskUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskUpdate not implemented")
}
func (UnimplementedManagerServiceServer) TaskComplete(context.Context, *TaskCompleteRequest) (*TaskCompleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskComplete not implemented")
}
func (UnimplementedManagerServiceServer) UpdateNodePool(context.Context, *UpdateNodePoolRequest) (*UpdateNodePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodePool not implemented")
}
func (UnimplementedManagerServiceServer) mustEmbedUnimplementedManagerServiceServer() {}
func (UnimplementedManagerServiceServer) testEmbeddedByValue()                        {}

// UnsafeManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServiceServer will
// result in compilation errors.
type UnsafeManagerServiceServer interface {
	mustEmbedUnimplementedManagerServiceServer()
}

func RegisterManagerServiceServer(s grpc.ServiceRegistrar, srv ManagerServiceServer) {
	// If the following call pancis, it indicates UnimplementedManagerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ManagerService_ServiceDesc, srv)
}

func _ManagerService_UpsertManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).UpsertManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_UpsertManifest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).UpsertManifest(ctx, req.(*UpsertManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerService_MarkForDeletion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkForDeletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).MarkForDeletion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_MarkForDeletion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).MarkForDeletion(ctx, req.(*MarkForDeletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerService_ListConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).ListConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_ListConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).ListConfigs(ctx, req.(*ListConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerService_NextTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).NextTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_NextTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).NextTask(ctx, req.(*NextTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerService_TaskUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).TaskUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_TaskUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).TaskUpdate(ctx, req.(*TaskUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerService_TaskComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).TaskComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_TaskComplete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).TaskComplete(ctx, req.(*TaskCompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerService_UpdateNodePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).UpdateNodePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_UpdateNodePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).UpdateNodePool(ctx, req.(*UpdateNodePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagerService_ServiceDesc is the grpc.ServiceDesc for ManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "claudie.ManagerService",
	HandlerType: (*ManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertManifest",
			Handler:    _ManagerService_UpsertManifest_Handler,
		},
		{
			MethodName: "MarkForDeletion",
			Handler:    _ManagerService_MarkForDeletion_Handler,
		},
		{
			MethodName: "ListConfigs",
			Handler:    _ManagerService_ListConfigs_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _ManagerService_GetConfig_Handler,
		},
		{
			MethodName: "NextTask",
			Handler:    _ManagerService_NextTask_Handler,
		},
		{
			MethodName: "TaskUpdate",
			Handler:    _ManagerService_TaskUpdate_Handler,
		},
		{
			MethodName: "TaskComplete",
			Handler:    _ManagerService_TaskComplete_Handler,
		},
		{
			MethodName: "UpdateNodePool",
			Handler:    _ManagerService_UpdateNodePool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}
