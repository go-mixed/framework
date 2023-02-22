package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "gopkg.in/go-mixed/framework.v1/contracts/console/mocks"
	"gopkg.in/go-mixed/framework.v1/support/file"
)

func TestListenerMakeCommand(t *testing.T) {
	listenerMakeCommand := &ListenerMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := listenerMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("GoravelListen").Once()
	err = listenerMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("app/listeners/goravel_listen.go"))
	assert.True(t, file.Remove("app"))
}
