// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/aribandy/work/src/sigs.k8s.io/cluster-api-provider-oci/cloud/scope/cluster_client.go

// Package mock_scope is a generated GoMock package.
package mock_scope

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/oracle/cluster-api-provider-oci/api/v1beta1"
)

// MockClusterScopeClient is a mock of ClusterScopeClient interface.
type MockClusterScopeClient struct {
	ctrl     *gomock.Controller
	recorder *MockClusterScopeClientMockRecorder
}

// MockClusterScopeClientMockRecorder is the mock recorder for MockClusterScopeClient.
type MockClusterScopeClientMockRecorder struct {
	mock *MockClusterScopeClient
}

// NewMockClusterScopeClient creates a new mock instance.
func NewMockClusterScopeClient(ctrl *gomock.Controller) *MockClusterScopeClient {
	mock := &MockClusterScopeClient{ctrl: ctrl}
	mock.recorder = &MockClusterScopeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterScopeClient) EXPECT() *MockClusterScopeClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClusterScopeClient) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClusterScopeClientMockRecorder) Close(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClusterScopeClient)(nil).Close), ctx)
}

// DeleteApiServerLB mocks base method.
func (m *MockClusterScopeClient) DeleteApiServerLB(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApiServerLB", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApiServerLB indicates an expected call of DeleteApiServerLB.
func (mr *MockClusterScopeClientMockRecorder) DeleteApiServerLB(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApiServerLB", reflect.TypeOf((*MockClusterScopeClient)(nil).DeleteApiServerLB), ctx)
}

// DeleteInternetGateway mocks base method.
func (m *MockClusterScopeClient) DeleteInternetGateway(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInternetGateway", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInternetGateway indicates an expected call of DeleteInternetGateway.
func (mr *MockClusterScopeClientMockRecorder) DeleteInternetGateway(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInternetGateway", reflect.TypeOf((*MockClusterScopeClient)(nil).DeleteInternetGateway), ctx)
}

// DeleteNSGs mocks base method.
func (m *MockClusterScopeClient) DeleteNSGs(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNSGs", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNSGs indicates an expected call of DeleteNSGs.
func (mr *MockClusterScopeClientMockRecorder) DeleteNSGs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNSGs", reflect.TypeOf((*MockClusterScopeClient)(nil).DeleteNSGs), ctx)
}

// DeleteNatGateway mocks base method.
func (m *MockClusterScopeClient) DeleteNatGateway(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNatGateway", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNatGateway indicates an expected call of DeleteNatGateway.
func (mr *MockClusterScopeClientMockRecorder) DeleteNatGateway(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNatGateway", reflect.TypeOf((*MockClusterScopeClient)(nil).DeleteNatGateway), ctx)
}

// DeleteRouteTables mocks base method.
func (m *MockClusterScopeClient) DeleteRouteTables(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRouteTables", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRouteTables indicates an expected call of DeleteRouteTables.
func (mr *MockClusterScopeClientMockRecorder) DeleteRouteTables(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRouteTables", reflect.TypeOf((*MockClusterScopeClient)(nil).DeleteRouteTables), ctx)
}

// DeleteSecurityLists mocks base method.
func (m *MockClusterScopeClient) DeleteSecurityLists(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecurityLists", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecurityLists indicates an expected call of DeleteSecurityLists.
func (mr *MockClusterScopeClientMockRecorder) DeleteSecurityLists(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecurityLists", reflect.TypeOf((*MockClusterScopeClient)(nil).DeleteSecurityLists), ctx)
}

// DeleteServiceGateway mocks base method.
func (m *MockClusterScopeClient) DeleteServiceGateway(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceGateway", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServiceGateway indicates an expected call of DeleteServiceGateway.
func (mr *MockClusterScopeClientMockRecorder) DeleteServiceGateway(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceGateway", reflect.TypeOf((*MockClusterScopeClient)(nil).DeleteServiceGateway), ctx)
}

// DeleteSubnets mocks base method.
func (m *MockClusterScopeClient) DeleteSubnets(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubnets", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubnets indicates an expected call of DeleteSubnets.
func (mr *MockClusterScopeClientMockRecorder) DeleteSubnets(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubnets", reflect.TypeOf((*MockClusterScopeClient)(nil).DeleteSubnets), ctx)
}

// DeleteVCN mocks base method.
func (m *MockClusterScopeClient) DeleteVCN(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVCN", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVCN indicates an expected call of DeleteVCN.
func (mr *MockClusterScopeClientMockRecorder) DeleteVCN(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVCN", reflect.TypeOf((*MockClusterScopeClient)(nil).DeleteVCN), ctx)
}

// GetOCICluster mocks base method.
func (m *MockClusterScopeClient) GetOCICluster() *v1beta1.OCICluster {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOCICluster")
	ret0, _ := ret[0].(*v1beta1.OCICluster)
	return ret0
}

// GetOCICluster indicates an expected call of GetOCICluster.
func (mr *MockClusterScopeClientMockRecorder) GetOCICluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOCICluster", reflect.TypeOf((*MockClusterScopeClient)(nil).GetOCICluster))
}

// ReconcileApiServerLB mocks base method.
func (m *MockClusterScopeClient) ReconcileApiServerLB(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileApiServerLB", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileApiServerLB indicates an expected call of ReconcileApiServerLB.
func (mr *MockClusterScopeClientMockRecorder) ReconcileApiServerLB(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileApiServerLB", reflect.TypeOf((*MockClusterScopeClient)(nil).ReconcileApiServerLB), ctx)
}

// ReconcileFailureDomains mocks base method.
func (m *MockClusterScopeClient) ReconcileFailureDomains(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFailureDomains", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFailureDomains indicates an expected call of ReconcileFailureDomains.
func (mr *MockClusterScopeClientMockRecorder) ReconcileFailureDomains(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFailureDomains", reflect.TypeOf((*MockClusterScopeClient)(nil).ReconcileFailureDomains), ctx)
}

// ReconcileInternetGateway mocks base method.
func (m *MockClusterScopeClient) ReconcileInternetGateway(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileInternetGateway", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileInternetGateway indicates an expected call of ReconcileInternetGateway.
func (mr *MockClusterScopeClientMockRecorder) ReconcileInternetGateway(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileInternetGateway", reflect.TypeOf((*MockClusterScopeClient)(nil).ReconcileInternetGateway), ctx)
}

// ReconcileNSG mocks base method.
func (m *MockClusterScopeClient) ReconcileNSG(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileNSG", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileNSG indicates an expected call of ReconcileNSG.
func (mr *MockClusterScopeClientMockRecorder) ReconcileNSG(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNSG", reflect.TypeOf((*MockClusterScopeClient)(nil).ReconcileNSG), ctx)
}

// ReconcileNatGateway mocks base method.
func (m *MockClusterScopeClient) ReconcileNatGateway(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileNatGateway", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileNatGateway indicates an expected call of ReconcileNatGateway.
func (mr *MockClusterScopeClientMockRecorder) ReconcileNatGateway(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNatGateway", reflect.TypeOf((*MockClusterScopeClient)(nil).ReconcileNatGateway), ctx)
}

// ReconcileRouteTable mocks base method.
func (m *MockClusterScopeClient) ReconcileRouteTable(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRouteTable", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRouteTable indicates an expected call of ReconcileRouteTable.
func (mr *MockClusterScopeClientMockRecorder) ReconcileRouteTable(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRouteTable", reflect.TypeOf((*MockClusterScopeClient)(nil).ReconcileRouteTable), ctx)
}

// ReconcileServiceGateway mocks base method.
func (m *MockClusterScopeClient) ReconcileServiceGateway(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileServiceGateway", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileServiceGateway indicates an expected call of ReconcileServiceGateway.
func (mr *MockClusterScopeClientMockRecorder) ReconcileServiceGateway(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileServiceGateway", reflect.TypeOf((*MockClusterScopeClient)(nil).ReconcileServiceGateway), ctx)
}

// ReconcileSubnet mocks base method.
func (m *MockClusterScopeClient) ReconcileSubnet(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSubnet", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileSubnet indicates an expected call of ReconcileSubnet.
func (mr *MockClusterScopeClientMockRecorder) ReconcileSubnet(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSubnet", reflect.TypeOf((*MockClusterScopeClient)(nil).ReconcileSubnet), ctx)
}

// ReconcileVCN mocks base method.
func (m *MockClusterScopeClient) ReconcileVCN(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVCN", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVCN indicates an expected call of ReconcileVCN.
func (mr *MockClusterScopeClientMockRecorder) ReconcileVCN(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVCN", reflect.TypeOf((*MockClusterScopeClient)(nil).ReconcileVCN), ctx)
}
