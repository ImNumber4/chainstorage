// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coinbase/chainstorage/protos/coinbase/chainstorage (interfaces: ChainStorageClient,ChainStorage_StreamChainEventsClient)
//
// Generated by this command:
//
//	mockgen -destination protos/coinbase/chainstorage/mocks/mocks.go -package chainstoragemocks github.com/coinbase/chainstorage/protos/coinbase/chainstorage ChainStorageClient,ChainStorage_StreamChainEventsClient
//
// Package chainstoragemocks is a generated GoMock package.
package chainstoragemocks

import (
	context "context"
	reflect "reflect"

	chainstorage "github.com/coinbase/chainstorage/protos/coinbase/chainstorage"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockChainStorageClient is a mock of ChainStorageClient interface.
type MockChainStorageClient struct {
	ctrl     *gomock.Controller
	recorder *MockChainStorageClientMockRecorder
}

// MockChainStorageClientMockRecorder is the mock recorder for MockChainStorageClient.
type MockChainStorageClientMockRecorder struct {
	mock *MockChainStorageClient
}

// NewMockChainStorageClient creates a new mock instance.
func NewMockChainStorageClient(ctrl *gomock.Controller) *MockChainStorageClient {
	mock := &MockChainStorageClient{ctrl: ctrl}
	mock.recorder = &MockChainStorageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainStorageClient) EXPECT() *MockChainStorageClientMockRecorder {
	return m.recorder
}

// GetBlockByTransaction mocks base method.
func (m *MockChainStorageClient) GetBlockByTransaction(arg0 context.Context, arg1 *chainstorage.GetBlockByTransactionRequest, arg2 ...grpc.CallOption) (*chainstorage.GetBlockByTransactionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockByTransaction", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetBlockByTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByTransaction indicates an expected call of GetBlockByTransaction.
func (mr *MockChainStorageClientMockRecorder) GetBlockByTransaction(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByTransaction", reflect.TypeOf((*MockChainStorageClient)(nil).GetBlockByTransaction), varargs...)
}

// GetBlockFile mocks base method.
func (m *MockChainStorageClient) GetBlockFile(arg0 context.Context, arg1 *chainstorage.GetBlockFileRequest, arg2 ...grpc.CallOption) (*chainstorage.GetBlockFileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockFile", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetBlockFileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockFile indicates an expected call of GetBlockFile.
func (mr *MockChainStorageClientMockRecorder) GetBlockFile(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockFile", reflect.TypeOf((*MockChainStorageClient)(nil).GetBlockFile), varargs...)
}

// GetBlockFilesByRange mocks base method.
func (m *MockChainStorageClient) GetBlockFilesByRange(arg0 context.Context, arg1 *chainstorage.GetBlockFilesByRangeRequest, arg2 ...grpc.CallOption) (*chainstorage.GetBlockFilesByRangeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockFilesByRange", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetBlockFilesByRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockFilesByRange indicates an expected call of GetBlockFilesByRange.
func (mr *MockChainStorageClientMockRecorder) GetBlockFilesByRange(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockFilesByRange", reflect.TypeOf((*MockChainStorageClient)(nil).GetBlockFilesByRange), varargs...)
}

// GetChainEvents mocks base method.
func (m *MockChainStorageClient) GetChainEvents(arg0 context.Context, arg1 *chainstorage.GetChainEventsRequest, arg2 ...grpc.CallOption) (*chainstorage.GetChainEventsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChainEvents", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetChainEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainEvents indicates an expected call of GetChainEvents.
func (mr *MockChainStorageClientMockRecorder) GetChainEvents(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainEvents", reflect.TypeOf((*MockChainStorageClient)(nil).GetChainEvents), varargs...)
}

// GetChainMetadata mocks base method.
func (m *MockChainStorageClient) GetChainMetadata(arg0 context.Context, arg1 *chainstorage.GetChainMetadataRequest, arg2 ...grpc.CallOption) (*chainstorage.GetChainMetadataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChainMetadata", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetChainMetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainMetadata indicates an expected call of GetChainMetadata.
func (mr *MockChainStorageClientMockRecorder) GetChainMetadata(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainMetadata", reflect.TypeOf((*MockChainStorageClient)(nil).GetChainMetadata), varargs...)
}

// GetLatestBlock mocks base method.
func (m *MockChainStorageClient) GetLatestBlock(arg0 context.Context, arg1 *chainstorage.GetLatestBlockRequest, arg2 ...grpc.CallOption) (*chainstorage.GetLatestBlockResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLatestBlock", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetLatestBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestBlock indicates an expected call of GetLatestBlock.
func (mr *MockChainStorageClientMockRecorder) GetLatestBlock(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestBlock", reflect.TypeOf((*MockChainStorageClient)(nil).GetLatestBlock), varargs...)
}

// GetNativeBlock mocks base method.
func (m *MockChainStorageClient) GetNativeBlock(arg0 context.Context, arg1 *chainstorage.GetNativeBlockRequest, arg2 ...grpc.CallOption) (*chainstorage.GetNativeBlockResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNativeBlock", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetNativeBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNativeBlock indicates an expected call of GetNativeBlock.
func (mr *MockChainStorageClientMockRecorder) GetNativeBlock(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNativeBlock", reflect.TypeOf((*MockChainStorageClient)(nil).GetNativeBlock), varargs...)
}

// GetNativeBlocksByRange mocks base method.
func (m *MockChainStorageClient) GetNativeBlocksByRange(arg0 context.Context, arg1 *chainstorage.GetNativeBlocksByRangeRequest, arg2 ...grpc.CallOption) (*chainstorage.GetNativeBlocksByRangeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNativeBlocksByRange", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetNativeBlocksByRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNativeBlocksByRange indicates an expected call of GetNativeBlocksByRange.
func (mr *MockChainStorageClientMockRecorder) GetNativeBlocksByRange(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNativeBlocksByRange", reflect.TypeOf((*MockChainStorageClient)(nil).GetNativeBlocksByRange), varargs...)
}

// GetNativeTransaction mocks base method.
func (m *MockChainStorageClient) GetNativeTransaction(arg0 context.Context, arg1 *chainstorage.GetNativeTransactionRequest, arg2 ...grpc.CallOption) (*chainstorage.GetNativeTransactionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNativeTransaction", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetNativeTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNativeTransaction indicates an expected call of GetNativeTransaction.
func (mr *MockChainStorageClientMockRecorder) GetNativeTransaction(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNativeTransaction", reflect.TypeOf((*MockChainStorageClient)(nil).GetNativeTransaction), varargs...)
}

// GetRawBlock mocks base method.
func (m *MockChainStorageClient) GetRawBlock(arg0 context.Context, arg1 *chainstorage.GetRawBlockRequest, arg2 ...grpc.CallOption) (*chainstorage.GetRawBlockResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRawBlock", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetRawBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawBlock indicates an expected call of GetRawBlock.
func (mr *MockChainStorageClientMockRecorder) GetRawBlock(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawBlock", reflect.TypeOf((*MockChainStorageClient)(nil).GetRawBlock), varargs...)
}

// GetRawBlocksByRange mocks base method.
func (m *MockChainStorageClient) GetRawBlocksByRange(arg0 context.Context, arg1 *chainstorage.GetRawBlocksByRangeRequest, arg2 ...grpc.CallOption) (*chainstorage.GetRawBlocksByRangeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRawBlocksByRange", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetRawBlocksByRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawBlocksByRange indicates an expected call of GetRawBlocksByRange.
func (mr *MockChainStorageClientMockRecorder) GetRawBlocksByRange(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawBlocksByRange", reflect.TypeOf((*MockChainStorageClient)(nil).GetRawBlocksByRange), varargs...)
}

// GetRosettaBlock mocks base method.
func (m *MockChainStorageClient) GetRosettaBlock(arg0 context.Context, arg1 *chainstorage.GetRosettaBlockRequest, arg2 ...grpc.CallOption) (*chainstorage.GetRosettaBlockResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRosettaBlock", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetRosettaBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRosettaBlock indicates an expected call of GetRosettaBlock.
func (mr *MockChainStorageClientMockRecorder) GetRosettaBlock(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRosettaBlock", reflect.TypeOf((*MockChainStorageClient)(nil).GetRosettaBlock), varargs...)
}

// GetRosettaBlocksByRange mocks base method.
func (m *MockChainStorageClient) GetRosettaBlocksByRange(arg0 context.Context, arg1 *chainstorage.GetRosettaBlocksByRangeRequest, arg2 ...grpc.CallOption) (*chainstorage.GetRosettaBlocksByRangeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRosettaBlocksByRange", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetRosettaBlocksByRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRosettaBlocksByRange indicates an expected call of GetRosettaBlocksByRange.
func (mr *MockChainStorageClientMockRecorder) GetRosettaBlocksByRange(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRosettaBlocksByRange", reflect.TypeOf((*MockChainStorageClient)(nil).GetRosettaBlocksByRange), varargs...)
}

// GetVerifiedAccountState mocks base method.
func (m *MockChainStorageClient) GetVerifiedAccountState(arg0 context.Context, arg1 *chainstorage.GetVerifiedAccountStateRequest, arg2 ...grpc.CallOption) (*chainstorage.GetVerifiedAccountStateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVerifiedAccountState", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetVerifiedAccountStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVerifiedAccountState indicates an expected call of GetVerifiedAccountState.
func (mr *MockChainStorageClientMockRecorder) GetVerifiedAccountState(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVerifiedAccountState", reflect.TypeOf((*MockChainStorageClient)(nil).GetVerifiedAccountState), varargs...)
}

// GetVersionedChainEvent mocks base method.
func (m *MockChainStorageClient) GetVersionedChainEvent(arg0 context.Context, arg1 *chainstorage.GetVersionedChainEventRequest, arg2 ...grpc.CallOption) (*chainstorage.GetVersionedChainEventResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVersionedChainEvent", varargs...)
	ret0, _ := ret[0].(*chainstorage.GetVersionedChainEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersionedChainEvent indicates an expected call of GetVersionedChainEvent.
func (mr *MockChainStorageClientMockRecorder) GetVersionedChainEvent(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionedChainEvent", reflect.TypeOf((*MockChainStorageClient)(nil).GetVersionedChainEvent), varargs...)
}

// StreamChainEvents mocks base method.
func (m *MockChainStorageClient) StreamChainEvents(arg0 context.Context, arg1 *chainstorage.ChainEventsRequest, arg2 ...grpc.CallOption) (chainstorage.ChainStorage_StreamChainEventsClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamChainEvents", varargs...)
	ret0, _ := ret[0].(chainstorage.ChainStorage_StreamChainEventsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamChainEvents indicates an expected call of StreamChainEvents.
func (mr *MockChainStorageClientMockRecorder) StreamChainEvents(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamChainEvents", reflect.TypeOf((*MockChainStorageClient)(nil).StreamChainEvents), varargs...)
}

// MockChainStorage_StreamChainEventsClient is a mock of ChainStorage_StreamChainEventsClient interface.
type MockChainStorage_StreamChainEventsClient struct {
	ctrl     *gomock.Controller
	recorder *MockChainStorage_StreamChainEventsClientMockRecorder
}

// MockChainStorage_StreamChainEventsClientMockRecorder is the mock recorder for MockChainStorage_StreamChainEventsClient.
type MockChainStorage_StreamChainEventsClientMockRecorder struct {
	mock *MockChainStorage_StreamChainEventsClient
}

// NewMockChainStorage_StreamChainEventsClient creates a new mock instance.
func NewMockChainStorage_StreamChainEventsClient(ctrl *gomock.Controller) *MockChainStorage_StreamChainEventsClient {
	mock := &MockChainStorage_StreamChainEventsClient{ctrl: ctrl}
	mock.recorder = &MockChainStorage_StreamChainEventsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainStorage_StreamChainEventsClient) EXPECT() *MockChainStorage_StreamChainEventsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockChainStorage_StreamChainEventsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockChainStorage_StreamChainEventsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockChainStorage_StreamChainEventsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockChainStorage_StreamChainEventsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockChainStorage_StreamChainEventsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChainStorage_StreamChainEventsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockChainStorage_StreamChainEventsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockChainStorage_StreamChainEventsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockChainStorage_StreamChainEventsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockChainStorage_StreamChainEventsClient) Recv() (*chainstorage.ChainEventsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*chainstorage.ChainEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockChainStorage_StreamChainEventsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockChainStorage_StreamChainEventsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockChainStorage_StreamChainEventsClient) RecvMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockChainStorage_StreamChainEventsClientMockRecorder) RecvMsg(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChainStorage_StreamChainEventsClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockChainStorage_StreamChainEventsClient) SendMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockChainStorage_StreamChainEventsClientMockRecorder) SendMsg(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChainStorage_StreamChainEventsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockChainStorage_StreamChainEventsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockChainStorage_StreamChainEventsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockChainStorage_StreamChainEventsClient)(nil).Trailer))
}
