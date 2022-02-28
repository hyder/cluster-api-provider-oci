// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/aribandy/work/src/sigs.k8s.io/cluster-api-provider-oci/cloud/services/identity/client.go

// Package mock_identity is a generated GoMock package.
package mock_identity

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	identity "github.com/oracle/oci-go-sdk/v53/identity"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ListAvailabilityDomains mocks base method.
func (m *MockClient) ListAvailabilityDomains(ctx context.Context, request identity.ListAvailabilityDomainsRequest) (identity.ListAvailabilityDomainsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAvailabilityDomains", ctx, request)
	ret0, _ := ret[0].(identity.ListAvailabilityDomainsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAvailabilityDomains indicates an expected call of ListAvailabilityDomains.
func (mr *MockClientMockRecorder) ListAvailabilityDomains(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailabilityDomains", reflect.TypeOf((*MockClient)(nil).ListAvailabilityDomains), ctx, request)
}

// ListFaultDomains mocks base method.
func (m *MockClient) ListFaultDomains(ctx context.Context, request identity.ListFaultDomainsRequest) (identity.ListFaultDomainsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFaultDomains", ctx, request)
	ret0, _ := ret[0].(identity.ListFaultDomainsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFaultDomains indicates an expected call of ListFaultDomains.
func (mr *MockClientMockRecorder) ListFaultDomains(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFaultDomains", reflect.TypeOf((*MockClient)(nil).ListFaultDomains), ctx, request)
}

// ListRegions mocks base method.
func (m *MockClient) ListRegions(ctx context.Context) (identity.ListRegionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRegions", ctx)
	ret0, _ := ret[0].(identity.ListRegionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegions indicates an expected call of ListRegions.
func (mr *MockClientMockRecorder) ListRegions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegions", reflect.TypeOf((*MockClient)(nil).ListRegions), ctx)
}
