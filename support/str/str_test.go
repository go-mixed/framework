package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	assert.Len(t, Random(10), 10)
}

func TestCase2Camel(t *testing.T) {
	assert.Equal(t, "LaravelFramework", Case2Camel("laravel_framework"))
	assert.Equal(t, "LaravelFramework1", Case2Camel("laravel_framework1"))
}

func TestCamel2Case(t *testing.T) {
	assert.Equal(t, "laravel_framework", Camel2Case("LaravelFramework"))
	assert.Equal(t, "laravel_framework1", Camel2Case("LaravelFramework1"))
}
