package required

// import (
// 	"testing"

// 	"github.com/matryer/is"
// )

// func TestParse(t *testing.T) {
// 	is := is.New(t)

// 	tt := []struct {
// 		Tag []string
// 		Max int
// 		Min int
// 		Err error
// 	}{
// 		{
// 			Tag: []string{},
// 			Max: 0,
// 			Min: 0,
// 			Err: nil,
// 		},
// 		{
// 			Tag: []string{"min=4", "max=6"},
// 			Max: 6,
// 			Min: 4,
// 			Err: nil,
// 		},
// 		{
// 			Tag: []string{"min=4", "max=0"},
// 			Max: 0,
// 			Min: 4,
// 			Err: ErrMaxLowerMin,
// 		},
// 		{
// 			Tag: []string{"min=0", "max=4"},
// 			Max: 4,
// 			Min: 0,
// 			Err: nil,
// 		},
// 		{
// 			Tag: []string{"min=1", "max=6"},
// 			Max: 6,
// 			Min: 1,
// 			Err: nil,
// 		},
// 	}

// 	for _, tc := range tt {
// 		err := parse(tc.Tag...)
// 		is.Equal(err, tc.Err)
// 		is.Equal(maxValue, tc.Max)
// 		is.Equal(minValue, tc.Min)
// 	}
// }

// type TestStringNoOpts struct {
// 	A string `required:"yes"`
// }

// type TestStringMinMax struct {
// 	A string `required:"yes, min=4, max=15"`
// }

// type TestIntNoOpts struct {
// 	A int16 `required:"yes"`
// }

// type TestIntMinMax struct {
// 	A int32 `required:"yes, min=6,max=12"`
// }

// type TestPtr struct {
// 	A *TestStringNoOpts `required:"yes"`
// }

// type TestSliceNoOpts struct {
// 	A []int `required:"yes"`
// }
// type TestScliceMinMax struct {
// 	A []int `required:"yes, min=3, max=7"`
// }

// type TestMinNA struct {
// 	A string `required:"yes, min=two, max=4"`
// }
// type TestMaxNA struct {
// 	A string `required:"yes, min=2, max=six"`
// }

// type TestMaxLowerMin struct {
// 	A string `required:"yes, min=10, max=8"`
// }

// func TestAll(t *testing.T) {
// 	is := is.New(t)

// 	tt := []struct {
// 		v   interface{}
// 		err error
// 	}{
// 		{
// 			v:   TestMinNA{A: "does not matter"},
// 			err: ErrNotANumber,
// 		},
// 		{
// 			v:   TestMaxNA{A: "does not matter"},
// 			err: ErrNotANumber,
// 		},
// 		{
// 			v:   TestMaxLowerMin{A: "does not matter"},
// 			err: ErrMaxLowerMin,
// 		},
// 		{
// 			v:   TestStringNoOpts{A: "hello world"},
// 			err: nil,
// 		},
// 		{
// 			v:   TestStringNoOpts{A: ""},
// 			err: ErrDefaultFound,
// 		},
// 		{
// 			v:   TestStringMinMax{A: "hello world"},
// 			err: nil,
// 		},
// 		{
// 			v:   TestStringMinMax{A: "hel"},
// 			err: ErrConditionFail,
// 		},
// 		{
// 			v:   TestStringMinMax{A: "hello, this is a word or tow to long"},
// 			err: ErrConditionFail,
// 		},
// 		{
// 			v:   TestIntNoOpts{A: int16(42)},
// 			err: nil,
// 		},
// 		{
// 			v:   TestIntNoOpts{A: int16(0)},
// 			err: ErrDefaultFound,
// 		},
// 		{
// 			v:   TestIntMinMax{A: int32(8)},
// 			err: nil,
// 		},
// 		{
// 			v:   TestIntMinMax{A: int32(5)},
// 			err: ErrConditionFail,
// 		},
// 		{
// 			v:   TestIntMinMax{A: int32(14)},
// 			err: ErrConditionFail,
// 		},
// 		{
// 			v:   TestPtr{A: &TestStringNoOpts{A: "hello"}},
// 			err: nil,
// 		},
// 		{
// 			v:   TestPtr{A: nil},
// 			err: ErrDefaultFound,
// 		},
// 		{
// 			v:   TestSliceNoOpts{A: nil},
// 			err: ErrDefaultFound,
// 		},
// 		{
// 			v:   TestSliceNoOpts{A: []int{24, 42}},
// 			err: nil,
// 		},
// 		{
// 			v:   TestScliceMinMax{A: []int{}},
// 			err: ErrConditionFail,
// 		},
// 		{
// 			v:   TestScliceMinMax{A: []int{1, 2}},
// 			err: ErrConditionFail,
// 		},
// 		{
// 			v:   TestScliceMinMax{A: []int{1, 2, 3, 4}},
// 			err: nil,
// 		},
// 		{
// 			v:   TestScliceMinMax{A: []int{1, 2, 3, 4, 5, 6, 7, 8}},
// 			err: ErrConditionFail,
// 		},
// 		{
// 			v:   0,
// 			err: ErrNoPtr,
// 		},
// 		{
// 			v:   nil,
// 			err: ErrNoPtr,
// 		},
// 	}

