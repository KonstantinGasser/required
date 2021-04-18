package required

import (
	"fmt"
	"reflect"
)

// isValid validates whether the value follows the
// tag options (min,max). For intX,uintX,floatX comparison
// is min <= value <= max. For slices and arrays min <= len(value) <= max
func isValid(value reflect.Value) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%v\n %v\n", value, r)
		}
	}()
	switch value.Type().Kind() {
	case reflect.String:
		v := value.String()
		if (minValue != 0 && len(v) < minValue) || (maxValue != 0 && len(v) > maxValue) {
			return false
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := value.Int()
		if (minValue != 0 && v < int64(minValue)) || (maxValue != 0 && v > int64(maxValue)) {
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v := value.Uint()
		if (minValue != 0 && v < uint64(minValue)) || (maxValue != 0 && v > uint64(maxValue)) {
			return false
		}
	case reflect.Float32:
		v := value.Float()
		if (minValue != 0 && v < float64(minValue)) || (maxValue != 0 && v > float64(maxValue)) {
			return false
		}
	case reflect.Slice, reflect.Array:
		v := value.Len()
		if (minValue != 0 && v < minValue) || (maxValue != 0 && v > maxValue) {
			return false
		}
	default:
		return true
	}
	return true
}

// isAllowedType first check if v != nil else ValueOf panics
func isAllowedType(v interface{}) bool {
	return v != nil && !reflect.ValueOf(v).IsNil() && reflect.ValueOf(v).Kind() == reflect.Ptr
}

func isNotZero(v reflect.Value) bool {
	if ok := v.IsValid(); !ok {
		return false
	}
	return !v.IsZero()
}
