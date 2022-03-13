package service

import (
	"context"

	"go.uber.org/zap"
)

// LoggerAdapter wraps the usecase interface
// with a logging adapter which can be swapped out
type LoggerAdapter struct {
	Logger  *zap.Logger
	Usecase DeviceService
}

func (a *LoggerAdapter) logErr(err error) {
	if err != nil {
		a.Logger.Error(err.Error())
	}
}

// Get a single user
func (a *LoggerAdapter) Get(ctx context.Context, id string) (*Device, error) {
	defer a.Logger.Sync()
	a.Logger.With(zap.String("id", id))
	a.Logger.Info("getting a single device")
	user, err := a.Usecase.Get(ctx, id)
	a.logErr(err)
	return user, err
}

// Create a single user
func (a *LoggerAdapter) Create(ctx context.Context, device *Device) error {
	defer a.Logger.Sync()
	a.Logger.Info("creating a single device")
	err := a.Usecase.Create(ctx, device)
	a.logErr(err)
	return err
}
