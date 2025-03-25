package fp_samples

import (
	RIOE "github.com/IBM/fp-go/context/readerioeither"
)

type Result = RIOE.ReaderIOEither[int]

func fetchData() Result {
	return RIOE.Of(42)
}
