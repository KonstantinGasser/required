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
		ok       bool
	}{
		{
			// checks for recover
			v:   reflect.ValueOf(nil),
			min: 4,
			max: 32,
			ok:  false,
		},
		{
			v:   reflect.ValueOf(int(32)),
			min: 4,
			max: 32,
			ok:  true,
		},
		{
			v:   reflect.ValueOf(uint(32)),
			min: 4,
			max: 30,
			ok:  false,
		},
		{
			v:   reflect.ValueOf("KonstantinGasser"),
			min: 12,
			max: 25,
			ok:  true,
		},
		{
			v:   reflect.ValueOf("Gasser"),
			min: 12,
			max: 25,
			ok:  false,
		},
		{
			v:   reflect.ValueOf([]int{}),
			min: 1,
			max: 0,
			ok:  false,
		},
		{
			v:   reflect.ValueOf([]int{1, 2, 3}),
			min: 3,
			max: 10,
			ok:  true,
		},
	}

	for _, t := range tt {
		minValue = t.min
		maxValue = t.max
		ok := isValid(t.v)
		is.Equal(t.ok, ok)
	}
}
