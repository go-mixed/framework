package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "gopkg.in/go-mixed/framework/contracts/console/mocks"
	"gopkg.in/go-mixed/framework/support/file"
)

func TestModelMakeCommand(t *testing.T) {
	modelMakeCommand := &ModelMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := modelMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("User").Once()
	err = modelMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("app/models/user.go"))
	assert.True(t, file.Remove("app"))
}
