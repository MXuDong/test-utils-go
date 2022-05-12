package base

import (
	"testing"
	"time"
)

func TestFreezeWithTime(t *testing.T) {
	FreezeTime()
	a := time.Now()
	time.Sleep(1 * time.Second)
	b := time.Now()

	if !a.Equal(b) {
		t.Errorf("time is not freeze ")
	}
	UnFreezeTime()
	c := time.Now()
	if a.Equal(c){
		t.Errorf("unfreeze error")
	}
}

func TestDurationFreeze(t *testing.T) {
	FreezeTime()
	a := time.Now()
	AddDuration(1 * time.Minute)
	b := time.Now()

	if b.Sub(a) != 1*time.Minute {
		t.Errorf("time duration is error, e: %v, but: %v", 1*time.Minute, b.Sub(a))
	}
	AddDuration(1 * time.Minute)
	c := time.Now()
	if c.Sub(a) != 2*time.Minute{
		t.Errorf("time duration is error, e: %v, but: %v", 2*time.Minute, c.Sub(a))
	}
	CleanDuration()
	d := time.Now()
	if d.Sub(a) != 0 {
		t.Errorf("Clean duration is error")
	}
}
