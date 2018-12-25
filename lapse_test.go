package lapse

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Lapse(t *testing.T) {
	r := require.New(t)
	d := Lapse(func() {
		time.Sleep(5 * time.Millisecond)
	})
}
