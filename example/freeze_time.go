package example

import (
	"github.com/MXuDong/test-utils-go/base"
	"time"
)

func ExampleFreezeTime() {
	if err := base.FreezeTime(); err != nil {
		panic(err)
	}

	// your unit test code
	println(time.Now())
	time.Sleep(10 * time.Second)
	println(time.Now())

	// the output is same time.

	// restore time flow
	base.UnFreezeTime()
}

func ExampleCustomeFreezeTime() {
	t := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	if err := base.FreezeWithTimeStruct(t); err != nil {
		panic(err)
	}

	// or use :
	// if err := base.FreezeWithTime(2000, 1, 1, 0, 0, 0, 0, time.Local); err != nil {

	// your unit test code
	println(time.Now())

	// the output is 2000-01-01 00:00:00

	// restore time flow
	base.UnFreezeTime()
}

func ExampleChangeFreezeTime() {
	if err := base.FreezeTime(); err != nil {
		panic(err)
	}

	// your unit test code
	println(time.Now())

	// custom flow
	base.AddDuration(1 * time.Minute)
	println(time.Now())

	// the two outputs differ by 1 minute
	base.UnFreezeTime()
}
