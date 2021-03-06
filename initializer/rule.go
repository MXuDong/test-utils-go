package initializer

import (
	"reflect"
	"regexp"
	"strings"
)

// RuleHockFunction return true when value is real inject, if not inject, return false.
type RuleHockFunction func(path string, obj interface{}, valueObj reflect.Value) bool

// Rule define how value will be created. Use with Initializer.
type Rule struct {
	// match mod, default is pattern
	mod int

	// the json path
	pattern string

	// the regex pattern
	_regexP *regexp.Regexp

	// Hock should inject the value to valueObj. Only invoke Hock func when type if specify type.
	// This function's choice of type is completely trusted, it means the valueObj can be set value directly
	// by valueObj.Set*(*)
	Hock RuleHockFunction

	// if typ is nil, allow any type, else, only allow value which type can be match in typ,
	typ map[reflect.Type]struct{}
	// if kin is nil, allow any kind, else, only allow value which kind can be match in kin
	kin map[reflect.Kind]struct{}
}

const (
	patternPrefix = "p:"

	JsonPathMod = 1
	RegexMod    = 2
)

// NewRule return a new rule.
func NewRule(pattern string, mod int, hock RuleHockFunction, triggerType []reflect.Type, triggerKind []reflect.Kind) (*Rule, error) {
	r := &Rule{
		Hock: hock,
	}
	if err := r.SetPattern(pattern, mod); err != nil {
		return nil, err
	}

	if triggerKind != nil {
		r.AddKind(triggerKind...)
	}
	if triggerType != nil {
		r.AddType(triggerType...)
	}

	return r, nil
}

// inner initializer, ignore the err
func newRule(pattern string, mod int, hock RuleHockFunction, triggerType []reflect.Type, triggerKind []reflect.Kind) *Rule {
	r, _ := NewRule(pattern, mod, hock, triggerType, triggerKind)
	return r
}

// GetType return the type of rule.
func (r Rule) GetType() map[reflect.Type]struct{} {
	return r.typ
}

// AddType append a types of rule
func (r *Rule) AddType(t ...reflect.Type) {
	if r.typ == nil {
		r.typ = make(map[reflect.Type]struct{}, len(t))
	}
	if t != nil {
		for _, item := range t {
			r.typ[item] = struct{}{}
		}
	}
}

// DeleteType will delete the type
func (r *Rule) DeleteType(t reflect.Type) bool {
	if _, ok := r.typ[t]; ok {
		delete(r.typ, t)
		return true
	}
	return false
}

// IsTypeAllow return true when specify type can pass.
func (r *Rule) IsTypeAllow(t reflect.Type) bool {
	if len(r.typ) == 0 {
		return true
	}

	_, ok := r.typ[t]
	return ok
}

// GetKind return the kinds of rule.
func (r Rule) GetKind() map[reflect.Kind]struct{} {
	return r.kin
}

// AddKind append a kind of rule
func (r *Rule) AddKind(t ...reflect.Kind) {
	if r.kin == nil {
		r.kin = make(map[reflect.Kind]struct{}, len(t))
	}
	if t != nil {
		for _, item := range t {
			r.kin[item] = struct{}{}
		}
	}
}

// DeleteKind will delete the kind.
func (r *Rule) DeleteKind(t reflect.Kind) bool {
	if _, ok := r.kin[t]; ok {
		delete(r.kin, t)
		return true
	}
	return false
}

// IsKindAllow return true when specify kind can pass.
func (r *Rule) IsKindAllow(t reflect.Kind) bool {
	if len(r.kin) == 0 {
		return true
	}

	_, ok := r.kin[t]
	return ok
}

// SetPattern set the pattern to rule, and only path can be rule hold, the hock will be invoked.
func (r *Rule) SetPattern(pattern string, mod int) error {
	var err error
	r.mod = mod

	switch mod {
	case RegexMod:
		r._regexP, err = regexp.Compile(pattern)
		if err != nil {
			return err
		}
	case JsonPathMod:
	default:
		r.pattern = pattern
	}
	return nil
}

