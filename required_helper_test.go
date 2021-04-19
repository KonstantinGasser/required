package required

import (
	"reflect"
	"testing"

	"github.com/matryer/is"
)

func TestIsValid(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		v        reflect.Value
		min, max int
		err      error
	}{
		{
			v:   reflect.ValueOf(int(32)),
			min: 4,
			max: 32,
			err: nil,
		},
		{
			v:   reflect.ValueOf(uint(32)),
			min: 4,
			max: 30,
			err: ErrConditionFail,
		},
		{
			v:   reflect.ValueOf("KonstantinGasser"),
			min: 12,
			max: 25,
			err: nil,
		},
		{
			v:   reflect.ValueOf("Gasser"),
			min: 12,
			max: 25,
			err: ErrConditionFail,
		},
		{
			v:   reflect.ValueOf([]int{}),
			min: 1,
			max: 0,
			err: ErrConditionFail,
		},
		{
			v:   reflect.ValueOf([]int{1, 2, 3}),
			min: 3,
			max: 10,
			err: nil,
		},
		{
			v:   reflect.ValueOf(int(12)),
			min: 10,
			max: 3,
			err: ErrMaxLowerMin,
		},
	}

	for _, t := range tt {
		err := isValid(t.v, t.min, t.max)
		is.Equal(t.err, err)
	}
}
