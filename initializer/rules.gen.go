//DO NOT EDIT THE FILE.

package initializer

import (
	"reflect"
)

func FixedFloat32Rule(value float64) *Rule {
	return FixedFloatRules(value, reflect.Float32)
}

func FixedFloat64Rule(value float64) *Rule {
	return FixedFloatRules(value, reflect.Float64)
}

func FixedIntRule(value int64) *Rule {
	return FixedIntRules(value, reflect.Int)
}

func FixedInt16Rule(value int64) *Rule {
	return FixedIntRules(value, reflect.Int16)
}

func FixedInt32Rule(value int64) *Rule {
	return FixedIntRules(value, reflect.Int32)
}

func FixedInt64Rule(value int64) *Rule {
	return FixedIntRules(value, reflect.Int64)
}

func FixedInt8Rule(value int64) *Rule {
	return FixedIntRules(value, reflect.Int8)
}

func FixedUintRule(value uint64) *Rule {
	return FixedUintRules(value, reflect.Uint)
}

func FixedUint16Rule(value uint64) *Rule {
	return FixedUintRules(value, reflect.Uint16)
}

func FixedUint32Rule(value uint64) *Rule {
	return FixedUintRules(value, reflect.Uint32)
}

func FixedUint64Rule(value uint64) *Rule {
	return FixedUintRules(value, reflect.Uint64)
}

func FixedUint8Rule(value uint64) *Rule {
	return FixedUintRules(value, reflect.Uint8)
}
