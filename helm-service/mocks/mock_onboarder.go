// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/keptn/keptn/helm-service/controller (interfaces: Onboarder)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	event "github.com/cloudevents/sdk-go/v2/event"
	gomock "github.com/golang/mock/gomock"
	keptn "github.com/keptn/go-utils/pkg/lib"
	v0_2_0 "github.com/keptn/go-utils/pkg/lib/v0_2_0"
	chart "helm.sh/helm/v3/pkg/chart"
)

// MockOnboarder is a mock of Onboarder interface.
type MockOnboarder struct {
	ctrl     *gomock.Controller
	recorder *MockOnboarderMockRecorder
}

// MockOnboarderMockRecorder is the mock recorder for MockOnboarder.
type MockOnboarderMockRecorder struct {
	mock *MockOnboarder
}

// NewMockOnboarder creates a new mock instance.
func NewMockOnboarder(ctrl *gomock.Controller) *MockOnboarder {
	mock := &MockOnboarder{ctrl: ctrl}
	mock.recorder = &MockOnboarderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOnboarder) EXPECT() *MockOnboarderMockRecorder {
	return m.recorder
}

// HandleEvent mocks base method.
func (m *MockOnboarder) HandleEvent(arg0 event.Event, arg1 func(*v0_2_0.Keptn)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleEvent", arg0, arg1)
}

// HandleEvent indicates an expected call of HandleEvent.
func (mr *MockOnboarderMockRecorder) HandleEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvent", reflect.TypeOf((*MockOnboarder)(nil).HandleEvent), arg0, arg1)
}

// OnboardGeneratedChart mocks base method.
func (m *MockOnboarder) OnboardGeneratedChart(arg0 string, arg1 v0_2_0.EventData, arg2 keptn.DeploymentStrategy) (*chart.Chart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnboardGeneratedChart", arg0, arg1, arg2)
	ret0, _ := ret[0].(*chart.Chart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnboardGeneratedChart indicates an expected call of OnboardGeneratedChart.
func (mr *MockOnboarderMockRecorder) OnboardGeneratedChart(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnboardGeneratedChart", reflect.TypeOf((*MockOnboarder)(nil).OnboardGeneratedChart), arg0, arg1, arg2)
}
