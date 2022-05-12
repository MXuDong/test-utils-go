package patch

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

// patch is an applied patch
// needed to undo a patch
type patch struct {
	originalBytes []byte
	replacement   *reflect.Value
}

var (
	operatorLock *sync.Mutex
	patches      map[uintptr]patch
)

func init() {
	operatorLock = &sync.Mutex{}
	patches = make(map[uintptr]patch)
}

type value struct {
	_   uintptr
	ptr unsafe.Pointer
}

func getPtr(v reflect.Value) unsafe.Pointer {
	return (*value)(unsafe.Pointer(&v)).ptr
}

type Record struct {
	target      reflect.Value
	replacement reflect.Value
}

func (g *Record) Unpatch() {
	restoreValue(g.target)
}

func (g *Record) Restore() {
	coverValue(g.target, g.replacement)
}

//--------------------------------------------------

// Cover replaces a function with another
func Cover(target, replacement interface{}) *Record {
	t := reflect.ValueOf(target)
	r := reflect.ValueOf(replacement)
	coverValue(t, r)

	return &Record{t, r}
}

// Restore removes any monkey patches on target.
func Restore(target interface{}) bool {
	return restoreValue(reflect.ValueOf(target))
}

func RestoreSome(target ...interface{}) {
	if target != nil {
		for _, item := range target {
			Restore(item)
		}
	}
}

func RestoreAll() {
	operatorLock.Lock()
	defer operatorLock.Unlock()
	for target, p := range patches {
		restore(target, p)
		delete(patches, target)
	}
}

//--------------------------------------------------

// CoverInstanceFunction replaces an instance method methodName for the type target with replacement.
func CoverInstanceFunction(target reflect.Type, methodName string, replacement interface{}) *Record {
	m, ok := target.MethodByName(methodName)
	if !ok {
		panic(fmt.Sprintf("unknown method %s", methodName))
	}
	r := reflect.ValueOf(replacement)
	coverValue(m.Func, r)

	return &Record{m.Func, r}
}

// RestoreInstanceMethod removes the patch on methodName of the target.
func RestoreInstanceMethod(target reflect.Type, methodName string) bool {
	m, ok := target.MethodByName(methodName)
	if !ok {
		panic(fmt.Sprintf("unknown method %s", methodName))
	}
	return restoreValue(m.Func)
}

//--------------------------------------------------

func coverValue(target, replacement reflect.Value) {
	operatorLock.Lock()
	defer operatorLock.Unlock()

	if target.Kind() != reflect.Func {
		panic("target has to be a Func")
	}

	if replacement.Kind() != reflect.Func {
		panic("replacement has to be a Func")
	}

	if target.Type() != replacement.Type() {
		panic(fmt.Sprintf("target and replacement have to have the same type %s != %s", target.Type(), replacement.Type()))
	}

	if patch, ok := patches[target.Pointer()]; ok {
		restore(target.Pointer(), patch)
	}

	bytes := replaceFunction(target.Pointer(), (uintptr)(getPtr(replacement)))
	patches[target.Pointer()] = patch{bytes, &replacement}
}

// Restore removes a monkeypatch from the specified function
func restoreValue(target reflect.Value) bool {
	operatorLock.Lock()
	defer operatorLock.Unlock()
	patch, ok := patches[target.Pointer()]
	if !ok {
		return false
	}
	restore(target.Pointer(), patch)
	delete(patches, target.Pointer())
	return true
}

func restore(target uintptr, p patch) {
	cp(target, p.originalBytes)
}
