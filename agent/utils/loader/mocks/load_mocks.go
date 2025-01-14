// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
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
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/utils/loader (interfaces: Loader)

// Package mock_loader is a generated GoMock package.
package mock_loader

import (
	context "context"
	reflect "reflect"

	config "github.com/aws/amazon-ecs-agent/agent/config"
	dockerapi "github.com/aws/amazon-ecs-agent/agent/dockerclient/dockerapi"
	types "github.com/docker/docker/api/types"
	gomock "github.com/golang/mock/gomock"
)

// MockLoader is a mock of Loader interface
type MockLoader struct {
	ctrl     *gomock.Controller
	recorder *MockLoaderMockRecorder
}

// MockLoaderMockRecorder is the mock recorder for MockLoader
type MockLoaderMockRecorder struct {
	mock *MockLoader
}

// NewMockLoader creates a new mock instance
func NewMockLoader(ctrl *gomock.Controller) *MockLoader {
	mock := &MockLoader{ctrl: ctrl}
	mock.recorder = &MockLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLoader) EXPECT() *MockLoaderMockRecorder {
	return m.recorder
}

// IsLoaded mocks base method
func (m *MockLoader) IsLoaded(arg0 dockerapi.DockerClient) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLoaded", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLoaded indicates an expected call of IsLoaded
func (mr *MockLoaderMockRecorder) IsLoaded(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLoaded", reflect.TypeOf((*MockLoader)(nil).IsLoaded), arg0)
}

// LoadImage mocks base method
func (m *MockLoader) LoadImage(arg0 context.Context, arg1 *config.Config, arg2 dockerapi.DockerClient) (*types.ImageInspect, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.ImageInspect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadImage indicates an expected call of LoadImage
func (mr *MockLoaderMockRecorder) LoadImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadImage", reflect.TypeOf((*MockLoader)(nil).LoadImage), arg0, arg1, arg2)
}
