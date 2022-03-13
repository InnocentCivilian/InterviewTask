package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	dto "github.com/innocentcivilian/interviewtask/dto/createdevicerequest"
	"github.com/innocentcivilian/interviewtask/helpers"
	"github.com/innocentcivilian/interviewtask/model"
	"github.com/innocentcivilian/interviewtask/service"
	"github.com/innocentcivilian/interviewtask/util"
)

// var usecase service.DeviceService

// type book struct {
// 	ID    int    `dynamodbav:"id"`
// 	Title string `dynamodbav:"title"`
// }
type handler struct {
	usecase service.DeviceService
}

// Get a single device
func (h *handler) Get(ctx context.Context, id string) (helpers.Response, error) {
	device, err := h.usecase.Get(ctx, id)
	if err != nil {
		if strings.HasSuffix(err.Error(), util.NotFound) {
			return util.ResponseMessage(util.NotFound, http.StatusNotFound)
		}
		return util.ResponseMessage(util.InternalError, http.StatusInternalServerError)
	}
	return util.ResponseData(util.OK, device, http.StatusOK)
}

// Create a device
func (h *handler) Create(ctx context.Context, body []byte) (helpers.Response, error) {
	deviceRequest := &dto.CreateDeviceRequest{}
	if err := json.Unmarshal(body, &deviceRequest); err != nil {
		return util.ResponseMessage(util.InternalError, http.StatusInternalServerError)
	}
	// validate attempt
	errMsg, errValidation := util.Validate(deviceRequest)

	if errValidation != nil {
		//validation failed
		return util.ResponseMessage(errMsg, http.StatusBadRequest)
	}
	device := &model.Device{
		Id:          deviceRequest.Id,
		DeviceModel: deviceRequest.DeviceModel,
		Name:        deviceRequest.Name,
		Note:        deviceRequest.Note,
		Serial:      deviceRequest.Serial,
	}

	if err := h.usecase.Create(ctx, device); err != nil {
		return util.ResponseMessage(util.InternalError, http.StatusInternalServerError)
	}
	return util.ResponseData(util.OK, device, http.StatusCreated)
}
func main() {
	usecase, err := service.Init()
	if err != nil {
		log.Panic(err)
	}
	h := &handler{usecase}
	lambda.Start(helpers.Router(h))
}
