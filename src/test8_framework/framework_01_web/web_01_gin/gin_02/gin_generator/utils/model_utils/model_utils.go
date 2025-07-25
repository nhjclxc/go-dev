package model_utils

import (
	"reflect"
)

// 递归获取结构体的所有字段，包括匿名字段
func flattenFields(v reflect.Value) map[string]reflect.Value {
	fields := make(map[string]reflect.Value)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)
		fieldType := t.Field(i)

		if fieldType.Anonymous && fieldVal.Kind() == reflect.Struct {
			// 匿名结构体递归处理
			nestedFields := flattenFields(fieldVal)
			for k, v := range nestedFields {
				fields[k] = v
			}
		} else {
			fields[fieldType.Name] = fieldVal
		}
	}

	return fields
}

// 泛型结构体字段拷贝，支持嵌套结构体
func CopyStructFields[S any, D any](src S, dst *D) {
	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	dstVal := reflect.ValueOf(dst).Elem()

	srcFields := flattenFields(srcVal)
	dstFields := flattenFields(dstVal)

	for name, srcField := range srcFields {
		if dstField, ok := dstFields[name]; ok {
			if dstField.Type() == srcField.Type() && dstField.CanSet() {
				dstField.Set(srcField)
			}
		}
	}
}
