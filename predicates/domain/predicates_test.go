package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainComplaint(t *testing.T) {
	assert.False(t, DomainCompliant("IEA"))
	assert.True(t, DomainCompliant("EA"))
	assert.True(t, DomainCompliant("MFG"))
	assert.True(t, DomainCompliant("ENG"))
}
