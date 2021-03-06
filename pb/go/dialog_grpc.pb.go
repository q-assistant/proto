// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DialogClient is the client API for Dialog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DialogClient interface {
	// Sends a greeting
	Recognize(ctx context.Context, in *RecognizeRequest, opts ...grpc.CallOption) (*RecognizeResponse, error)
}

type dialogClient struct {
	cc grpc.ClientConnInterface
}

func NewDialogClient(cc grpc.ClientConnInterface) DialogClient {
	return &dialogClient{cc}
}

func (c *dialogClient) Recognize(ctx context.Context, in *RecognizeRequest, opts ...grpc.CallOption) (*RecognizeResponse, error) {
	out := new(RecognizeResponse)
	err := c.cc.Invoke(ctx, "/dialog.Dialog/Recognize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DialogServer is the server API for Dialog service.
// All implementations must embed UnimplementedDialogServer
// for forward compatibility
type DialogServer interface {
	// Sends a greeting
	Recognize(context.Context, *RecognizeRequest) (*RecognizeResponse, error)
	mustEmbedUnimplementedDialogServer()
}

// UnimplementedDialogServer must be embedded to have forward compatible implementations.
type UnimplementedDialogServer struct {
}

func (UnimplementedDialogServer) Recognize(context.Context, *RecognizeRequest) (*RecognizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recognize not implemented")
}
func (UnimplementedDialogServer) mustEmbedUnimplementedDialogServer() {}

// UnsafeDialogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DialogServer will
// result in compilation errors.
type UnsafeDialogServer interface {
	mustEmbedUnimplementedDialogServer()
}

func RegisterDialogServer(s grpc.ServiceRegistrar, srv DialogServer) {
	s.RegisterService(&_Dialog_serviceDesc, srv)
}

func _Dialog_Recognize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecognizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DialogServer).Recognize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Dialog/Recognize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DialogServer).Recognize(ctx, req.(*RecognizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dialog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dialog.Dialog",
	HandlerType: (*DialogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recognize",
			Handler:    _Dialog_Recognize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dialog.proto",
}
