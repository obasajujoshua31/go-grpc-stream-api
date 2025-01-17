// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorClient interface {
	Add(ctx context.Context, opts ...grpc.CallOption) (Calculator_AddClient, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

var calculatorAddStreamDesc = &grpc.StreamDesc{
	StreamName:    "Add",
	ClientStreams: true,
}

func (c *calculatorClient) Add(ctx context.Context, opts ...grpc.CallOption) (Calculator_AddClient, error) {
	stream, err := c.cc.NewStream(ctx, calculatorAddStreamDesc, "/calculator/Add", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorAddClient{stream}
	return x, nil
}

type Calculator_AddClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type calculatorAddClient struct {
	grpc.ClientStream
}

func (x *calculatorAddClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorAddClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorService is the service API for Calculator service.
// Fields should be assigned to their respective handler implementations only before
// RegisterCalculatorService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type CalculatorService struct {
	Add func(Calculator_AddServer) error
}

func (s *CalculatorService) add(_ interface{}, stream grpc.ServerStream) error {
	return s.Add(&calculatorAddServer{stream})
}

type Calculator_AddServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type calculatorAddServer struct {
	grpc.ServerStream
}

func (x *calculatorAddServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorAddServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterCalculatorService registers a service implementation with a gRPC server.
func RegisterCalculatorService(s grpc.ServiceRegistrar, srv *CalculatorService) {
	srvCopy := *srv
	if srvCopy.Add == nil {
		srvCopy.Add = func(Calculator_AddServer) error {
			return status.Errorf(codes.Unimplemented, "method Add not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "calculator",
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Add",
				Handler:       srvCopy.add,
				ClientStreams: true,
			},
		},
		Metadata: "proto/spec.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewCalculatorService creates a new CalculatorService containing the
// implemented methods of the Calculator service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewCalculatorService(s interface{}) *CalculatorService {
	ns := &CalculatorService{}
	if h, ok := s.(interface {
		Add(Calculator_AddServer) error
	}); ok {
		ns.Add = h.Add
	}
	return ns
}

// UnstableCalculatorService is the service API for Calculator service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableCalculatorService interface {
	Add(Calculator_AddServer) error
}
