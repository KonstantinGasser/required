package required

import (
	"testing"

	"github.com/matryer/is"
)

type AtomicSimple struct {
	A string `required:"yes"`
	B int    `required:"yes"`
}

func TestAtomicSimple(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		v   AtomicSimple
		err error
	}{
		{
			v: AtomicSimple{
				A: "Konstantin",
				B: 42,
			},
			err: nil,
		},
		{
			v: AtomicSimple{
				A: "",
				B: 42,
			},
			err: ErrDefaultFound,
		},
		{
			v: AtomicSimple{
				A: "Konstantin",
				B: 0,
			},
			err: ErrDefaultFound,
		},
	}

	for _, t := range tt {
		err := Atomic(&t.v)
		is.Equal(t.err, err)
	}
}

type AtomicMinMax struct {
	A string  `required:"yes, min=4, max=6"`
	B int     `required:"yes, min=3,max=12"`
	C float32 `required:"yes, min=1, max=2"`
}

func TestAtomicMinMax(t *testing.T) {

	is := is.New(t)

	tt := []struct {
		v   *AtomicMinMax // setting v as pointer allows to check for nil
		err error
	}{
		{
			v:   nil,
			err: ErrInvalidType,
		},
		{
			// all fields ok
			v: &AtomicMinMax{
				A: "hello",
				B: 8,
				C: float32(1.1),
			},
			err: nil,
		},
		{
			// string condition not ok
			v: &AtomicMinMax{
				A: "Konstantin",
				B: 8,
				C: float32(1.1),
			},
			err: ErrConditionFail,
		},
		{
			// int condition not ok
			v: &AtomicMinMax{
				A: "Konstantin",
				B: 2,
				C: float32(1.1),
			},
			err: ErrConditionFail,
		},
		{
			// float32 condition not ok
			v: &AtomicMinMax{
				A: "Konstantin",
				B: 8,
				C: float32(2.01),
			},
			err: ErrConditionFail,
		},
	}

	for _, t := range tt {
		err := Atomic(t.v)
		is.Equal(t.err, err)

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

var r error

func BenchmarkAllMinMax(b *testing.B) {
	b.ReportAllocs()
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
	var rr error
	for n := 0; n < b.N; n++ {
		rr = Atomic(&tt)
	}
	r = rr
}

type BenchTest struct {
	A string   `required:"yes"`
	B int16    `required:"yes"`
	C []string `required:"yes"`
	D string   `required:"yes"`
	E float64  `required:"yes"`
	F uint32   `required:"yes"`
	G []int    `required:"yes"`
	H int64    `required:"yes"`
	I string   `required:"yes"`
	J string   `required:"yes"`
}

var result error

func BenchmarkAllNoOptions(b *testing.B) {
	b.ReportAllocs()
	tt := BenchTest{
		A: "hello world",
		B: int16(16),
		C: []string{"hello", "world"},
		D: "hello world",
		E: float64(64.64),
		F: uint32(32),
		G: []int{1, 2200, 444567, 1337},
		H: int64(2000010),
		I: "hello world",
		J: "J is the last field in the struct",
	}

	var err error
	for n := 0; n < b.N; n++ {
		err = Atomic(&tt)
	}

	result = err
}

type BenchTestMix struct {
	A string   `required:"yes" m:"jhk"`
	B int16    `required:"yes, min=2, max=200"`
	C []string `required:"yes, min=1"`
	D string   `required:"yes, min=5,max=12"`
	E float64  `required:"yes"`
	F uint32   `required:"yes"`
	G []int    `required:"yes, max=20"`
	H int64    `required:"yes"`
	I string   `required:"yes"`
}

var result3 error

func BenchmarkAllMix(b *testing.B) {
	b.ReportAllocs()
	tt := BenchTestMix{
		A: "hello world",
		B: int16(16),
		C: []string{"hello", "world"},
		D: "hello world",
		E: float64(64.64),
		F: uint32(32),
		G: []int{1, 2200, 444567, 1337},
		H: int64(2000010),
		I: "hello world",
	}

	var err error
	for n := 0; n < b.N; n++ {
		err = Atomic(&tt)
	}

	result3 = err
}