// MatchMod return the mod in rule
func (r Rule) MatchMod() int {
	return r.mod
}

// MatchPattern return ture when path can be r hold.
func (r Rule) MatchPattern(path string) bool {
	switch r.mod {
	case RegexMod:
		return r.matchRegex(path)
	case JsonPathMod:
	default:
		return r.matchJsonPath(path)
	}
	// set default value.
	return false
}

func (r Rule) matchRegex(path string) bool {
	if r._regexP == nil {
		// in common, it will not trigger
		return false
	}
	return r._regexP.MatchString(path)
}

func (r Rule) matchJsonPath(path string) bool {
	return strings.HasSuffix(path, r.pattern)
}

// ---------------------------------------------------------------------------------------------------------------------
// default rule

const RegexAll = ".*"

func newDefaultRule(hock RuleHockFunction, kind ...reflect.Kind) *Rule {
	return newRule(RegexAll, RegexMod, hock, nil, kind)
}

var (
	DefaultFloatRule = newDefaultRule(func(path string, obj interface{}, valueObj reflect.Value) bool {
		valueObj.SetFloat(0)
		return true
	}, reflect.Float64, reflect.Float32)
	DefaultIntRule = newDefaultRule(func(path string, obj interface{}, valueObj reflect.Value) bool {
		valueObj.SetInt(0)
		return true
	}, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64)
	DefaultUIntRule = newDefaultRule(func(path string, obj interface{}, valueObj reflect.Value) bool {
		valueObj.SetUint(0)
		return true
	}, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64)
	DefaultBooleanRule = newDefaultRule(func(path string, obj interface{}, valueObj reflect.Value) bool {
		valueObj.SetBool(false)
		return true
	}, reflect.Bool)
	DefaultStringRule = newDefaultRule(func(path string, obj interface{}, valueObj reflect.Value) bool {
		valueObj.SetString("")
		return true
	}, reflect.String)
	DefaultPtrTypeRule = newDefaultRule(func(path string, obj interface{}, valueObj reflect.Value) bool {
		v := reflect.New(valueObj.Type())
		valueObj.Set(v.Elem())
		return true
	}, reflect.Map, reflect.Array, reflect.Slice)
)

func FixedUintBaseRule(pattern string, mod int, value uint64, kinds ...reflect.Kind) (*Rule, error) {
	return NewRule(pattern, mod, func(path string, obj interface{}, valueObj reflect.Value) bool {
		valueObj.SetUint(value)
		return true
	}, nil, kinds)
}

func FixedIntBaseRule(pattern string, mod int, value int64, kinds ...reflect.Kind) (*Rule, error) {
	return NewRule(pattern, mod, func(path string, obj interface{}, valueObj reflect.Value) bool {
		valueObj.SetInt(value)
		return true
	}, nil, kinds)
}

func FixedBoolBaseRule(pattern string, mod int, value bool, kinds ...reflect.Kind) (*Rule, error) {
	return NewRule(pattern, mod, func(path string, obj interface{}, valueObj reflect.Value) bool {
		valueObj.SetBool(value)
		return true
	}, nil, kinds)
}

func FixedFloatBaseRule(pattern string, mod int, value float64, kinds ...reflect.Kind) (*Rule, error) {
	return NewRule(pattern, mod, func(path string, obj interface{}, valueObj reflect.Value) bool {
		valueObj.SetFloat(value)
		return true
	}, nil, kinds)
}

func FixedStringBaseRule(pattern string, mod int, value string, kinds ...reflect.Kind) (*Rule, error) {
	return NewRule(pattern, mod, func(path string, obj interface{}, valueObj reflect.Value) bool {
		valueObj.SetString(value)
		return true
	}, nil, kinds)
}
