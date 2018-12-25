package lapse

import (
	"fmt"
	"time"

	"github.com/markbates/safe"
)

func Lapse(fn func()) time.Duration {
	now := time.Now()
	safe.Run(fn)
	return time.Now().Sub(now)
}

func LapseP(msg string, fn func()) time.Duration {
	d := Lapse(fn)
	fmt.Println(msg, d)
	return d
}
