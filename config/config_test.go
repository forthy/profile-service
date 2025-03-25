package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	O "github.com/IBM/fp-go/option"
)

func TestReadReservedPort(t *testing.T) {
	err1 := os.Setenv("TCP_PORT", "1023")
	if err1 == nil {
		portOpt := ReadReservedPort("TCP_PORT")
		assert.True(t, O.IsSome(portOpt))
	} else {
		assert.Fail(t, "Failed to set environment variable")
	}

	err2 := os.Setenv("UDP_PORT", "8080")
	if err2 == nil {
		portOpt := ReadReservedPort("UDP_PORT")
		assert.True(t, O.IsNone(portOpt))
	} else {
		assert.Fail(t, "Failed to set environment variable")
	}
}
