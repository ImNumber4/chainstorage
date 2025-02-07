// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coinbase/chainstorage/internal/storage/blobstorage/downloader (interfaces: BlockDownloader)
//
// Generated by this command:
//
//	mockgen -destination internal/storage/blobstorage/downloader/mocks/mocks.go -package downloadermocks github.com/coinbase/chainstorage/internal/storage/blobstorage/downloader BlockDownloader
//
// Package downloadermocks is a generated GoMock package.
package downloadermocks

import (
	context "context"
	reflect "reflect"

	chainstorage "github.com/coinbase/chainstorage/protos/coinbase/chainstorage"
	gomock "go.uber.org/mock/gomock"
)

// MockBlockDownloader is a mock of BlockDownloader interface.
type MockBlockDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockBlockDownloaderMockRecorder
}

// MockBlockDownloaderMockRecorder is the mock recorder for MockBlockDownloader.
type MockBlockDownloaderMockRecorder struct {
	mock *MockBlockDownloader
}

// NewMockBlockDownloader creates a new mock instance.
func NewMockBlockDownloader(ctrl *gomock.Controller) *MockBlockDownloader {
	mock := &MockBlockDownloader{ctrl: ctrl}
	mock.recorder = &MockBlockDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockDownloader) EXPECT() *MockBlockDownloaderMockRecorder {
	return m.recorder
}

// Download mocks base method.
func (m *MockBlockDownloader) Download(arg0 context.Context, arg1 *chainstorage.BlockFile) (*chainstorage.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1)
	ret0, _ := ret[0].(*chainstorage.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockBlockDownloaderMockRecorder) Download(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockBlockDownloader)(nil).Download), arg0, arg1)
}