// 	for _, t := range tt {
// 		switch t.v.(type) {
// 		case TestStringNoOpts:
// 			v, _ := t.v.(TestStringNoOpts)
// 			err := All(&v)
// 			is.Equal(err, t.err)
// 		case TestStringMinMax:
// 			v, _ := t.v.(TestStringMinMax)
// 			err := All(&v)
// 			is.Equal(err, t.err)
// 		case TestIntNoOpts:
// 			v, _ := t.v.(TestIntNoOpts)
// 			err := All(&v)
// 			is.Equal(err, t.err)
// 		case TestIntMinMax:
// 			v, _ := t.v.(TestIntMinMax)
// 			err := All(&v)
// 			is.Equal(err, t.err)
// 		case TestPtr:
// 			v, _ := t.v.(TestPtr)
// 			err := All(&v)
// 			is.Equal(err, t.err)
// 		case TestSliceNoOpts:
// 			v, _ := t.v.(TestSliceNoOpts)
// 			err := All(&v)
// 			is.Equal(err, t.err)
// 		case TestScliceMinMax:
// 			v, _ := t.v.(TestScliceMinMax)
// 			err := All(&v)
// 			is.Equal(err, t.err)
// 		case TestMinNA:
// 			v, _ := t.v.(TestMinNA)
// 			err := All(&v)
// 			is.Equal(err, t.err)
// 		case TestMaxNA:
// 			v, _ := t.v.(TestMaxNA)
// 			err := All(&v)
// 			is.Equal(err, t.err)
// 		case TestMaxLowerMin:
// 			v, _ := t.v.(TestMaxLowerMin)
// 			err := All(&v)
// 			is.Equal(err, t.err)
// 		default:
// 			err := All(t.v)
// 			is.Equal(err, t.err)
// 		}
// 	}
// }

// type BenchTestMinMax struct {
// 	A string `required:"yes, min=2,max=25"`
// 	B string `required:"yes, min=2,max=25"`
// 	C string `required:"yes, min=2,max=25"`
// 	D string `required:"yes, min=2,max=25"`
// 	E string `required:"yes, min=2,max=25"`
// 	F string `required:"yes, min=2,max=25"`
// 	G string `required:"yes, min=2,max=25"`
// 	H string `required:"yes, min=2,max=25"`
// 	I string `required:"yes, min=2,max=25"`
// 	J string `required:"yes, min=2,max=25"`
// }

// var r error

// func BenchmarkAllMinMax(b *testing.B) {
// 	b.ReportAllocs()
// 	tt := BenchTestMinMax{
// 		A: "hello world",
// 		B: "hello world",
// 		C: "hello world",
// 		D: "hello world",
// 		E: "hello world",
// 		F: "hello world",
// 		G: "hello world",
// 		H: "hello world",
// 		I: "hello world",
// 		J: "hello world",
// 	}
// 	var rr error
// 	for n := 0; n < b.N; n++ {
// 		rr = Match(&tt)
// 	}
// 	r = rr
// }

// type BenchTest struct {
// 	A string   `required:"yes"`
// 	B int16    `required:"yes"`
// 	C []string `required:"yes"`
// 	D string   `required:"yes"`
// 	E float64  `required:"yes"`
// 	F uint32   `required:"yes"`
// 	G []int    `required:"yes"`
// 	H int64    `required:"yes"`
// 	I string   `required:"yes"`
// 	J string   `required:"yes"`
// }

// var result error

// func BenchmarkAllNoOptions(b *testing.B) {
// 	b.ReportAllocs()
// 	tt := BenchTest{
// 		A: "hello world",
// 		B: int16(16),
// 		C: []string{"hello", "world"},
// 		D: "hello world",
// 		E: float64(64.64),
// 		F: uint32(32),
// 		G: []int{1, 2200, 444567, 1337},
// 		H: int64(2000010),
// 		I: "hello world",
// 		J: "J is the last field in the struct",
// 	}

// 	var err error
// 	for n := 0; n < b.N; n++ {
// 		err = All(&tt)
// 	}

// 	result = err
// }

// type BenchTestMix struct {
// 	A string   `required:"yes" m:"jhk"`
// 	B int16    `required:"yes, min=2, max=200"`
// 	C []string `required:"yes, min=1"`
// 	D string   `required:"yes, min=5,max=12"`
// 	E float64  `required:"yes"`
// 	F uint32   `required:"yes"`
// 	G []int    `required:"yes, max=20"`
// 	H int64    `required:"yes"`
// 	I string   `required:"yes"`
// }

// var result3 error

// func BenchmarkAllMix(b *testing.B) {
// 	b.ReportAllocs()
// 	tt := BenchTestMix{
// 		A: "hello world",
// 		B: int16(16),
// 		C: []string{"hello", "world"},
// 		D: "hello world",
// 		E: float64(64.64),
// 		F: uint32(32),
// 		G: []int{1, 2200, 444567, 1337},
// 		H: int64(2000010),
// 		I: "hello world",
// 	}

// 	var err error
// 	for n := 0; n < b.N; n++ {
// 		err = All(&tt)
// 	}

// 	result3 = err
// }
