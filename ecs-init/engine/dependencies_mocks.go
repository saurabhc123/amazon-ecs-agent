// Copyright 2015-2023 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
// Source: dependencies.go in package engine
// Code generated by MockGen. DO NOT EDIT.

// Package engine is a generated GoMock package.
package engine

import (
	io "io"
	reflect "reflect"

	cache "github.com/aws/amazon-ecs-agent/ecs-init/cache"
	gomock "github.com/golang/mock/gomock"
)

// Mockdownloader is a mock of downloader interface
type Mockdownloader struct {
	ctrl     *gomock.Controller
	recorder *MockdownloaderMockRecorder
}

// MockdownloaderMockRecorder is the mock recorder for Mockdownloader
type MockdownloaderMockRecorder struct {
	mock *Mockdownloader
}

// NewMockdownloader creates a new mock instance
func NewMockdownloader(ctrl *gomock.Controller) *Mockdownloader {
	mock := &Mockdownloader{ctrl: ctrl}
	mock.recorder = &MockdownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockdownloader) EXPECT() *MockdownloaderMockRecorder {
	return m.recorder
}

// IsAgentCached mocks base method
func (m *Mockdownloader) IsAgentCached() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAgentCached")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAgentCached indicates an expected call of IsAgentCached
func (mr *MockdownloaderMockRecorder) IsAgentCached() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAgentCached", reflect.TypeOf((*Mockdownloader)(nil).IsAgentCached))
}

// DownloadAgent mocks base method
func (m *Mockdownloader) DownloadAgent() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadAgent")
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadAgent indicates an expected call of DownloadAgent
func (mr *MockdownloaderMockRecorder) DownloadAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadAgent", reflect.TypeOf((*Mockdownloader)(nil).DownloadAgent))
}

// LoadCachedAgent mocks base method
func (m *Mockdownloader) LoadCachedAgent() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadCachedAgent")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadCachedAgent indicates an expected call of LoadCachedAgent
func (mr *MockdownloaderMockRecorder) LoadCachedAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadCachedAgent", reflect.TypeOf((*Mockdownloader)(nil).LoadCachedAgent))
}

// LoadDesiredAgent mocks base method
func (m *Mockdownloader) LoadDesiredAgent() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadDesiredAgent")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadDesiredAgent indicates an expected call of LoadDesiredAgent
func (mr *MockdownloaderMockRecorder) LoadDesiredAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadDesiredAgent", reflect.TypeOf((*Mockdownloader)(nil).LoadDesiredAgent))
}

// RecordCachedAgent mocks base method
func (m *Mockdownloader) RecordCachedAgent() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordCachedAgent")
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordCachedAgent indicates an expected call of RecordCachedAgent
func (mr *MockdownloaderMockRecorder) RecordCachedAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordCachedAgent", reflect.TypeOf((*Mockdownloader)(nil).RecordCachedAgent))
}

// AgentCacheStatus mocks base method
func (m *Mockdownloader) AgentCacheStatus() cache.CacheStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentCacheStatus")
	ret0, _ := ret[0].(cache.CacheStatus)
	return ret0
}

// AgentCacheStatus indicates an expected call of AgentCacheStatus
func (mr *MockdownloaderMockRecorder) AgentCacheStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentCacheStatus", reflect.TypeOf((*Mockdownloader)(nil).AgentCacheStatus))
}

// MockdockerClient is a mock of dockerClient interface
type MockdockerClient struct {
	ctrl     *gomock.Controller
	recorder *MockdockerClientMockRecorder
}

// MockdockerClientMockRecorder is the mock recorder for MockdockerClient
type MockdockerClientMockRecorder struct {
	mock *MockdockerClient
}

// NewMockdockerClient creates a new mock instance
func NewMockdockerClient(ctrl *gomock.Controller) *MockdockerClient {
	mock := &MockdockerClient{ctrl: ctrl}
	mock.recorder = &MockdockerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockdockerClient) EXPECT() *MockdockerClientMockRecorder {
	return m.recorder
}

// GetContainerLogTail mocks base method
func (m *MockdockerClient) GetContainerLogTail(logWindowSize string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerLogTail", logWindowSize)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetContainerLogTail indicates an expected call of GetContainerLogTail
func (mr *MockdockerClientMockRecorder) GetContainerLogTail(logWindowSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerLogTail", reflect.TypeOf((*MockdockerClient)(nil).GetContainerLogTail), logWindowSize)
}

// IsAgentImageLoaded mocks base method
func (m *MockdockerClient) IsAgentImageLoaded() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAgentImageLoaded")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAgentImageLoaded indicates an expected call of IsAgentImageLoaded
func (mr *MockdockerClientMockRecorder) IsAgentImageLoaded() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAgentImageLoaded", reflect.TypeOf((*MockdockerClient)(nil).IsAgentImageLoaded))
}

