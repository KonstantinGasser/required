package required

import (
	"testing"

	"github.com/matryer/is"
)

type TestRequiredOnly struct {
	A string   `required:"yes"`
	B int      `required:"yes"`
	C []string `required:"yes"`
}

func TestAtomicRequiredOnly(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		v   TestRequiredOnly
		err error
	}{
		{
			v: TestRequiredOnly{
				A: "hello world",
				B: 24,
				C: []string{"hello", "friend"},
			},
			err: nil,
		},
		{
			v: TestRequiredOnly{
				A: "hello world",
				B: 24,
				C: nil,
			},
			err: ErrDefaultFound,
		},
		{
			v: TestRequiredOnly{
				A: "hello world",
				B: 0,
				C: []string{"hello", "friend"},
			},
			err: ErrDefaultFound,
		},
	}

	for _, tc := range tt {
		err := Atomic(&tc.v)
		is.Equal(err, tc.err)
	}
}

type TestMinMax struct {
	A string   `required:"yes" min:"4" max:"25"`
	B int      `required:"yes" min:"100" max:"150"`
	C []string `required:"yes" min:"1"`
	D uint16   `required:"yes" max:"24"`
	E float64  `required:"yes" min:"1" max:"2"`
}

func TestAtomicMinMax(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		v   TestMinMax
		err error
	}{
		{
			v: TestMinMax{
				A: "hello world",
				B: 124,
				C: []string{"index_0", "index_1"},
				D: uint16(24),
				E: float64(1.99),
			},
			err: nil,
		},
		{
			v: TestMinMax{
				A: "hel",
				B: 124,
				C: []string{"index_0", "index_1"},
				D: uint16(24),
				E: float64(1.99),
			},
			err: ErrConditionFail,
		},
		{
			v: TestMinMax{
				A: "hello world, hello friend, hello everyone",
				B: 124,
				C: []string{"index_0", "index_1"},
				D: uint16(24),
				E: float64(1.99),
			},
			err: ErrConditionFail,
		},
		{
			v: TestMinMax{
				A: "hello world",
				B: 42,
				C: []string{"index_0", "index_1"},
				D: uint16(24),
				E: float64(1.99),
			},
			err: ErrConditionFail,
		},
		{
			v: TestMinMax{
				A: "hello world",
				B: 142,
				C: []string{"index_0", "index_1"},
				D: uint16(24),
				E: float64(0.45),
			},
			err: ErrConditionFail,
		},
		{
			v: TestMinMax{
				A: "hello world",
				B: 142,
				C: []string{},
				D: uint16(24),
				E: float64(1.45),
			},
			err: ErrConditionFail,
		},
	}

	for _, tc := range tt {
		err := Atomic(&tc.v)
		is.Equal(err, tc.err)
	}
}

type TestNoTag struct {
	A string
	B int
}

func TestAtomicNoTag(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		v   TestNoTag
		err error
	}{
		{
			v:   TestNoTag{},
			err: nil,
		},
		{
			v: TestNoTag{
				A: "",
				B: 42,
			},
			err: nil,
		},
		{
			v: TestNoTag{
				A: "hello",
				B: 0,
			},
			err: nil,
		},
	}

	for _, tc := range tt {
		err := Atomic(&tc.v)
		is.Equal(err, tc.err)
	}
}

type TestStruct struct {
	a string
}
type BenchRequiredOnly struct {
	A string     `required:"yes"`
	B int        `required:"yes"`
	C []string   `required:"yes"`
	D uint16     `required:"yes"`
	E float64    `required:"yes"`
	F string     `required:"yes"`
	G []byte     `required:"yes"`
	H string     `required:"yes"`
	I int16      `required:"yes"`
	J TestStruct `required:"yes"`
}

var benchRes1 error

func BenchmarkRequired(b *testing.B) {
	b.ReportAllocs()

	t := BenchRequiredOnly{
		A: "hello world",
		B: 24,
		C: []string{"hello", "friend"},
		D: uint16(42),
		E: float64(64),
		F: "konstantin.gasser@com",
		G: []byte("mr. robot"),
		H: "KonstantinGasser",
		I: int16(16),
		J: TestStruct{a: "kgasser"},
	}
	var err error
	for n := 0; n < b.N; n++ {
		err = Atomic(&t)
	}
	benchRes1 = err

}

type BenchMinMax struct {
	A string   `required:"yes" min:"4" max:"15"`
	B int      `required:"yes" min:"4"`
	C []string `required:"yes" min:"2"`
	D uint16   `required:"yes" min:"4" max:"15"`
	E float64  `required:"yes" min:"7" max:"10"`
	F string   `required:"yes" max:"15"`
	G []byte   `required:"yes" min:"1"`
	H string   `required:"yes" min:"15" max:"35"`
	I int16    `required:"yes" min:"400"`
	J int16    `required:"yes" min:"4"`
}

var benchRes2 error

func BenchmarkMinMax(b *testing.B) {
	b.ReportAllocs()

	t := BenchMinMax{
		A: "hello world",
		B: 24,
		C: []string{"hello", "friend"},
		D: uint16(14),
		E: float64(9.314),
		F: "konstantin",
		G: []byte("mr. robot"),
		H: "KonstantinGasser@Konstantin.com",
		I: int16(1600),
		J: int16(5),
	}
	var err error
	for n := 0; n < b.N; n++ {
		err = Atomic(&t)
	}
	benchRes2 = err
}
