package required

import (
	"strconv"
	"strings"
)

// var (
// 	// ErrNoPtr indicated that input type is no pointer
// 	ErrNoPtr = errors.New("required: v of type interface{} in required.All must be pointer to a struct")
// 	// ErrBadSyntax indicates a syntax error in the tag
// 	ErrBadSyntax = errors.New("required: syntax for tag requires is wrong")

// 	// ErrNotANumber indicates that either the min or max or is not a number
// 	ErrNotANumber = errors.New("required: option for tag is not a number")

// 	// ErrMaxLowerMin indicates that the max for a field is lower then the min
// 	ErrMaxLowerMin = errors.New("required: max must be grater (>) then min option")

// 	// ErrDefaultFound indicates that the field value is its types default value
// 	ErrDefaultFound = errors.New("required: at least one tagged field has a default value")

// 	// ErrConditionFail indicates that the set min or max condition failed
// 	ErrConditionFail = errors.New("required: at least one tagged field does not satisfies its tag condition")

// 	// minValue refers to the under boundary a value must have (string and slices use len operation for comparison)
// 	minValue = 0
// 	// maxValue refers to the upper boundary a value must have (string and slices use len operation for comparison)
// 	maxValue = 0
// )

// All(v ...interface{}) error
// Single(v interface{}, opts ...func() string) error
// Multi(v ...interface{}, opts ...func() string) error
// GetValid(v []interface{}) ([]interface{}, error)

// type Func func() string
// type WithFeedback func() string
// required.Match(v ...interface{}, opts ...Func) error
// required.WithFeedback(v ...interface{}) ([]string, error)

// All verifies that all fields in a struct with a `required` tag
// pass the tags condition(s)

// All returns an error if fields tagged with `required` fail the validation.
// if v is not a pointer to a struct All returns an ErrNoPtr error.
// All is atomic which means that either all conditions pass with no error or
// if one field fails the whole operation will fail with ErrConditionFail or ErrDefaultFound.
// If one of the options evaluates to not a number All returns an ErrNotANumber error
// func All(v interface{}) error {
// 	if ok := isPtr(reflect.ValueOf(v).Kind()); !ok {
// 		return ErrNoPtr
// 	}
// 	_struct := reflect.ValueOf(v).Elem()
// 	for i := 0; i < _struct.NumField(); i++ {
// 		f := _struct.Field(i)
// 		tag, ok := _struct.Type().Field(i).Tag.Lookup("required")
// 		if !ok {
// 			continue
// 		}
// 		actions := strings.Split(tag, ",")
// 		if ok := isNotZero(f); !ok {
// 			return ErrDefaultFound
// 		}
// 		if len(actions) < 2 {
// 			continue
// 		}
// 		err := parse(actions[1:]...)
// 		if err != nil {
// 			return err
// 		}
// 		if minValue == 0 && maxValue == 0 {
// 			continue
// 		}
// 		if ok := isValid(f); !ok {
// 			return ErrConditionFail
// 		}
// 	}
// 	return nil
// }

// isValid validates whether the value follows the
// tag options (min,max). For intX,uintX,floatX comparison
// is min <= value <= max. For slices and arrays min <= len(value) <= max
// func isValid(value reflect.Value) bool {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Printf("recover: %v\n", r)
// 		}
// 	}()
// 	switch value.Type().Kind() {
// 	case reflect.String:
// 		v := value.String()
// 		if (maxValue != 0 && len(v) > maxValue) || (minValue != 0 && len(v) < minValue) {
// 			return false
// 		}
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		v := value.Int()
// 		if (maxValue != 0 && v > int64(maxValue)) || (minValue != 0 && v < int64(minValue)) {
// 			return false
// 		}
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
// 		v := value.Uint()
// 		if (minValue != 0 && v < uint64(minValue)) || (maxValue != 0 && v > uint64(maxValue)) {
// 			return false
// 		}
// 	case reflect.Float32:
// 		v := value.Float()
// 		if (maxValue != 0 && v > float64(maxValue)) || (minValue != 0 && v < float64(minValue)) {
// 			return false
// 		}
// 	case reflect.Slice, reflect.Array:
// 		v := value.Len()
// 		if (maxValue != 0 && v > maxValue) || (minValue != 0 && v < minValue) {
// 			return false
// 		}
// 	default:
// 		return true
// 	}
// 	return true
// }

// parse assigns the min and max value from the tag
// options. in -> "min=42","max=99". If either option
// holds not a number parse returns a ErrNotANumber error
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

// func isPtr(kind reflect.Kind) bool {
// 	return (kind == reflect.Ptr)
// }

// func isNotZero(v reflect.Value) bool {
// 	if ok := v.IsValid(); !ok {
// 		return false
// 	}
// 	return !v.IsZero()
// }
