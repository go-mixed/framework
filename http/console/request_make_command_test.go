package console

import (
	"testing"

	consolemocks "gopkg.in/go-mixed/framework/contracts/console/mocks"
	"gopkg.in/go-mixed/framework/support/file"

	"github.com/stretchr/testify/assert"
)

func TestRequestMakeCommand(t *testing.T) {
	requestMakeCommand := &RequestMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := requestMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("CreateUser").Once()
	err = requestMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("app/http/requests/create_user.go"))
	assert.True(t, file.Remove("app"))
}
