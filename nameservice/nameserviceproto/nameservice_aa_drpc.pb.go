// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.34
// source: nameservice/nameserviceproto/protos/nameservice_aa.proto

package nameserviceproto

import (
	context "context"
	errors "errors"
	drpc1 "github.com/planetscale/vtprotobuf/codec/drpc"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto struct{}

func (drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return drpc1.Marshal(msg)
}

func (drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return drpc1.Unmarshal(buf, msg)
}

func (drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	return drpc1.JSONMarshal(msg)
}

func (drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return drpc1.JSONUnmarshal(buf, msg)
}

type DRPCAnynsAccountAbstractionClient interface {
	DRPCConn() drpc.Conn

	GetOperation(ctx context.Context, in *GetOperationStatusRequest) (*OperationResponse, error)
	AdminFundUserAccount(ctx context.Context, in *AdminFundUserAccountRequestSigned) (*OperationResponse, error)
	AdminFundGasOperations(ctx context.Context, in *AdminFundGasOperationsRequestSigned) (*OperationResponse, error)
	GetUserAccount(ctx context.Context, in *GetUserAccountRequest) (*UserAccount, error)
	GetDataNameRegister(ctx context.Context, in *NameRegisterRequest) (*GetDataNameRegisterResponse, error)
	GetDataNameRegisterForSpace(ctx context.Context, in *NameRegisterForSpaceRequest) (*GetDataNameRegisterResponse, error)
	CreateUserOperation(ctx context.Context, in *CreateUserOperationRequestSigned) (*OperationResponse, error)
}

type drpcAnynsAccountAbstractionClient struct {
	cc drpc.Conn
}

func NewDRPCAnynsAccountAbstractionClient(cc drpc.Conn) DRPCAnynsAccountAbstractionClient {
	return &drpcAnynsAccountAbstractionClient{cc}
}

func (c *drpcAnynsAccountAbstractionClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcAnynsAccountAbstractionClient) GetOperation(ctx context.Context, in *GetOperationStatusRequest) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/AnynsAccountAbstraction/GetOperation", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcAnynsAccountAbstractionClient) AdminFundUserAccount(ctx context.Context, in *AdminFundUserAccountRequestSigned) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/AnynsAccountAbstraction/AdminFundUserAccount", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcAnynsAccountAbstractionClient) AdminFundGasOperations(ctx context.Context, in *AdminFundGasOperationsRequestSigned) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/AnynsAccountAbstraction/AdminFundGasOperations", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcAnynsAccountAbstractionClient) GetUserAccount(ctx context.Context, in *GetUserAccountRequest) (*UserAccount, error) {
	out := new(UserAccount)
	err := c.cc.Invoke(ctx, "/AnynsAccountAbstraction/GetUserAccount", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcAnynsAccountAbstractionClient) GetDataNameRegister(ctx context.Context, in *NameRegisterRequest) (*GetDataNameRegisterResponse, error) {
	out := new(GetDataNameRegisterResponse)
	err := c.cc.Invoke(ctx, "/AnynsAccountAbstraction/GetDataNameRegister", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcAnynsAccountAbstractionClient) GetDataNameRegisterForSpace(ctx context.Context, in *NameRegisterForSpaceRequest) (*GetDataNameRegisterResponse, error) {
	out := new(GetDataNameRegisterResponse)
	err := c.cc.Invoke(ctx, "/AnynsAccountAbstraction/GetDataNameRegisterForSpace", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcAnynsAccountAbstractionClient) CreateUserOperation(ctx context.Context, in *CreateUserOperationRequestSigned) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/AnynsAccountAbstraction/CreateUserOperation", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCAnynsAccountAbstractionServer interface {
	GetOperation(context.Context, *GetOperationStatusRequest) (*OperationResponse, error)
	AdminFundUserAccount(context.Context, *AdminFundUserAccountRequestSigned) (*OperationResponse, error)
	AdminFundGasOperations(context.Context, *AdminFundGasOperationsRequestSigned) (*OperationResponse, error)
	GetUserAccount(context.Context, *GetUserAccountRequest) (*UserAccount, error)
	GetDataNameRegister(context.Context, *NameRegisterRequest) (*GetDataNameRegisterResponse, error)
	GetDataNameRegisterForSpace(context.Context, *NameRegisterForSpaceRequest) (*GetDataNameRegisterResponse, error)
	CreateUserOperation(context.Context, *CreateUserOperationRequestSigned) (*OperationResponse, error)
}

type DRPCAnynsAccountAbstractionUnimplementedServer struct{}

func (s *DRPCAnynsAccountAbstractionUnimplementedServer) GetOperation(context.Context, *GetOperationStatusRequest) (*OperationResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCAnynsAccountAbstractionUnimplementedServer) AdminFundUserAccount(context.Context, *AdminFundUserAccountRequestSigned) (*OperationResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCAnynsAccountAbstractionUnimplementedServer) AdminFundGasOperations(context.Context, *AdminFundGasOperationsRequestSigned) (*OperationResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCAnynsAccountAbstractionUnimplementedServer) GetUserAccount(context.Context, *GetUserAccountRequest) (*UserAccount, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCAnynsAccountAbstractionUnimplementedServer) GetDataNameRegister(context.Context, *NameRegisterRequest) (*GetDataNameRegisterResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCAnynsAccountAbstractionUnimplementedServer) GetDataNameRegisterForSpace(context.Context, *NameRegisterForSpaceRequest) (*GetDataNameRegisterResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCAnynsAccountAbstractionUnimplementedServer) CreateUserOperation(context.Context, *CreateUserOperationRequestSigned) (*OperationResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCAnynsAccountAbstractionDescription struct{}

func (DRPCAnynsAccountAbstractionDescription) NumMethods() int { return 7 }

func (DRPCAnynsAccountAbstractionDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/AnynsAccountAbstraction/GetOperation", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCAnynsAccountAbstractionServer).
					GetOperation(
						ctx,
						in1.(*GetOperationStatusRequest),
					)
			}, DRPCAnynsAccountAbstractionServer.GetOperation, true
	case 1:
		return "/AnynsAccountAbstraction/AdminFundUserAccount", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCAnynsAccountAbstractionServer).
					AdminFundUserAccount(
						ctx,
						in1.(*AdminFundUserAccountRequestSigned),
					)
			}, DRPCAnynsAccountAbstractionServer.AdminFundUserAccount, true
	case 2:
		return "/AnynsAccountAbstraction/AdminFundGasOperations", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCAnynsAccountAbstractionServer).
					AdminFundGasOperations(
						ctx,
						in1.(*AdminFundGasOperationsRequestSigned),
					)
			}, DRPCAnynsAccountAbstractionServer.AdminFundGasOperations, true
	case 3:
		return "/AnynsAccountAbstraction/GetUserAccount", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCAnynsAccountAbstractionServer).
					GetUserAccount(
						ctx,
						in1.(*GetUserAccountRequest),
					)
			}, DRPCAnynsAccountAbstractionServer.GetUserAccount, true
	case 4:
		return "/AnynsAccountAbstraction/GetDataNameRegister", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCAnynsAccountAbstractionServer).
					GetDataNameRegister(
						ctx,
						in1.(*NameRegisterRequest),
					)
			}, DRPCAnynsAccountAbstractionServer.GetDataNameRegister, true
	case 5:
		return "/AnynsAccountAbstraction/GetDataNameRegisterForSpace", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCAnynsAccountAbstractionServer).
					GetDataNameRegisterForSpace(
						ctx,
						in1.(*NameRegisterForSpaceRequest),
					)
			}, DRPCAnynsAccountAbstractionServer.GetDataNameRegisterForSpace, true
	case 6:
		return "/AnynsAccountAbstraction/CreateUserOperation", drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCAnynsAccountAbstractionServer).
					CreateUserOperation(
						ctx,
						in1.(*CreateUserOperationRequestSigned),
					)
			}, DRPCAnynsAccountAbstractionServer.CreateUserOperation, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterAnynsAccountAbstraction(mux drpc.Mux, impl DRPCAnynsAccountAbstractionServer) error {
	return mux.Register(impl, DRPCAnynsAccountAbstractionDescription{})
}

