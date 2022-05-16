package initializer

import (
	"fmt"
	"testing"
)

type A struct {
	Value float64
}

type B struct {
	Value float64

	CValue *C
}

type C struct {
	Value float64
}

func TestNewType(t *testing.T) {

	x := struct {
		Value  *float64
		IntV   *int
		Int8V  *int8
		Int16V *int16
		Int32V *int32
		Int64V *int64

		UintV   *uint
		Uint8V  *uint8
		Uint16V *uint16
		Uint32V *uint32
		Uint64V *uint64

		BoolV *bool

		StringV *string

		CommonMap map[string]string
		StructMap map[string]A

		CommonArray []int
		StructArray []A

		X1 struct {
			Value **float64
		}

		Av *A
	}{}
	y := &x

	r := Initializer{}

	r.AddRule(FixedFloat64MatchAllRule(100))

	r.InjectValue(y)

	fmt.Println("Float====")
	fmt.Println(*x.Value)
	fmt.Println(**x.X1.Value)
	fmt.Println("Int=====")
	fmt.Println(*x.IntV)
	fmt.Println(*x.Int8V)
	fmt.Println(*x.Int16V)
	fmt.Println(*x.Int32V)
	fmt.Println(*x.Int64V)
	fmt.Println("UInt=====")
	fmt.Println(*x.UintV)
	fmt.Println(*x.Uint8V)
	fmt.Println(*x.Uint16V)
	fmt.Println(*x.Uint32V)
	fmt.Println(*x.Uint64V)
	fmt.Println("Boolean=====")
	fmt.Println(*x.BoolV)
	fmt.Println("String=====")
	fmt.Println(*x.StringV)
	fmt.Println("Struct=====")
	fmt.Println(x.Av.Value)
	fmt.Println("Map=====")
	fmt.Println(x.CommonMap)
	fmt.Println(x.StructMap)
	fmt.Println("Array======")
	fmt.Println(x.CommonArray)
	fmt.Println(x.StructArray)
}
