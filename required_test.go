package required

import (
	"testing"

	"github.com/matryer/is"
)

func TestParse(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		TagList []string
		Max     int
		Min     int
		Err     error
	}{
		{
			TagList: []string{},
			Max:     0,
			Min:     0,
			Err:     nil,
		},
		{
			TagList: []string{"min=4", "max=6"},
			Max:     6,
			Min:     4,
			Err:     nil,
		},
		{
			TagList: []string{"min=4", "max=0"},
			Max:     0,
			Min:     4,
			Err:     nil,
		},
		{
			TagList: []string{"min=0", "max=4"},
			Max:     4,
			Min:     0,
			Err:     nil,
		},
		{
			TagList: []string{"min=12", "max=6"},
			Max:     0,
			Min:     0,
			Err:     ErrMaxLowerMin,
		},
	}

	for _, tc := range tt {
		min, max, err := parse(tc.TagList...)
		is.Equal(err, tc.Err)
		is.Equal(max, tc.Max)
		is.Equal(min, tc.Min)
	}
}

type TestStringNoOpts struct {
	A string `required:"yes"`
}

type TestStringMinMax struct {
	A string `required:"yes, min=4, max=15"`
}

type TestIntNoOpts struct {
	A int16 `required:"yes"`
}

type TestIntMinMax struct {
	A int32 `required:"yes, min=6,max=12"`
}

type TestPtr struct {
	A *TestStringNoOpts `required:"yes"`
}

type TestSliceNoOpts struct {
	A []int `required:"yes"`
}
type TestScliceMinMax struct {
	A []int `required:"yes, min=3, max=7"`
}

type TestMinNA struct {
	A string `required:"yes, min=two, max=4"`
}
type TestMaxNA struct {
	A string `required:"yes, min=2, max=six"`
}

type TestMaxLowerMin struct {
	A string `required:"yes min=10, max=8"`
}

func TestAll(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		v   interface{}
		err error
	}{
		{
			v:   TestMinNA{A: "does not matter"},
			err: ErrNotANumber,
		},
		{
			v:   TestMaxNA{A: "does not matter"},
			err: ErrNotANumber,
		},
		{
			v:   TestMaxLowerMin{A: "dose not matter"},
			err: ErrMaxLowerMin,
		},
		{
			v:   TestStringNoOpts{A: "hello world"},
			err: nil,
		},
		{
			v:   TestStringNoOpts{A: ""},
			err: ErrRequiredFailed,
		},
		{
			v:   TestStringMinMax{A: "hello world"},
			err: nil,
		},
		{
			v:   TestStringMinMax{A: "hel"},
			err: ErrRequiredFailed,
		},
		{
			v:   TestStringMinMax{A: "hello, this is a word or tow to long"},
			err: ErrRequiredFailed,
		},
		{
			v:   TestIntNoOpts{A: int16(42)},
			err: nil,
		},
		{
			v:   TestIntNoOpts{A: int16(0)},
			err: ErrRequiredFailed,
		},
		{
			v:   TestIntMinMax{A: int32(8)},
			err: nil,
		},
		{
			v:   TestIntMinMax{A: int32(5)},
			err: ErrRequiredFailed,
		},
		{
			v:   TestIntMinMax{A: int32(14)},
			err: ErrRequiredFailed,
		},
		{
			v:   TestPtr{A: &TestStringNoOpts{A: "hello"}},
			err: nil,
		},
		{
			v:   TestPtr{A: nil},
			err: ErrRequiredFailed,
		},
		{
			v:   TestSliceNoOpts{A: []int{24, 42}},
			err: nil,
		},
		{
			v:   TestScliceMinMax{A: []int{}},
			err: ErrRequiredFailed,
		},
		{
			v:   TestScliceMinMax{A: []int{1, 2}},
			err: ErrRequiredFailed,
		},
		{
			v:   TestScliceMinMax{A: []int{1, 2, 3, 4}},
			err: nil,
		},
		{
			v:   TestScliceMinMax{A: []int{1, 2, 3, 4, 5, 6, 7, 8}},
			err: ErrRequiredFailed,
		},
	}

	for _, t := range tt {
		switch t.v.(type) {
		case TestStringNoOpts:
			v, _ := t.v.(TestStringNoOpts)
			err := All(&v)
			is.Equal(err, t.err)
		case TestStringMinMax:
			v, _ := t.v.(TestStringMinMax)
			err := All(&v)
			is.Equal(err, t.err)
		case TestIntNoOpts:
			v, _ := t.v.(TestIntNoOpts)
			err := All(&v)
			is.Equal(err, t.err)
		case TestIntMinMax:
			v, _ := t.v.(TestIntMinMax)
			err := All(&v)
			is.Equal(err, t.err)
		case TestPtr:
			v, _ := t.v.(TestPtr)
			err := All(&v)
			is.Equal(err, t.err)
		case TestSliceNoOpts:
			v, _ := t.v.(TestSliceNoOpts)
			err := All(&v)
			is.Equal(err, t.err)
		case TestScliceMinMax:
			v, _ := t.v.(TestScliceMinMax)
			err := All(&v)
			is.Equal(err, t.err)
		}
	}
}

type BenchTestMinMax struct {
	A string `required:"yes, min=2,max=25"`
	B string `required:"yes, min=2,max=25"`
	C string `required:"yes, min=2,max=25"`
	D string `required:"yes, min=2,max=25"`
	E string `required:"yes, min=2,max=25"`
	F string `required:"yes, min=2,max=25"`
	G string `required:"yes, min=2,max=25"`
	H string `required:"yes, min=2,max=25"`
	I string `required:"yes, min=2,max=25"`
	J string `required:"yes, min=2,max=25"`
}

func BenchmarkAllMinMax(b *testing.B) {

	tt := BenchTestMinMax{
		A: "hello world",
		B: "hello world",
		C: "hello world",
		D: "hello world",
		E: "hello world",
		F: "hello world",
		G: "hello world",
		H: "hello world",
		I: "hello world",
		J: "hello world",
	}
	for n := 0; n < b.N; n++ {
		All(&tt)
	}
}

type BenchTest struct {
	A string `required:"yes"`
	B string `required:"yes"`
	C string `required:"yes"`
	D string `required:"yes"`
	E string `required:"yes"`
	F string `required:"yes"`
	G string `required:"yes"`
	H string `required:"yes"`
	I string `required:"yes"`
	J string `required:"yes"`
}

var result error

func BenchmarkAllNoOptions(b *testing.B) {

	tt := BenchTest{
		A: "hello world",
		B: "hello world",
		C: "hello world",
		D: "hello world",
		E: "hello world",
		F: "hello world",
		G: "hello world",
		H: "hello world",
		I: "hello world",
		J: "hello world",
	}

	var err error
	for n := 0; n < b.N; n++ {
		err = All(&tt)
	}

	result = err
}
