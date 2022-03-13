package userservice

import (
	"context"

	"github.com/pkg/errors"
)

type repository interface {
	Get(ctx context.Context, id string) (*Device, error)
	Create(ctx context.Context, device *Device) error
}

// Usecase for interacting with devices
type Usecase struct {
	Repository repository
}

// Get a single device
func (u *Usecase) Get(ctx context.Context, id string) (*Device, error) {
	user, err := u.Repository.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching a single device")
	}
	return user, nil
}

// Create a single device
func (u *Usecase) Create(ctx context.Context, device *Device) error {

	if err := u.Repository.Create(ctx, device); err != nil {
		return errors.Wrap(err, "error creating new device")
	}

	return nil
}
