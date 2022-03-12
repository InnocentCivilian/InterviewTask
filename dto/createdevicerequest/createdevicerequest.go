package dto

//data transfer object with validation rule under https://github.com/go-playground/validator rule definition
type CreateDeviceRequest struct {
	Id          string `json:"id" validate:"required"`
	DeviceModel string `json:"deviceModel" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Note        string `json:"note" validate:"required"`
	Serial      string `json:"serial" validate:"required"`
}
