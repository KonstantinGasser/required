package required

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
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

	// minValue refers to the under boundary a value must have (string and slices use len operation for comparison)
	minValue = 0
	// maxValue refers to the upper boundary a value must have (string and slices use len operation for comparison)
	maxValue = 0
)

// Atomic evaluates each field in the struct(s) which are tagged with
// the `required` tag. If the field value is its default value Atomic
// returns an ErrDefaultFound error. In case options (min/max) are provided
// in the tag, Atomic will return an ErrNotANumber if either option is
// not a number. Lastly Atomic returns an ErrConditionFail in case the
// field value does not satisfy the option(s)
func Atomic(values ...interface{}) error {
	for _, v := range values {
		// interface value must be a pointer else reflect.ValueOf(v).Elem()
		// will panic
		if ok := isAllowedType(v); !ok {
			return ErrInvalidType
		}
		_struct := reflect.ValueOf(v).Elem()
		for i := 0; i < _struct.NumField(); i++ {
			f := _struct.Field(i)
			tag, ok := _struct.Type().Field(i).Tag.Lookup("required")
			if !ok {
				continue
			}
			if ok := isNotZero(f); !ok {
				return ErrDefaultFound
			}
			actions := strings.Split(tag, ",")
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
	}
	return nil
}

// Debug does the same as Atomic however if a conditions fails
// it will not return immediately but rather continue looping through
// the struct(s) collecting all field names which value failed the
// tag conditions. Since no field will be excluded this operation is
// fare more expensive then the Atomic function and generally intended
// for well debugging.
func Debug(values ...interface{}) Issues {
	var issues Issues = make(Issues, len(values))

	for k, v := range values {
		if ok := isAllowedType(v); !ok {
			return nil
		}
		_struct := reflect.ValueOf(v).Elem()
		issues[k] = []Issue{}
		for i := 0; i < _struct.NumField(); i++ {
			f := _struct.Field(i)
			tag, ok := _struct.Type().Field(i).Tag.Lookup("required")
			if !ok {
				continue
			}
			if ok := isNotZero(f); !ok {
				issues[k] = append(issues[k], Issue{
					Struct: _struct.Type().Name(),
					Name:   _struct.Type().Field(i).Name,
					Err:    fmt.Sprintf("field value is its default value: <%v>", f.Interface()),
				})
				continue
			}
			actions := strings.Split(tag, ",")
			if len(actions) < 2 {
				continue
			}
			err := parse(actions[1:]...)
			if err != nil {
				issues[k] = append(issues[k], Issue{
					Struct: _struct.Type().Name(),
					Name:   _struct.Type().Field(i).Name,
					Err:    fmt.Sprintf("tag does has wrong syntax: %s", tag),
				})
				continue
			}
			if minValue == 0 && maxValue == 0 {
				continue
			}
			if ok := isValid(f); !ok {
				issues[k] = append(issues[k], Issue{
					Struct: _struct.Type().Name(),
					Name:   _struct.Type().Field(i).Name,
					Err:    fmt.Sprintf("field valid does not align with conditions: %v != %v", f.Interface(), actions[1:]),
				})
				continue
			}
		}
	}
	return issues
}

// Issues represents any issue found during debugging
type Issues [][]Issue

type Issue struct {
	Struct string
	Name   string
	Err    string
}

// Pretty formats the issues in a descriptive way
func (issues Issues) Pretty() {
	if len(issues) == 0 {
		fmt.Println("required: you're good to go ~ no issues found")
		return
	}
	for i := 0; i < len(issues); i++ {
		fmt.Printf("----------\n")
		for _, j := range issues[i] {
			fmt.Printf("%v:Field: %s\n\t Issue: %s\n", j.Struct, j.Name, j.Err)
		}
		fmt.Printf("----------\n")
	}
}
