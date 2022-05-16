package ran

import (
	"fmt"
	"reflect"
	"strings"
)

type Rander struct {
	hocks []*Rule
}

func (r *Rander) AddRule(rule *Rule) {
	if r.hocks == nil {
		r.hocks = make([]*Rule, 0)
	}
	r.hocks = append(r.hocks, rule)
}

func (r Rander) InjectValue(x interface{}) {
	xv := reflect.ValueOf(x)
	if xv.Kind() != reflect.Ptr {
		// error
	}

	r.initObj(xv.Elem(), x, "$")
}

func (r Rander) initObj(value reflect.Value, obj interface{}, path string) {
	fmt.Println(path)
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
	case reflect.Float64:
		// create by zero
		DefaultFloat64Rule.Hock(path, obj, value)
	}
}

func (r *Rander) invokeHock(k reflect.Kind, t reflect.Type, path string, obj interface{}, value reflect.Value) bool {
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

type TypeInitializer func(typ reflect.Type, tags ...string) interface{}