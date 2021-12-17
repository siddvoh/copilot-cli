// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/cli/prompter.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	prompt "github.com/aws/copilot-cli/internal/pkg/term/prompt"
	gomock "github.com/golang/mock/gomock"
)

// Mockprompter is a mock of prompter interface.
type Mockprompter struct {
	ctrl     *gomock.Controller
	recorder *MockprompterMockRecorder
}

// MockprompterMockRecorder is the mock recorder for Mockprompter.
type MockprompterMockRecorder struct {
	mock *Mockprompter
}

// NewMockprompter creates a new mock instance.
func NewMockprompter(ctrl *gomock.Controller) *Mockprompter {
	mock := &Mockprompter{ctrl: ctrl}
	mock.recorder = &MockprompterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockprompter) EXPECT() *MockprompterMockRecorder {
	return m.recorder
}

// Confirm mocks base method.
func (m *Mockprompter) Confirm(message, help string, promptOpts ...prompt.PromptConfig) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{message, help}
	for _, a := range promptOpts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Confirm", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Confirm indicates an expected call of Confirm.
func (mr *MockprompterMockRecorder) Confirm(message, help interface{}, promptOpts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{message, help}, promptOpts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Confirm", reflect.TypeOf((*Mockprompter)(nil).Confirm), varargs...)
}

// Get mocks base method.
func (m *Mockprompter) Get(message, help string, validator prompt.ValidatorFunc, promptOpts ...prompt.PromptConfig) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{message, help, validator}
	for _, a := range promptOpts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockprompterMockRecorder) Get(message, help, validator interface{}, promptOpts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{message, help, validator}, promptOpts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*Mockprompter)(nil).Get), varargs...)
}

// GetSecret mocks base method.
func (m *Mockprompter) GetSecret(message, help string, promptOpts ...prompt.PromptConfig) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{message, help}
	for _, a := range promptOpts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSecret", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockprompterMockRecorder) GetSecret(message, help interface{}, promptOpts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{message, help}, promptOpts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*Mockprompter)(nil).GetSecret), varargs...)
}

// MultiSelect mocks base method.
func (m *Mockprompter) MultiSelect(message, help string, options []string, validator prompt.ValidatorFunc, promptOpts ...prompt.PromptConfig) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{message, help, options, validator}
	for _, a := range promptOpts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiSelect", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiSelect indicates an expected call of MultiSelect.
func (mr *MockprompterMockRecorder) MultiSelect(message, help, options, validator interface{}, promptOpts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{message, help, options, validator}, promptOpts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiSelect", reflect.TypeOf((*Mockprompter)(nil).MultiSelect), varargs...)
}

// SelectOne mocks base method.
func (m *Mockprompter) SelectOne(message, help string, options []string, promptOpts ...prompt.PromptConfig) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{message, help, options}
	for _, a := range promptOpts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectOne", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectOne indicates an expected call of SelectOne.
func (mr *MockprompterMockRecorder) SelectOne(message, help, options interface{}, promptOpts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{message, help, options}, promptOpts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectOne", reflect.TypeOf((*Mockprompter)(nil).SelectOne), varargs...)
}

// SelectOption mocks base method.
func (m *Mockprompter) SelectOption(message, help string, opts []prompt.Option, promptCfgs ...prompt.PromptConfig) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{message, help, opts}
	for _, a := range promptCfgs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectOption", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectOption indicates an expected call of SelectOption.
func (mr *MockprompterMockRecorder) SelectOption(message, help, opts interface{}, promptCfgs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{message, help, opts}, promptCfgs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectOption", reflect.TypeOf((*Mockprompter)(nil).SelectOption), varargs...)
}
