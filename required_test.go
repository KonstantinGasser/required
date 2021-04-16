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

type TestAllStruct struct {
	A string     `required:"yes, min=4,max=20"`
	B int        `required:"yes,min=2,max=4"`
	C int32      `required:"yes,min=2,max=4"`
	D uint32     `required:"yes,min=2,max=4"`
	E float32    `required:"yes,min=2,max=20"`
	F *testing.T `required:"yes"`
}

func TestAll(t *testing.T) {
	is := is.New(t)

	tt := TestAllStruct{
		A: "hello world",
		B: 3,
		C: int32(3),
		D: uint32(3),
		E: float32(12.6),
		F: t,
	}

	err := All(&tt)
	is.NoErr(err)
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
