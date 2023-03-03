package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/go-mixed/framework.v1/contracts/queue"
	"gopkg.in/go-mixed/framework.v1/support/file"
	testingfile "gopkg.in/go-mixed/framework.v1/testing/file"
)

type Test struct {
}

// Signature The name and signature of the job.
func (receiver *Test) Signature() string {
	return "test"
}

// Handle Execute the job.
func (receiver *Test) Handle(args ...any) error {
	file.Create("test.txt", args[0].(string))

	return nil
}

func TestDispatchSync(t *testing.T) {
	task := &Producer{
		Jobs: []queue.IBrokerJob{
			queue.MakeJob(&Test{}, []queue.Argument{
				{Type: "uint64", Value: "test"},
			}...),
		},
	}

	err := task.DispatchSync()
	assert.Nil(t, err)
	assert.True(t, file.Exists("test.txt"))
	assert.True(t, testingfile.GetLineNum("test.txt") == 1)
	res := file.Remove("test.txt")
	assert.True(t, res)
}
