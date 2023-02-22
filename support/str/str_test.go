package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	assert.Len(t, Random(10), 10)
}

func TestCase2Camel(t *testing.T) {
	assert.Equal(t, "GoravelFramework", Case2Camel("goravel_framework"))
	assert.Equal(t, "GoravelFramework1", Case2Camel("goravel_framework1"))
}

func TestCamel2Case(t *testing.T) {
	assert.Equal(t, "goravel_framework", Camel2Case("GoravelFramework"))
	assert.Equal(t, "goravel_framework1", Camel2Case("GoravelFramework1"))
}
