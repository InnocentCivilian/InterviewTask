package dto_test

import (
	"testing"
	//
	dto "github.com/innocentcivilian/interviewtask/dto/createdevicerequest"
	"github.com/innocentcivilian/interviewtask/util"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	RequiredEmptyValidationFails(t)
	RequiredMissingValidationFails(t)
	RequiredPresentsValidationPass(t)
}
func RequiredEmptyValidationFails(t *testing.T) {
	msg, err := util.Validate(dto.CreateDeviceRequest{})

	assert.NotEqual(t, msg, "")
	assert.NotNil(t, err)
}
func RequiredMissingValidationFails(t *testing.T) {
	msg, err := util.Validate(dto.CreateDeviceRequest{Id: "id1"})

	assert.NotEqual(t, msg, "")
	assert.NotNil(t, err)
}
func RequiredPresentsValidationPass(t *testing.T) {
	msg, err := util.Validate(dto.CreateDeviceRequest{Id: "id1", Name: "devicename", DeviceModel: "devicemodel", Note: "some note", Serial: "sn1234"})

	assert.Equal(t, msg, "")
	assert.Nil(t, err)
}
