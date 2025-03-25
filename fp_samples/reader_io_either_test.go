package fp_samples

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	E "github.com/IBM/fp-go/either"
)

func TestFetchData(t *testing.T) {
	ctx := context.Background()
	E.Fold(
		func(e error) bool {
			return assert.Fail(t, "Should have no error:", e)
		},
		func(v int) bool {
			return assert.Equal(t, 42, v)
		},
	)(fetchData()(ctx)())
}
