// Code generated by mockery v2.13.1. DO NOT EDIT.

package archivemocks

import (
	io "io"

	internal "github.com/wal-g/wal-g/internal"

	mock "github.com/stretchr/testify/mock"

	models "github.com/wal-g/wal-g/internal/databases/mongo/models"
)

// Uploader is an autogenerated mock type for the Uploader type
type Uploader struct {
	mock.Mock
}

// UploadBackup provides a mock function with given fields: stream, cmd, metaConstructor
func (_m *Uploader) UploadBackup(stream io.Reader, cmd internal.ErrWaiter, metaConstructor internal.MetaConstructor) error {
	ret := _m.Called(stream, cmd, metaConstructor)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Reader, internal.ErrWaiter, internal.MetaConstructor) error); ok {
		r0 = rf(stream, cmd, metaConstructor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadGapArchive provides a mock function with given fields: err, firstTS, lastTS
func (_m *Uploader) UploadGapArchive(err error, firstTS models.Timestamp, lastTS models.Timestamp) error {
	ret := _m.Called(err, firstTS, lastTS)

	var r0 error
	if rf, ok := ret.Get(0).(func(error, models.Timestamp, models.Timestamp) error); ok {
		r0 = rf(err, firstTS, lastTS)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadOplogArchive provides a mock function with given fields: stream, firstTS, lastTS
func (_m *Uploader) UploadOplogArchive(stream io.Reader, firstTS models.Timestamp, lastTS models.Timestamp) error {
	ret := _m.Called(stream, firstTS, lastTS)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Reader, models.Timestamp, models.Timestamp) error); ok {
		r0 = rf(stream, firstTS, lastTS)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUploader interface {
	mock.TestingT
	Cleanup(func())
}

// NewUploader creates a new instance of Uploader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUploader(t mockConstructorTestingTNewUploader) *Uploader {
	mock := &Uploader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
