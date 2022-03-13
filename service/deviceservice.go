package service

import (
	"context"
	"os"
	"time"

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

const fiveSecondsTimeout = time.Second * 5

func Init() (DeviceService, error) {
	tableName := os.Getenv("TABLE_NAME")

	ddb := infrastructure.Database()
	repository := NewDynamoDBRepository(ddb, tableName)
	logger, _ := zap.NewProduction()
	usecase := &LoggerAdapter{
		Logger:  logger,
		Usecase: &Usecase{Repository: repository},
	}
	return usecase, nil

}
