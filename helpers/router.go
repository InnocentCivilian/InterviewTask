package helpers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type Request events.APIGatewayProxyRequest

type Response events.APIGatewayProxyResponse
type handler interface {
	Get(ctx context.Context, id string) (Response, error)
	Create(ctx context.Context, body []byte) (Response, error)
}

const fiveSecondsTimeout = time.Second * 5

// Router takes a lambda handler and returns a higher order function
// which can route each request using the HTTP verb and path params.
func Router(handler handler) func(context.Context, Request) (Response, error) {
	return func(ctx context.Context, req Request) (Response, error) {

		// Add cancellation deadline to context
		ctx, cancel := context.WithTimeout(ctx, fiveSecondsTimeout)
		defer cancel()
		switch req.HTTPMethod {
		case http.MethodGet:
			id, _ := req.PathParameters["deviceId"]
			id = "/devices/" + id
			return handler.Get(ctx, id)
		case http.MethodPost:
			return handler.Create(ctx, []byte(req.Body))
		default:
			return Response{}, errors.New("invalid method")
		}
	}
}
