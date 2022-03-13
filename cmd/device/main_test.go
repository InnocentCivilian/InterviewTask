package main

import (
	"testing"
	//
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	// _ := []struct {
	// 	request events.APIGatewayProxyRequest
	// 	expect  string
	// 	err     error
	// }{
	// 	{
	// 		// Test that the handler responds with the correct response
	// 		// when a valid name is provided in the HTTP body
	// 		request: events.APIGatewayProxyRequest{},
	// 		expect:  `{"message":"hello world"}`,
	// 		err:     nil,
	// 	},
	// }

	assert.Equal(t, "a", "ab")
}
