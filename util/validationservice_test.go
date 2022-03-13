// package util
package util_test

import (
	"testing"
	//
	"github.com/innocentcivilian/interviewtask/util"
	"github.com/stretchr/testify/assert"
)

type TestPartial struct {
	Required string `validate:"required"`
}

func TestHandler(t *testing.T) {
	RequiredMissingValidationFails(t)
	RequiredPresentsValidationPass(t)
}
func RequiredMissingValidationFails(t *testing.T) {
	msg, err := util.Validate(TestPartial{})

	assert.NotEqual(t, msg, "")
	assert.NotNil(t, err)
}
func RequiredPresentsValidationPass(t *testing.T) {
	msg, err := util.Validate(TestPartial{Required: "value"})

	assert.Equal(t, msg, "")
	assert.Nil(t, err)
}
