package fp_samples

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//IO "github.com/IBM/fp-go/io"
)

func TestGetCurrentTime(t *testing.T) {
	assert.Equal(t, "2025-03-13 12:00:00", GetCurrentTime()())
}
