package helpers

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelpers(t *testing.T) {

}

type MockHandler struct {
	invokedMethod string
	ctx           context.Context
	id            string
	body          []byte
}

func (m *MockHandler) Get(ctx context.Context, id string) (Response, error) {
	m.invokedMethod = "Get"
	m.ctx = ctx
	m.id = id
	return Response{}, nil
}

func (m *MockHandler) Create(ctx context.Context, body []byte) (Response, error) {
	m.invokedMethod = "Create"
	m.ctx = ctx
	m.body = body
	return Response{}, nil
}

func Router_GivenAnyRequest_ShouldAddCancellationDeadlineToContext(t *testing.T) {
	// Arrange
	var req Request

	var handler MockHandler

	// Act
	Router(&handler)(context.Background(), req)

	// Assert
	_, ok := handler.ctx.Deadline()
	assert.True(t, ok)
}

func Router_GivenValidGetRequest_ShouldInvokeGetHandler(t *testing.T) {

	// Arrange
	var req Request
	req.HTTPMethod = http.MethodGet
	req.PathParameters["deviceId"] = "123"

	var handler MockHandler

	// Act
	_, err := Router(&handler)(context.Background(), req)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "Get", handler.invokedMethod)
	assert.Equal(t, "/devices/123", handler.id)
}

func Router_GivenValidCreateRequest_ShouldInvokeGetHandler(t *testing.T) {
	// Arrange
	var req Request
	req.HTTPMethod = http.MethodPost
	req.Body = "{\"deviceId\":\"123\",\"deviceName\":\"test\"}"

	var handler MockHandler

	// Act
	_, err := Router(&handler)(context.Background(), req)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "Create", handler.invokedMethod)
	assert.Equal(t, "{\"deviceId\":\"123\",\"deviceName\":\"test\"}", string(handler.body))
}

func Router_GivenInvalidRequest_ShouldReturnError(t *testing.T) {
	// Arrange
	var req Request

	var handler MockHandler

	// Act
	_, err := Router(&handler)(context.Background(), req)

	// Assert
	assert.NotNil(t, err)
}
