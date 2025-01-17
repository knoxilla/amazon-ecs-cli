// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients/aws/ecr (interfaces: Client)

// Package mock_ecr is a generated GoMock package.
package mock_ecr

import (
	reflect "reflect"

	ecr "github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients/aws/ecr"
	gomock "github.com/golang/mock/gomock"
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

// CreateRepository mocks base method.
func (m *MockClient) CreateRepository(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRepository", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRepository indicates an expected call of CreateRepository.
func (mr *MockClientMockRecorder) CreateRepository(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRepository", reflect.TypeOf((*MockClient)(nil).CreateRepository), arg0)
}

// GetAuthorizationToken mocks base method.
func (m *MockClient) GetAuthorizationToken(arg0 string) (*ecr.Auth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorizationToken", arg0)
	ret0, _ := ret[0].(*ecr.Auth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorizationToken indicates an expected call of GetAuthorizationToken.
func (mr *MockClientMockRecorder) GetAuthorizationToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizationToken", reflect.TypeOf((*MockClient)(nil).GetAuthorizationToken), arg0)
}

// GetAuthorizationTokenByID mocks base method.
func (m *MockClient) GetAuthorizationTokenByID(arg0 string) (*ecr.Auth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorizationTokenByID", arg0)
	ret0, _ := ret[0].(*ecr.Auth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorizationTokenByID indicates an expected call of GetAuthorizationTokenByID.
func (mr *MockClientMockRecorder) GetAuthorizationTokenByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizationTokenByID", reflect.TypeOf((*MockClient)(nil).GetAuthorizationTokenByID), arg0)
}

// GetImages mocks base method.
func (m *MockClient) GetImages(arg0 []*string, arg1, arg2 string, arg3 ecr.ProcessImageDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImages", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetImages indicates an expected call of GetImages.
func (mr *MockClientMockRecorder) GetImages(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImages", reflect.TypeOf((*MockClient)(nil).GetImages), arg0, arg1, arg2, arg3)
}

// RepositoryExists mocks base method.
func (m *MockClient) RepositoryExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RepositoryExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RepositoryExists indicates an expected call of RepositoryExists.
func (mr *MockClientMockRecorder) RepositoryExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepositoryExists", reflect.TypeOf((*MockClient)(nil).RepositoryExists), arg0)
}
