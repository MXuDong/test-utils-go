//DO NOT EDIT THE FILE.

// All function of this file will return a rule with specify value and ignore the error. If some function cloud create
// an error, it will return the nil.

package initializer

import (
	"reflect"
)

// ---------------------------------------------------------------------------------------------------------------------
// pattern all

func FixedFloat32MatchAllRule(value float64) *Rule {
	if ruleItem, err := FixedFloatBaseRule(RegexAll, RegexMod, value, reflect.Float32); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedFloat64MatchAllRule(value float64) *Rule {
	if ruleItem, err := FixedFloatBaseRule(RegexAll, RegexMod, value, reflect.Float64); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedIntMatchAllRule(value int64) *Rule {
	if ruleItem, err := FixedIntBaseRule(RegexAll, RegexMod, value, reflect.Int); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt16MatchAllRule(value int64) *Rule {
	if ruleItem, err := FixedIntBaseRule(RegexAll, RegexMod, value, reflect.Int16); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt32MatchAllRule(value int64) *Rule {
	if ruleItem, err := FixedIntBaseRule(RegexAll, RegexMod, value, reflect.Int32); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt64MatchAllRule(value int64) *Rule {
	if ruleItem, err := FixedIntBaseRule(RegexAll, RegexMod, value, reflect.Int64); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt8MatchAllRule(value int64) *Rule {
	if ruleItem, err := FixedIntBaseRule(RegexAll, RegexMod, value, reflect.Int8); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUintMatchAllRule(value uint64) *Rule {
	if ruleItem, err := FixedUintBaseRule(RegexAll, RegexMod, value, reflect.Uint); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint16MatchAllRule(value uint64) *Rule {
	if ruleItem, err := FixedUintBaseRule(RegexAll, RegexMod, value, reflect.Uint16); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint32MatchAllRule(value uint64) *Rule {
	if ruleItem, err := FixedUintBaseRule(RegexAll, RegexMod, value, reflect.Uint32); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint64MatchAllRule(value uint64) *Rule {
	if ruleItem, err := FixedUintBaseRule(RegexAll, RegexMod, value, reflect.Uint64); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint8MatchAllRule(value uint64) *Rule {
	if ruleItem, err := FixedUintBaseRule(RegexAll, RegexMod, value, reflect.Uint8); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// pattern with path

func FixedFloat32MatchJsonPathRule(value float64, path string) *Rule {
	if ruleItem, err := FixedFloatBaseRule(path, JsonPathMod, value, reflect.Float32); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedFloat64MatchJsonPathRule(value float64, path string) *Rule {
	if ruleItem, err := FixedFloatBaseRule(path, JsonPathMod, value, reflect.Float64); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedIntMatchJsonPathRule(value int64, path string) *Rule {
	if ruleItem, err := FixedIntBaseRule(path, JsonPathMod, value, reflect.Int); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt16MatchJsonPathRule(value int64, path string) *Rule {
	if ruleItem, err := FixedIntBaseRule(path, JsonPathMod, value, reflect.Int16); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt32MatchJsonPathRule(value int64, path string) *Rule {
	if ruleItem, err := FixedIntBaseRule(path, JsonPathMod, value, reflect.Int32); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt64MatchJsonPathRule(value int64, path string) *Rule {
	if ruleItem, err := FixedIntBaseRule(path, JsonPathMod, value, reflect.Int64); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt8MatchJsonPathRule(value int64, path string) *Rule {
	if ruleItem, err := FixedIntBaseRule(path, JsonPathMod, value, reflect.Int8); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUintMatchJsonPathRule(value uint64, path string) *Rule {
	if ruleItem, err := FixedUintBaseRule(path, JsonPathMod, value, reflect.Uint); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint16MatchJsonPathRule(value uint64, path string) *Rule {
	if ruleItem, err := FixedUintBaseRule(path, JsonPathMod, value, reflect.Uint16); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint32MatchJsonPathRule(value uint64, path string) *Rule {
	if ruleItem, err := FixedUintBaseRule(path, JsonPathMod, value, reflect.Uint32); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint64MatchJsonPathRule(value uint64, path string) *Rule {
	if ruleItem, err := FixedUintBaseRule(path, JsonPathMod, value, reflect.Uint64); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint8MatchJsonPathRule(value uint64, path string) *Rule {
	if ruleItem, err := FixedUintBaseRule(path, JsonPathMod, value, reflect.Uint8); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// pattern with regex

func FixedFloat32MatchRegexRule(value float64, regex string) *Rule {
	if ruleItem, err := FixedFloatBaseRule(regex, RegexMod, value, reflect.Float32); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedFloat64MatchRegexRule(value float64, regex string) *Rule {
	if ruleItem, err := FixedFloatBaseRule(regex, RegexMod, value, reflect.Float64); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedIntMatchRegexRule(value int64, regex string) *Rule {
	if ruleItem, err := FixedIntBaseRule(regex, RegexMod, value, reflect.Int); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt16MatchRegexRule(value int64, regex string) *Rule {
	if ruleItem, err := FixedIntBaseRule(regex, RegexMod, value, reflect.Int16); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt32MatchRegexRule(value int64, regex string) *Rule {
	if ruleItem, err := FixedIntBaseRule(regex, RegexMod, value, reflect.Int32); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt64MatchRegexRule(value int64, regex string) *Rule {
	if ruleItem, err := FixedIntBaseRule(regex, RegexMod, value, reflect.Int64); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedInt8MatchRegexRule(value int64, regex string) *Rule {
	if ruleItem, err := FixedIntBaseRule(regex, RegexMod, value, reflect.Int8); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUintMatchRegexRule(value uint64, regex string) *Rule {
	if ruleItem, err := FixedUintBaseRule(regex, RegexMod, value, reflect.Uint); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint16MatchRegexRule(value uint64, regex string) *Rule {
	if ruleItem, err := FixedUintBaseRule(regex, RegexMod, value, reflect.Uint16); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint32MatchRegexRule(value uint64, regex string) *Rule {
	if ruleItem, err := FixedUintBaseRule(regex, RegexMod, value, reflect.Uint32); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint64MatchRegexRule(value uint64, regex string) *Rule {
	if ruleItem, err := FixedUintBaseRule(regex, RegexMod, value, reflect.Uint64); err != nil {
		return nil
	} else {
		return ruleItem
	}
}

func FixedUint8MatchRegexRule(value uint64, regex string) *Rule {
	if ruleItem, err := FixedUintBaseRule(regex, RegexMod, value, reflect.Uint8); err != nil {
		return nil
	} else {
		return ruleItem
	}
}
