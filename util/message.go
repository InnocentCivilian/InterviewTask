package util

import (
	"encoding/json"
	"errors"

	dto "github.com/innocentcivilian/interviewtask/dto/createdevicerequest"
	"github.com/innocentcivilian/interviewtask/helpers"
)

const InternalError string = "Internal Error"
const OK string = "OK"
const CREATED string = "Created"
const InvalidJsonError string = "Invalid/bad json format"
const InputDataInvalid string = "Input data invalid"

func internalErrorBody() string {
	var resp = dto.ResponseTemplate{
		Message: InternalError,
	}
	data, _ := json.MarshalIndent(resp, "", "    ")
	return string(data)
}
func jsonSerialze(data interface{}) (string, error) {
	var resp = dto.ResponseTemplate{
		Message: InternalError,
	}
	jsonSerialzed, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		return internalErrorBody(), errors.New("failed to serilze json")
	}
	return string(jsonSerialzed), nil
}
func ResponseMessage(message string, statusCode int) (helpers.Response, error) {
	var resp = dto.ResponseTemplate{
		Message: message,
	}
	data, err := jsonSerialze(resp)
	if err != nil {
		return helpers.Response{
			Body:       data,
			StatusCode: 500,
		}, nil
	}
	return helpers.Response{
		Body:       data,
		StatusCode: statusCode,
	}, nil
}
func ResponseData(message string, data interface{}, statusCode int) (helpers.Response, error) {

	jsonSerialzed, err := jsonSerialze(data)
	if err != nil {
		return helpers.Response{
			Body:       jsonSerialzed,
			StatusCode: 500,
		}, nil
	}
	return helpers.Response{
		Body:       jsonSerialzed,
		StatusCode: statusCode,
	}, nil
}