// LoadImage mocks base method
func (m *MockdockerClient) LoadImage(image io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadImage", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadImage indicates an expected call of LoadImage
func (mr *MockdockerClientMockRecorder) LoadImage(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadImage", reflect.TypeOf((*MockdockerClient)(nil).LoadImage), image)
}

// RemoveExistingAgentContainer mocks base method
func (m *MockdockerClient) RemoveExistingAgentContainer() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveExistingAgentContainer")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveExistingAgentContainer indicates an expected call of RemoveExistingAgentContainer
func (mr *MockdockerClientMockRecorder) RemoveExistingAgentContainer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveExistingAgentContainer", reflect.TypeOf((*MockdockerClient)(nil).RemoveExistingAgentContainer))
}

// StartAgent mocks base method
func (m *MockdockerClient) StartAgent() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartAgent")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartAgent indicates an expected call of StartAgent
func (mr *MockdockerClientMockRecorder) StartAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAgent", reflect.TypeOf((*MockdockerClient)(nil).StartAgent))
}

// StopAgent mocks base method
func (m *MockdockerClient) StopAgent() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopAgent")
	ret0, _ := ret[0].(error)
	return ret0
}

// StopAgent indicates an expected call of StopAgent
func (mr *MockdockerClientMockRecorder) StopAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAgent", reflect.TypeOf((*MockdockerClient)(nil).StopAgent))
}

// LoadEnvVars mocks base method
func (m *MockdockerClient) LoadEnvVars() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEnvVars")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// LoadEnvVars indicates an expected call of LoadEnvVars
func (mr *MockdockerClientMockRecorder) LoadEnvVars() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEnvVars", reflect.TypeOf((*MockdockerClient)(nil).LoadEnvVars))
}

// MockloopbackRouting is a mock of loopbackRouting interface
type MockloopbackRouting struct {
	ctrl     *gomock.Controller
	recorder *MockloopbackRoutingMockRecorder
}

// MockloopbackRoutingMockRecorder is the mock recorder for MockloopbackRouting
type MockloopbackRoutingMockRecorder struct {
	mock *MockloopbackRouting
}

// NewMockloopbackRouting creates a new mock instance
func NewMockloopbackRouting(ctrl *gomock.Controller) *MockloopbackRouting {
	mock := &MockloopbackRouting{ctrl: ctrl}
	mock.recorder = &MockloopbackRoutingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockloopbackRouting) EXPECT() *MockloopbackRoutingMockRecorder {
	return m.recorder
}

// Enable mocks base method
func (m *MockloopbackRouting) Enable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable")
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable
func (mr *MockloopbackRoutingMockRecorder) Enable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockloopbackRouting)(nil).Enable))
}

// RestoreDefault mocks base method
func (m *MockloopbackRouting) RestoreDefault() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreDefault")
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreDefault indicates an expected call of RestoreDefault
func (mr *MockloopbackRoutingMockRecorder) RestoreDefault() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreDefault", reflect.TypeOf((*MockloopbackRouting)(nil).RestoreDefault))
}

// MockcredentialsProxyRoute is a mock of credentialsProxyRoute interface
type MockcredentialsProxyRoute struct {
	ctrl     *gomock.Controller
	recorder *MockcredentialsProxyRouteMockRecorder
}

// MockcredentialsProxyRouteMockRecorder is the mock recorder for MockcredentialsProxyRoute
type MockcredentialsProxyRouteMockRecorder struct {
	mock *MockcredentialsProxyRoute
}

// NewMockcredentialsProxyRoute creates a new mock instance
func NewMockcredentialsProxyRoute(ctrl *gomock.Controller) *MockcredentialsProxyRoute {
	mock := &MockcredentialsProxyRoute{ctrl: ctrl}
	mock.recorder = &MockcredentialsProxyRouteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockcredentialsProxyRoute) EXPECT() *MockcredentialsProxyRouteMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockcredentialsProxyRoute) Create() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockcredentialsProxyRouteMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockcredentialsProxyRoute)(nil).Create))
}

// Remove mocks base method
func (m *MockcredentialsProxyRoute) Remove() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove")
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockcredentialsProxyRouteMockRecorder) Remove() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockcredentialsProxyRoute)(nil).Remove))
}

// Mockipv6RouterAdvertisements is a mock of ipv6RouterAdvertisements interface
type Mockipv6RouterAdvertisements struct {
	ctrl     *gomock.Controller
	recorder *Mockipv6RouterAdvertisementsMockRecorder
}

// Mockipv6RouterAdvertisementsMockRecorder is the mock recorder for Mockipv6RouterAdvertisements
type Mockipv6RouterAdvertisementsMockRecorder struct {
	mock *Mockipv6RouterAdvertisements
}

// NewMockipv6RouterAdvertisements creates a new mock instance
func NewMockipv6RouterAdvertisements(ctrl *gomock.Controller) *Mockipv6RouterAdvertisements {
	mock := &Mockipv6RouterAdvertisements{ctrl: ctrl}
	mock.recorder = &Mockipv6RouterAdvertisementsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockipv6RouterAdvertisements) EXPECT() *Mockipv6RouterAdvertisementsMockRecorder {
	return m.recorder
}

// Disable mocks base method
func (m *Mockipv6RouterAdvertisements) Disable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disable indicates an expected call of Disable
func (mr *Mockipv6RouterAdvertisementsMockRecorder) Disable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*Mockipv6RouterAdvertisements)(nil).Disable))
}
