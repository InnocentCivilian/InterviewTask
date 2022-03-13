package util

import (
	"encoding/json"
	"errors"

	"github.com/innocentcivilian/interviewtask/dto"
	"github.com/innocentcivilian/interviewtask/helpers"
)

const InternalError string = "Internal Error"
const OK string = "OK"
const CREATED string = "Created"
const NotFound string = "not found"
const InvalidJsonError string = "Invalid/bad json format"
const InputDataInvalid string = "Input data invalid"

type ResponseTemplate = dto.ResponseTemplate

func internalErrorBody() string {
	var resp = ResponseTemplate{
		Message: InternalError,
	}
	data, _ := json.MarshalIndent(resp, "", "    ")
	return string(data)
}
func jsonSerialze(data interface{}) (string, error) {
	jsonSerialzed, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return internalErrorBody(), errors.New("failed to serilze json")
	}
	return string(jsonSerialzed), nil
}
func ResponseMessage(message string, statusCode int) (helpers.Response, error) {
	var resp = ResponseTemplate{
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
