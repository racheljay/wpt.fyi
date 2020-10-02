// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/shared (interfaces: AppEngineAPI)

// Package sharedtest is a generated GoMock package.
package sharedtest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v31/github"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
	http "net/http"
	url "net/url"
	reflect "reflect"
	time "time"
)

// MockAppEngineAPI is a mock of AppEngineAPI interface
type MockAppEngineAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAppEngineAPIMockRecorder
}

// MockAppEngineAPIMockRecorder is the mock recorder for MockAppEngineAPI
type MockAppEngineAPIMockRecorder struct {
	mock *MockAppEngineAPI
}

// NewMockAppEngineAPI creates a new mock instance
func NewMockAppEngineAPI(ctrl *gomock.Controller) *MockAppEngineAPI {
	mock := &MockAppEngineAPI{ctrl: ctrl}
	mock.recorder = &MockAppEngineAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppEngineAPI) EXPECT() *MockAppEngineAPIMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockAppEngineAPI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockAppEngineAPIMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAppEngineAPI)(nil).Context))
}

// GetGitHubClient mocks base method
func (m *MockAppEngineAPI) GetGitHubClient() (*github.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGitHubClient")
	ret0, _ := ret[0].(*github.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGitHubClient indicates an expected call of GetGitHubClient
func (mr *MockAppEngineAPIMockRecorder) GetGitHubClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGitHubClient", reflect.TypeOf((*MockAppEngineAPI)(nil).GetGitHubClient))
}

// GetHTTPClient mocks base method
func (m *MockAppEngineAPI) GetHTTPClient() *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPClient")
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// GetHTTPClient indicates an expected call of GetHTTPClient
func (mr *MockAppEngineAPIMockRecorder) GetHTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPClient", reflect.TypeOf((*MockAppEngineAPI)(nil).GetHTTPClient))
}

// GetHTTPClientWithTimeout mocks base method
func (m *MockAppEngineAPI) GetHTTPClientWithTimeout(arg0 time.Duration) *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPClientWithTimeout", arg0)
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// GetHTTPClientWithTimeout indicates an expected call of GetHTTPClientWithTimeout
func (mr *MockAppEngineAPIMockRecorder) GetHTTPClientWithTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPClientWithTimeout", reflect.TypeOf((*MockAppEngineAPI)(nil).GetHTTPClientWithTimeout), arg0)
}

// GetHostname mocks base method
func (m *MockAppEngineAPI) GetHostname() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostname")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHostname indicates an expected call of GetHostname
func (mr *MockAppEngineAPIMockRecorder) GetHostname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostname", reflect.TypeOf((*MockAppEngineAPI)(nil).GetHostname))
}

// GetResultsURL mocks base method
func (m *MockAppEngineAPI) GetResultsURL(arg0 shared.TestRunFilter) *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResultsURL", arg0)
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// GetResultsURL indicates an expected call of GetResultsURL
func (mr *MockAppEngineAPIMockRecorder) GetResultsURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResultsURL", reflect.TypeOf((*MockAppEngineAPI)(nil).GetResultsURL), arg0)
}

// GetResultsUploadURL mocks base method
func (m *MockAppEngineAPI) GetResultsUploadURL() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResultsUploadURL")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// GetResultsUploadURL indicates an expected call of GetResultsUploadURL
func (mr *MockAppEngineAPIMockRecorder) GetResultsUploadURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResultsUploadURL", reflect.TypeOf((*MockAppEngineAPI)(nil).GetResultsUploadURL))
}

// GetRunsURL mocks base method
func (m *MockAppEngineAPI) GetRunsURL(arg0 shared.TestRunFilter) *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunsURL", arg0)
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// GetRunsURL indicates an expected call of GetRunsURL
func (mr *MockAppEngineAPIMockRecorder) GetRunsURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunsURL", reflect.TypeOf((*MockAppEngineAPI)(nil).GetRunsURL), arg0)
}

// GetServiceHostname mocks base method
func (m *MockAppEngineAPI) GetServiceHostname(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceHostname", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceHostname indicates an expected call of GetServiceHostname
func (mr *MockAppEngineAPIMockRecorder) GetServiceHostname(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceHostname", reflect.TypeOf((*MockAppEngineAPI)(nil).GetServiceHostname), arg0)
}

// GetUploader mocks base method
func (m *MockAppEngineAPI) GetUploader(arg0 string) (shared.Uploader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUploader", arg0)
	ret0, _ := ret[0].(shared.Uploader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUploader indicates an expected call of GetUploader
func (mr *MockAppEngineAPIMockRecorder) GetUploader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUploader", reflect.TypeOf((*MockAppEngineAPI)(nil).GetUploader), arg0)
}

// GetVersion mocks base method
func (m *MockAppEngineAPI) GetVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockAppEngineAPIMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockAppEngineAPI)(nil).GetVersion))
}

// GetVersionedHostname mocks base method
func (m *MockAppEngineAPI) GetVersionedHostname() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersionedHostname")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVersionedHostname indicates an expected call of GetVersionedHostname
func (mr *MockAppEngineAPIMockRecorder) GetVersionedHostname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionedHostname", reflect.TypeOf((*MockAppEngineAPI)(nil).GetVersionedHostname))
}

// IsFeatureEnabled mocks base method
func (m *MockAppEngineAPI) IsFeatureEnabled(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFeatureEnabled", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFeatureEnabled indicates an expected call of IsFeatureEnabled
func (mr *MockAppEngineAPIMockRecorder) IsFeatureEnabled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFeatureEnabled", reflect.TypeOf((*MockAppEngineAPI)(nil).IsFeatureEnabled), arg0)
}

// ScheduleTask mocks base method
func (m *MockAppEngineAPI) ScheduleTask(arg0, arg1, arg2 string, arg3 url.Values) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleTask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleTask indicates an expected call of ScheduleTask
func (mr *MockAppEngineAPIMockRecorder) ScheduleTask(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleTask", reflect.TypeOf((*MockAppEngineAPI)(nil).ScheduleTask), arg0, arg1, arg2, arg3)
}
