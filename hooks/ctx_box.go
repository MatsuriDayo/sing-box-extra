package hooks

import (
	"context"
	"reflect"
	"unsafe"
)

func GetCtxKeys(ctx interface{}) []interface{} {
	var keys = make([]interface{}, 0)

	if a, ok := ctx.(context.Context); ok && ctx != nil {
		ctx = a
	} else {
		return keys
	}

	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr()))
			reflectValueElem := reflectValue.Elem()
			reflectField := contextKeys.Field(i)

			if reflectField.Name == "key" {
				keys = append(keys, reflectValueElem.Interface())
			} else if reflectValueElem.Kind() == reflect.Struct {
				// timerCtx have a "cancelCtx" struct
				keys = append(keys, GetCtxKeys(reflectValue.Interface())...)
			} else if reflectField.Name == "Context" {
				keys = append(keys, GetCtxKeys(reflectValueElem.Interface())...)
			}
		}
	}

	return keys
}

func TransportNameFromContext(ctx context.Context) string {
	for _, k := range GetCtxKeys(ctx) {
		if reflect.TypeOf(k).Name() == "transportKey" {
			core := ctx.Value(k)
			return reflect.ValueOf(core).String()
		}
	}
	return ""
}
