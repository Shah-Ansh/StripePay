// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: payment.proto

package proto

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
	PaymentGateway_Register_FullMethodName      = "/payment.PaymentGateway/Register"
	PaymentGateway_Authenticate_FullMethodName  = "/payment.PaymentGateway/Authenticate"
	PaymentGateway_TransferMoney_FullMethodName = "/payment.PaymentGateway/TransferMoney"
	PaymentGateway_CheckBalance_FullMethodName  = "/payment.PaymentGateway/CheckBalance"
	PaymentGateway_BankRegister_FullMethodName  = "/payment.PaymentGateway/BankRegister"
)

// PaymentGatewayClient is the client API for PaymentGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Payment Gateway Service
type PaymentGatewayClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	TransferMoney(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	CheckBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	BankRegister(ctx context.Context, in *BankRegisterRequest, opts ...grpc.CallOption) (*BankRegisterResponse, error)
}

type paymentGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentGatewayClient(cc grpc.ClientConnInterface) PaymentGatewayClient {
	return &paymentGatewayClient{cc}
}

func (c *paymentGatewayClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_Authenticate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) TransferMoney(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_TransferMoney_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) CheckBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_CheckBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) BankRegister(ctx context.Context, in *BankRegisterRequest, opts ...grpc.CallOption) (*BankRegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BankRegisterResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_BankRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentGatewayServer is the server API for PaymentGateway service.
// All implementations must embed UnimplementedPaymentGatewayServer
// for forward compatibility.
//
// Payment Gateway Service
type PaymentGatewayServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Authenticate(context.Context, *AuthRequest) (*AuthResponse, error)
	TransferMoney(context.Context, *TransferRequest) (*TransferResponse, error)
	CheckBalance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	BankRegister(context.Context, *BankRegisterRequest) (*BankRegisterResponse, error)
	mustEmbedUnimplementedPaymentGatewayServer()
}

// UnimplementedPaymentGatewayServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentGatewayServer struct{}

func (UnimplementedPaymentGatewayServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedPaymentGatewayServer) Authenticate(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedPaymentGatewayServer) TransferMoney(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferMoney not implemented")
}
func (UnimplementedPaymentGatewayServer) CheckBalance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBalance not implemented")
}
func (UnimplementedPaymentGatewayServer) BankRegister(context.Context, *BankRegisterRequest) (*BankRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BankRegister not implemented")
}
func (UnimplementedPaymentGatewayServer) mustEmbedUnimplementedPaymentGatewayServer() {}
func (UnimplementedPaymentGatewayServer) testEmbeddedByValue()                        {}

// UnsafePaymentGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentGatewayServer will
// result in compilation errors.
type UnsafePaymentGatewayServer interface {
	mustEmbedUnimplementedPaymentGatewayServer()
}

func RegisterPaymentGatewayServer(s grpc.ServiceRegistrar, srv PaymentGatewayServer) {
	// If the following call pancis, it indicates UnimplementedPaymentGatewayServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentGateway_ServiceDesc, srv)
}

func _PaymentGateway_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_TransferMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).TransferMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_TransferMoney_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).TransferMoney(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_CheckBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).CheckBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_CheckBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).CheckBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_BankRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).BankRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_BankRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).BankRegister(ctx, req.(*BankRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentGateway_ServiceDesc is the grpc.ServiceDesc for PaymentGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymentGateway",
	HandlerType: (*PaymentGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PaymentGateway_Register_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _PaymentGateway_Authenticate_Handler,
		},
		{
			MethodName: "TransferMoney",
			Handler:    _PaymentGateway_TransferMoney_Handler,
		},
		{
			MethodName: "CheckBalance",
			Handler:    _PaymentGateway_CheckBalance_Handler,
		},
		{
			MethodName: "BankRegister",
			Handler:    _PaymentGateway_BankRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}

const (
	Bank_ProcessTransaction_FullMethodName = "/payment.Bank/ProcessTransaction"
	Bank_GetBalance_FullMethodName         = "/payment.Bank/GetBalance"
	Bank_DebitAccount_FullMethodName       = "/payment.Bank/DebitAccount"
	Bank_CreditAccount_FullMethodName      = "/payment.Bank/CreditAccount"
	Bank_PrepareDebit_FullMethodName       = "/payment.Bank/PrepareDebit"
	Bank_CommitDebit_FullMethodName        = "/payment.Bank/CommitDebit"
	Bank_AbortDebit_FullMethodName         = "/payment.Bank/AbortDebit"
	Bank_PrepareCredit_FullMethodName      = "/payment.Bank/PrepareCredit"
	Bank_CommitCredit_FullMethodName       = "/payment.Bank/CommitCredit"
	Bank_AbortCredit_FullMethodName        = "/payment.Bank/AbortCredit"
)

// BankClient is the client API for Bank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Bank Service
type BankClient interface {
	ProcessTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	// Existing operations.
	DebitAccount(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error)
	CreditAccount(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error)
	// New 2PC RPCs.
	PrepareDebit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error)
	CommitDebit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error)
	AbortDebit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error)
	PrepareCredit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error)
	CommitCredit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error)
	AbortCredit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error)
}

type bankClient struct {
	cc grpc.ClientConnInterface
}

func NewBankClient(cc grpc.ClientConnInterface) BankClient {
	return &bankClient{cc}
}

