package initializer

import (
	"fmt"
	"reflect"
	"strings"
)

type Initializer struct {
	hocks []*Rule
}

func (r *Initializer) AddRule(rule *Rule) {
	if r.hocks == nil {
		r.hocks = make([]*Rule, 0)
	}
	r.hocks = append(r.hocks, rule)
}

func (r Initializer) InjectValue(x interface{}) error {
	xv := reflect.ValueOf(x)
	if xv.Kind() != reflect.Ptr {
		return fmt.Errorf("Object should a pointer ")
	}

	r.initObj(xv.Elem(), x, "$")
	return nil
}

func (r Initializer) initObj(value reflect.Value, obj interface{}, path string) {
	if r.invokeHock(value.Kind(), value.Type(), path, obj, value) {
		return
	}

	switch value.Kind() {
	case reflect.Ptr:
		// need create object
		newValue := reflect.New(value.Type().Elem())
		value.Set(newValue)
		r.initObj(newValue.Elem(), obj, path)
	case reflect.Struct:
		// create by zero
		valueT := value.Type()
		for fieldIndex := 0; fieldIndex < valueT.NumField(); fieldIndex++ {
			field := valueT.Field(fieldIndex)
			fieldValue := value.FieldByName(field.Name)
			r.initObj(fieldValue, obj, strings.Join([]string{path, field.Name}, "."))
		}
	case reflect.Float64, reflect.Float32:
		DefaultFloatRule.Hock(path, obj, value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		DefaultIntRule.Hock(path, obj, value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		DefaultUIntRule.Hock(path, obj, value)
	case reflect.Bool:
		DefaultBooleanRule.Hock(path, obj, value)
	case reflect.String:
		DefaultStringRule.Hock(path, obj, value)
	case reflect.Map, reflect.Slice, reflect.Array:
		DefaultPtrTypeRule.Hock(path, obj, value)
	case reflect.Interface:
	case reflect.Func:
	case reflect.Chan:
	default:
		// do nothing
		break
	}
}

func (r *Initializer) invokeHock(k reflect.Kind, t reflect.Type, path string, obj interface{}, value reflect.Value) bool {
	if len(r.hocks) == 0 {
		return false
	}
	for _, hock := range r.hocks {
		if hock.IsKindAllow(k) && hock.IsTypeAllow(t) && hock.MatchPattern(path) {
			if hock.Hock(path, obj, value) {
				return true
			}
		}
	}
	return false
}
