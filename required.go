package required

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	// ErrBadSyntax indicates a syntax error in the tag
	ErrBadSyntax = errors.New("error: required: syntax for tag requires is wrong")

	// ErrNotANumber indicates that either the min or max or is not a number
	ErrNotANumber = errors.New("error: required: option for tag is not a number")

	// ErrMaxLowerMin indicates that the max for a field is lower then the min
	ErrMaxLowerMin = errors.New("error: required: max must be > then min option")

	// ErrRequiredFailed indicates that at least one field of the struct does not
	// satisfies the tag condition
	ErrRequiredFailed = errors.New("error: required: field does not satisfy conditions")
)

// All verifies that all fields in a struct with a `required` tag
// pass the tags condition(s)
func All(v interface{}) error {
	_struct := reflect.ValueOf(v).Elem()
	for i := 0; i < _struct.NumField(); i++ {
		f := _struct.Field(i)
		tag := _struct.Type().Field(i).Tag
		tagParams, ok := tag.Lookup("required")
		if !ok {
			continue
		}
		actions := strings.Split(tagParams, ",")
		if len(actions) < 1 {
			return ErrBadSyntax
		}
		if ok := isNotZero(f); !ok {
			return ErrRequiredFailed
		}
		if len(actions) < 2 {
			continue
		}
		min, max, err := parse(actions[1:]...)
		if err != nil {
			return err
		}
		if min == 0 && max == 0 {
			continue
		}
		if ok := isValid(f, min, max); !ok {
			return ErrRequiredFailed
		}
	}
	return nil
}

func isNotZero(value reflect.Value) bool {
	if ok := value.IsValid(); !ok {
		return false
	}
	return !value.IsZero()
}

func isValid(value reflect.Value, min, max int) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover: %v caused: %v\n", value, r)
		}
	}()
	switch value.Type().Kind() {
	case reflect.String:
		v := value.String()
		if (max != 0 && len(v) > max) || (min != 0 && len(v) < min) {
			return false
		}
		return true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := value.Int()
		if (max != 0 && v > int64(max)) || (min != 0 && v < int64(min)) {
			return false
		}
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v := value.Uint()
		if (min != 0 && v < uint64(min)) || (max != 0 && v > uint64(max)) {
			return false
		}
		return true
	case reflect.Float32:
		v := value.Float()
		if (max != 0 && v > float64(max)) || (min != 0 && v < float64(min)) {
			return false
		}
		return true
	case reflect.Slice, reflect.Array:
		v := value.Len()
		if (max != 0 && v > max) || (min != 0 && v < min) {
			return false
		}
		return true
	default:
		return true
	}
}

func parse(opts ...string) (int, int, error) {
	var min, max int
	for _, opt := range opts {
		v := strings.Split(opt, "=")
		i, err := strconv.Atoi(v[1])
		if err != nil {
			return 0, 0, ErrNotANumber
		}
		if strings.TrimSpace(v[0]) == "max" {
			max = i
		}
		if strings.TrimSpace(v[0]) == "min" {
			min = i
		}
	}
	if max != 0 && min != 0 && max < min {
		return 0, 0, ErrMaxLowerMin
	}
	return min, max, nil
}