type DRPCAnynsAccountAbstraction_GetOperationStream interface {
	drpc.Stream
	SendAndClose(*OperationResponse) error
}

type drpcAnynsAccountAbstraction_GetOperationStream struct {
	drpc.Stream
}

func (x *drpcAnynsAccountAbstraction_GetOperationStream) SendAndClose(m *OperationResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCAnynsAccountAbstraction_AdminFundUserAccountStream interface {
	drpc.Stream
	SendAndClose(*OperationResponse) error
}

type drpcAnynsAccountAbstraction_AdminFundUserAccountStream struct {
	drpc.Stream
}

func (x *drpcAnynsAccountAbstraction_AdminFundUserAccountStream) SendAndClose(m *OperationResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCAnynsAccountAbstraction_AdminFundGasOperationsStream interface {
	drpc.Stream
	SendAndClose(*OperationResponse) error
}

type drpcAnynsAccountAbstraction_AdminFundGasOperationsStream struct {
	drpc.Stream
}

func (x *drpcAnynsAccountAbstraction_AdminFundGasOperationsStream) SendAndClose(m *OperationResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCAnynsAccountAbstraction_GetUserAccountStream interface {
	drpc.Stream
	SendAndClose(*UserAccount) error
}

type drpcAnynsAccountAbstraction_GetUserAccountStream struct {
	drpc.Stream
}

func (x *drpcAnynsAccountAbstraction_GetUserAccountStream) SendAndClose(m *UserAccount) error {
	if err := x.MsgSend(m, drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCAnynsAccountAbstraction_GetDataNameRegisterStream interface {
	drpc.Stream
	SendAndClose(*GetDataNameRegisterResponse) error
}

type drpcAnynsAccountAbstraction_GetDataNameRegisterStream struct {
	drpc.Stream
}

func (x *drpcAnynsAccountAbstraction_GetDataNameRegisterStream) SendAndClose(m *GetDataNameRegisterResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCAnynsAccountAbstraction_GetDataNameRegisterForSpaceStream interface {
	drpc.Stream
	SendAndClose(*GetDataNameRegisterResponse) error
}

type drpcAnynsAccountAbstraction_GetDataNameRegisterForSpaceStream struct {
	drpc.Stream
}

func (x *drpcAnynsAccountAbstraction_GetDataNameRegisterForSpaceStream) SendAndClose(m *GetDataNameRegisterResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCAnynsAccountAbstraction_CreateUserOperationStream interface {
	drpc.Stream
	SendAndClose(*OperationResponse) error
}

type drpcAnynsAccountAbstraction_CreateUserOperationStream struct {
	drpc.Stream
}

func (x *drpcAnynsAccountAbstraction_CreateUserOperationStream) SendAndClose(m *OperationResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_nameservice_nameserviceproto_protos_nameservice_aa_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
