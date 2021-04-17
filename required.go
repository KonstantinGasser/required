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

	// // ErrRequiredFailed indicates that at least one field of the struct does not
	// // satisfies the tag condition
	// ErrRequiredFailed = errors.New("error: required: field does not satisfy conditions")

	// ErrDefaultFound indicates that the field value is its types default value
	ErrDefaultFound = errors.New("required: at least one tagged field has a default value")

	// ErrConditionFail indicates that the set min or max condition failed
	ErrConditionFail = errors.New("required: at least one tagged field does not satisfies its tag condition")
)

var (
	// minValue refers to the under boundary a value must have (string and slices use len operation)
	minValue = 0
	// maxValue refers to the upper boundary a value must have (string and slices use len operation)
	maxValue = 0
)

// All verifies that all fields in a struct with a `required` tag
// pass the tags condition(s)
func All(v interface{}) error {
	_struct := reflect.ValueOf(v).Elem()
	for i := 0; i < _struct.NumField(); i++ {
		f := _struct.Field(i)
		tag, ok := _struct.Type().Field(i).Tag.Lookup("required")
		if !ok {
			continue
		}
		actions := strings.Split(tag, ",")
		if ok := isNotZero(f); !ok {
			return ErrDefaultFound
		}
		if len(actions) < 2 {
			continue
		}
		err := parse(actions[1:]...)
		if err != nil {
			return err
		}
		if minValue == 0 && maxValue == 0 {
			continue
		}
		if ok := isValid(f); !ok {
			return ErrConditionFail
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

func isValid(value reflect.Value) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover: %v\n", r)
		}
	}()
	switch value.Type().Kind() {
	case reflect.String:
		v := value.String()
		if (maxValue != 0 && len(v) > maxValue) || (minValue != 0 && len(v) < minValue) {
			return false
		}
		return true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := value.Int()
		if (maxValue != 0 && v > int64(maxValue)) || (minValue != 0 && v < int64(minValue)) {
			return false
		}
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v := value.Uint()
		if (minValue != 0 && v < uint64(minValue)) || (maxValue != 0 && v > uint64(maxValue)) {
			return false
		}
		return true
	case reflect.Float32:
		v := value.Float()
		if (maxValue != 0 && v > float64(maxValue)) || (minValue != 0 && v < float64(minValue)) {
			return false
		}
		return true
	case reflect.Slice, reflect.Array:
		v := value.Len()
		if (maxValue != 0 && v > maxValue) || (minValue != 0 && v < minValue) {
			return false
		}
		return true
	default:
		return true
	}
}

func parse(opts ...string) error {
	for _, opt := range opts {
		v := strings.Split(opt, "=")
		i, err := strconv.Atoi(v[1])
		if err != nil {
			return ErrNotANumber
		}
		if strings.TrimSpace(v[0]) == "max" {
			maxValue = i
		}
		if strings.TrimSpace(v[0]) == "min" {
			minValue = i
		}
	}
	if minValue != 0 && len(opts) > 1 && maxValue < minValue {
		return ErrMaxLowerMin
	}
	return nil
}
