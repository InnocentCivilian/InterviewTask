package service

import (
	"context"

	infrastructure "github.com/innocentcivilian/interviewtask/infrastructure/datastore"
	"github.com/innocentcivilian/interviewtask/model"
	"go.uber.org/zap"
)

type Device = model.Device

// DeviceService is the top level signature of this service
type DeviceService interface {
	Get(ctx context.Context, id string) (*Device, error)
	Create(ctx context.Context, device *Device) error
}

func Init() (DeviceService, error) {

	ddb := infrastructure.Database()
	repository := NewDynamoDBRepository(ddb, "Devices")
	logger, _ := zap.NewProduction()
	usecase := &LoggerAdapter{
		Logger:  logger,
		Usecase: &Usecase{Repository: repository},
	}
	return usecase, nil

}
