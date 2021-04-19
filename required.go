package required

import (
	"errors"
	"reflect"
)

var (
	// ErrInvalidType indicated that input type is no pointer
	ErrInvalidType = errors.New("required: arguments in required.Atomic must be pointer to a struct")
	// ErrBadSyntax indicates a syntax error in the tag
	ErrBadSyntax = errors.New("required: syntax for tag requires is wrong")

	// ErrNotANumber indicates that either the min or max or is not a number
	ErrNotANumber = errors.New("required: option for tag is not a number")

	// ErrMaxLowerMin indicates that the max for a field is lower then the min
	ErrMaxLowerMin = errors.New("required: max must be grater (>) then min option")

	// ErrDefaultFound indicates that the field value is its types default value
	ErrDefaultFound = errors.New("required: at least one tagged field has a default value")

	// ErrConditionFail indicates that the set min or max condition failed
	ErrConditionFail = errors.New("required: at least one tagged field does not satisfies its tag condition")
)

// Atomic evaluates if the tagged fields satisfy the condition
// to be either not the zero value or confirm with the min/max option.
// It acts atomically meaning if only one field does not confirm either
// condition the whole operate evaluates as failed.
// If either vs's type is not a pointer to a struct Atomic returns an ErrInvalidType.
// If a tagged field is the zero value Atomic returns an ErrDefaultFound.
// If options are set but are not a number or max < min, Atomic returns ErrNotANumber / ErrMaxLowerMin.
func Atomic(vs ...interface{}) error {
	for _, v := range vs {
		if ok := isAllowedType(v); !ok {
			return ErrInvalidType
		}
		currElem := reflect.ValueOf(v).Elem()
		for i := 0; i < currElem.NumField(); i++ {
			tag := currElem.Type().Field(i).Tag
			_, ok := tag.Lookup("required")
			if !ok {
				continue
			}
			field := currElem.Field(i)
			if ok := isNotZero(field); !ok {
				return ErrDefaultFound
			}
			optMin, err := getOpt("min", tag)
			if err != nil {
				return err
			}
			optMax, err := getOpt("max", tag)
			if err != nil {
				return err
			}
			if err := isValid(field, optMin, optMax); err != nil {
				return err
			}
		}
	}
	return nil
}
