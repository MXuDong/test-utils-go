package base

import (
	"testing"
	"time"
)

func TestFreezeWithTime(t *testing.T) {
	FreezeTime()
	a := time.Now()
	time.Sleep(1*time.Second)
	b := time.Now()

	if !a.Equal(b) {
		t.Errorf("time is not freeze ")
	}
}
