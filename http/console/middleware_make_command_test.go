package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "gopkg.in/go-mixed/framework/contracts/console/mocks"
	"gopkg.in/go-mixed/framework/support/file"
)

func TestMiddlewareMakeCommand(t *testing.T) {
	middlewareMakeCommand := &MiddlewareMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := middlewareMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("VerifyCsrfToken").Once()
	err = middlewareMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("app/http/middleware/verify_csrf_token.go"))
	assert.True(t, file.Remove("app"))
}
