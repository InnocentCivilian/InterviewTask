package service

import (
	"context"
	reflect "reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// MockrepositoryMockRecorder is the mock recorder for Mockrepository
type MockrepositoryMockRecorder struct {
	mock *Mockrepository
}

// Mockrepository is a mock of repository interface
type Mockrepository struct {
	ctrl     *gomock.Controller
	recorder *MockrepositoryMockRecorder
}

// NewMockrepository creates a new mock instance
func NewMockrepository(ctrl *gomock.Controller) *Mockrepository {
	mock := &Mockrepository{ctrl: ctrl}
	mock.recorder = &MockrepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockrepository) EXPECT() *MockrepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *Mockrepository) Get(ctx context.Context, id string) (*Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockrepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*Mockrepository)(nil).Get), ctx, id)
}

// Create mocks base method
func (m *Mockrepository) Create(ctx context.Context, device *Device) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, device)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockrepositoryMockRecorder) Create(ctx, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Mockrepository)(nil).Create), ctx, device)
}
func TestCanGetUser(t *testing.T) {
	expected := &Device{
		Id: "id1",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockrepository(ctrl)
	repo.EXPECT().Get(context.Background(), "abc123").Return(expected, nil)

	uc := Usecase{repo}

	user, err := uc.Get(context.Background(), "abc123")

	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}
func TestCanCreateUser(t *testing.T) {
	expected := &Device{
		Id: "id1",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockrepository(ctrl)
	repo.EXPECT().Create(context.Background(), expected).Return(nil)

	uc := Usecase{repo}
	err := uc.Create(context.Background(), expected)

	assert.NoError(t, err)
}
