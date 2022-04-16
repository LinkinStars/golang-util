package gu

import (
	"fmt"
	"reflect"
)

// GetPtrTypeValue returns the value of the pointer and type
func GetPtrTypeValue(obj interface{}) (objT reflect.Type, objV reflect.Value, err error) {
	objT = reflect.TypeOf(obj)
	objV = reflect.ValueOf(obj)
	if !IsStructPtr(objT) && !IsSlicePtr(objT) {
		return nil, reflect.Value{}, fmt.Errorf("%v must be a struct pointer", obj)
	}
	objT = objT.Elem()
	objV = objV.Elem()
	return objT, objV, nil
}

// IsStructPtr check if struct pointer
func IsStructPtr(t reflect.Type) bool {
	return t != nil && t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

// IsSlicePtr check if slice pointer
func IsSlicePtr(t reflect.Type) bool {
	return t != nil && t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Slice
}
