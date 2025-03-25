package fp_samples

import (
	"testing"
	"time"

	I "github.com/IBM/fp-go/io"
	O "github.com/IBM/fp-go/option"
	R "github.com/IBM/fp-go/retry"
	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	fn := findUserByIdWithCount(2)

	policies := R.Monoid.Concat(R.ConstantDelay(2*time.Second), R.LimitRetries(4))

	x := I.Retrying(
		policies,
		func(rs R.RetryStatus) I.IO[O.Option[User]] {
			return I.Of(fn(UserId{"UT-28474"}))
		},
		func(ou O.Option[User]) bool {
			return O.IsNone(ou)
		},
	)

	assert.True(t, O.IsSome(x()))
}
