package tests

import (
	"testing"
)

func FuzzQueueRun(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int) {
		//assert := assert.New(t)

	})
}
