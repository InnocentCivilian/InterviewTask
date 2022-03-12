package util

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	dto "github.com/innocentcivilian/interviewtask/dto/createdevicerequest"
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
func ResponseMessage(message string, statusCode int) (events.APIGatewayProxyResponse, error) {
	var resp = dto.ResponseTemplate{
		Message: message,
	}
	data, err := jsonSerialze(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       data,
			StatusCode: 500,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       data,
		StatusCode: statusCode,
	}, nil
}
func ResponseData(message string, data interface{}, statusCode int) (events.APIGatewayProxyResponse, error) {

	jsonSerialzed, err := jsonSerialze(data)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       jsonSerialzed,
			StatusCode: 500,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       jsonSerialzed,
		StatusCode: statusCode,
	}, nil
}