func (c *bankClient) ProcessTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, Bank_ProcessTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, Bank_GetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) DebitAccount(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DebitCreditResponse)
	err := c.cc.Invoke(ctx, Bank_DebitAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) CreditAccount(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DebitCreditResponse)
	err := c.cc.Invoke(ctx, Bank_CreditAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) PrepareDebit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DebitCreditResponse)
	err := c.cc.Invoke(ctx, Bank_PrepareDebit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) CommitDebit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DebitCreditResponse)
	err := c.cc.Invoke(ctx, Bank_CommitDebit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) AbortDebit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DebitCreditResponse)
	err := c.cc.Invoke(ctx, Bank_AbortDebit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) PrepareCredit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DebitCreditResponse)
	err := c.cc.Invoke(ctx, Bank_PrepareCredit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) CommitCredit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DebitCreditResponse)
	err := c.cc.Invoke(ctx, Bank_CommitCredit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) AbortCredit(ctx context.Context, in *DebitCreditRequest, opts ...grpc.CallOption) (*DebitCreditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DebitCreditResponse)
	err := c.cc.Invoke(ctx, Bank_AbortCredit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankServer is the server API for Bank service.
// All implementations must embed UnimplementedBankServer
// for forward compatibility.
//
// Bank Service
type BankServer interface {
	ProcessTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error)
	GetBalance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	// Existing operations.
	DebitAccount(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error)
	CreditAccount(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error)
	// New 2PC RPCs.
	PrepareDebit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error)
	CommitDebit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error)
	AbortDebit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error)
	PrepareCredit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error)
	CommitCredit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error)
	AbortCredit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error)
	mustEmbedUnimplementedBankServer()
}

// UnimplementedBankServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBankServer struct{}

func (UnimplementedBankServer) ProcessTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessTransaction not implemented")
}
func (UnimplementedBankServer) GetBalance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedBankServer) DebitAccount(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DebitAccount not implemented")
}
func (UnimplementedBankServer) CreditAccount(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreditAccount not implemented")
}
func (UnimplementedBankServer) PrepareDebit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareDebit not implemented")
}
func (UnimplementedBankServer) CommitDebit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitDebit not implemented")
}
func (UnimplementedBankServer) AbortDebit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortDebit not implemented")
}
func (UnimplementedBankServer) PrepareCredit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareCredit not implemented")
}
func (UnimplementedBankServer) CommitCredit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitCredit not implemented")
}
func (UnimplementedBankServer) AbortCredit(context.Context, *DebitCreditRequest) (*DebitCreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortCredit not implemented")
}
func (UnimplementedBankServer) mustEmbedUnimplementedBankServer() {}
func (UnimplementedBankServer) testEmbeddedByValue()              {}

// UnsafeBankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankServer will
// result in compilation errors.
type UnsafeBankServer interface {
	mustEmbedUnimplementedBankServer()
}

func RegisterBankServer(s grpc.ServiceRegistrar, srv BankServer) {
	// If the following call pancis, it indicates UnimplementedBankServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Bank_ServiceDesc, srv)
}

func _Bank_ProcessTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).ProcessTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_ProcessTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).ProcessTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).GetBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_DebitAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitCreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).DebitAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_DebitAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).DebitAccount(ctx, req.(*DebitCreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_CreditAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitCreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).CreditAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_CreditAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).CreditAccount(ctx, req.(*DebitCreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_PrepareDebit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitCreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).PrepareDebit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_PrepareDebit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).PrepareDebit(ctx, req.(*DebitCreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_CommitDebit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitCreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).CommitDebit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_CommitDebit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).CommitDebit(ctx, req.(*DebitCreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_AbortDebit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitCreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).AbortDebit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_AbortDebit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).AbortDebit(ctx, req.(*DebitCreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_PrepareCredit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitCreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).PrepareCredit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_PrepareCredit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).PrepareCredit(ctx, req.(*DebitCreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_CommitCredit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitCreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).CommitCredit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_CommitCredit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).CommitCredit(ctx, req.(*DebitCreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_AbortCredit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitCreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).AbortCredit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_AbortCredit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).AbortCredit(ctx, req.(*DebitCreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bank_ServiceDesc is the grpc.ServiceDesc for Bank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.Bank",
	HandlerType: (*BankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessTransaction",
			Handler:    _Bank_ProcessTransaction_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Bank_GetBalance_Handler,
		},
		{
			MethodName: "DebitAccount",
			Handler:    _Bank_DebitAccount_Handler,
		},
		{
			MethodName: "CreditAccount",
			Handler:    _Bank_CreditAccount_Handler,
		},
		{
			MethodName: "PrepareDebit",
			Handler:    _Bank_PrepareDebit_Handler,
		},
		{
			MethodName: "CommitDebit",
			Handler:    _Bank_CommitDebit_Handler,
		},
		{
			MethodName: "AbortDebit",
			Handler:    _Bank_AbortDebit_Handler,
		},
		{
			MethodName: "PrepareCredit",
			Handler:    _Bank_PrepareCredit_Handler,
		},
		{
			MethodName: "CommitCredit",
			Handler:    _Bank_CommitCredit_Handler,
		},
		{
			MethodName: "AbortCredit",
			Handler:    _Bank_AbortCredit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}
