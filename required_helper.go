package required

import (
	"fmt"
	"reflect"
	"strconv"
)

// isValid validates whether the value follows the
// tag options (min,max). For intX,uintX,floatX comparison
// is min <= value <= max. For slices and arrays min <= len(value) <= max
func isValid(value reflect.Value, minValue, maxValue int) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%v\n %v\n", value, r)
		}
	}()

	// max must be grater min
	if (minValue != 0 && maxValue != 0) && (maxValue < minValue) {
		return ErrMaxLowerMin
	}
	switch value.Type().Kind() {
	case reflect.String, reflect.Slice, reflect.Array:
		v := value.Len()
		if (minValue != 0 && v < minValue) || (maxValue != 0 && v > maxValue) {
			return ErrConditionFail
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := value.Int()
		if (minValue != 0 && v < int64(minValue)) || (maxValue != 0 && v > int64(maxValue)) {
			return ErrConditionFail
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v := value.Uint()
		if (minValue != 0 && v < uint64(minValue)) || (maxValue != 0 && v > uint64(maxValue)) {
			return ErrConditionFail
		}
	case reflect.Float32, reflect.Float64:
		v := value.Float()
		if (minValue != 0 && v < float64(minValue)) || (maxValue != 0 && v > float64(maxValue)) {
			return ErrConditionFail
		}
	default:
		return nil
	}
	return nil
}

func isAllowedType(v interface{}) bool {
	return v != nil && !reflect.ValueOf(v).IsNil() && reflect.ValueOf(v).Kind() == reflect.Ptr
}

func isNotZero(v reflect.Value) bool {
	if ok := v.IsValid(); !ok {
		return false
	}
	return !v.IsZero()
}

// getOpt looks up an option (min/max) and if ok returns the integer
// representation of the string
func getOpt(opt string, tag reflect.StructTag) (int, error) {
	v, ok := tag.Lookup(opt)
	if !ok {
		return 0, nil
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, ErrNotANumber
	}
	return i, nil
}
