package fp_samples

import (
	IO "github.com/IBM/fp-go/io"
)

func GetCurrentTime() IO.IO[string] {
	return IO.Of("2025-03-13 12:00:00")
}
