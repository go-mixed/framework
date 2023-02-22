package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "gopkg.in/go-mixed/framework/contracts/console/mocks"
	"gopkg.in/go-mixed/framework/support/file"
)

func TestControllerMakeCommand(t *testing.T) {
	controllerMakeCommand := &ControllerMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := controllerMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("UsersController").Once()
	err = controllerMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("app/http/controllers/users_controller.go"))
	assert.True(t, file.Remove("app"))
}
