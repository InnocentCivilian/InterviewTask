package userservice

import (
	"context"

	"github.com/innocentcivilian/interviewtask/model"
)

type Device = model.Device

// DeviceService is the top level signature of this service
type DeviceService interface {
	Get(ctx context.Context, id string) (*Device, error)
	Create(ctx context.Context, device *Device) error
}
