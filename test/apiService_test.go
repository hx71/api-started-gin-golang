package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiVersion(t *testing.T) {
	version := "1.0.0"
	assert.NotNil(t, version)
	assert.Equal(t, version, "1.0.0")
	fmt.Println("running test api version")
}
