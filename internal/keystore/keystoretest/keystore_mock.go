// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: keystore.go

// Package keystoretest is a generated GoMock package.
package keystoretest

import (
	gomock "github.com/golang/mock/gomock"
	crypto "github.com/mailchain/mailchain/crypto"
	cipher "github.com/mailchain/mailchain/crypto/cipher"
	multi "github.com/mailchain/mailchain/internal/keystore/kdf/multi"
	mailbox "github.com/mailchain/mailchain/internal/mailbox"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetSigner mocks base method
func (m *MockStore) GetSigner(address []byte, chain string, deriveKeyOptions multi.OptionsBuilders) (mailbox.Signer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSigner", address, chain, deriveKeyOptions)
	ret0, _ := ret[0].(mailbox.Signer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSigner indicates an expected call of GetSigner
func (mr *MockStoreMockRecorder) GetSigner(address, chain, deriveKeyOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSigner", reflect.TypeOf((*MockStore)(nil).GetSigner), address, chain, deriveKeyOptions)
}

// GetDecrypter mocks base method
func (m *MockStore) GetDecrypter(address []byte, decrypterType byte, deriveKeyOptions multi.OptionsBuilders) (cipher.Decrypter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDecrypter", address, decrypterType, deriveKeyOptions)
	ret0, _ := ret[0].(cipher.Decrypter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDecrypter indicates an expected call of GetDecrypter
func (mr *MockStoreMockRecorder) GetDecrypter(address, decrypterType, deriveKeyOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDecrypter", reflect.TypeOf((*MockStore)(nil).GetDecrypter), address, decrypterType, deriveKeyOptions)
}

// Store mocks base method
func (m *MockStore) Store(private crypto.PrivateKey, curveType string, deriveKeyOptions multi.OptionsBuilders) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", private, curveType, deriveKeyOptions)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockStoreMockRecorder) Store(private, curveType, deriveKeyOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockStore)(nil).Store), private, curveType, deriveKeyOptions)
}

// HasAddress mocks base method
func (m *MockStore) HasAddress(address []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAddress", address)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasAddress indicates an expected call of HasAddress
func (mr *MockStoreMockRecorder) HasAddress(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAddress", reflect.TypeOf((*MockStore)(nil).HasAddress), address)
}

// GetAddresses mocks base method
func (m *MockStore) GetAddresses() ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddresses")
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddresses indicates an expected call of GetAddresses
func (mr *MockStoreMockRecorder) GetAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddresses", reflect.TypeOf((*MockStore)(nil).GetAddresses))
}
