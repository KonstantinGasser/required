# required [![GoDoc](https://godoc.org/github.com/KonstantinGasser/required?status.png)](http://godoc.org/github.com/KonstantinGasser/required) ![Go Report Card](https://goreportcard.com/badge/github.com/KonstantinGasser/required)

Small module helping you validating structs in Go. By adding `required:"yes"` to a struct field, you can ensure that it will not be the default value.

## Usage

- `required:"yes"`: only checks if value not the default of type
- `required:"yes, min=2, max=4"`: checks if value not the default and checks for conditions

`min` and `max` works for following types:
- `string` -> length check
- `intX/uintX/floatX` -> obvious, right?
- `slices` -> length check

## Gotchas
- Field names must be exported else validation will not find any fields and returns no error
- interface must be exported else validation will fail with `required.ErrRequiredFailed`

## Excluded types where tag has no effect
- `bool`: will be ignored
- `struct`: if struct is the empty struct validation will fail

## Example
```go
import "github.com/KonstantinGasser/required"

type User struct {
    Username string `required:"yes, min=6, max=25"`
    Password string `required:"yes, min=12"`
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        log.Fatal(err)
    }
    // will fail if:
    // - len(user.Username) < 6 or > 25
    // - len(user.Password) < 12
    if err := required.All(&user); err != nil {
        log.Fatalf("Not all fields in the User struct satisfy the tag conditions: %v", err)
    }

    ...

}
```

## Benchmarks
| Benchmark                | ns/op        | B/op      | allocs/op      |
|:------------------------ | ------------:| ---------:| --------------:|
|BenchmarkAllMinMax-16      | 4224 ns/op	  | 1360 B/op |  50 allocs/op  |
|BenchmarkAllNoOptions-16   | 1340 ns/op	  | 240 B/op  | 20 allocs/op   |
|BenchmarkAllMix-16         | 2151 ns/op	  | 544 B/op  | 27 allocs/op   |

\*first benchmark is only a struct with `required:"yes"`, the second one with `required:"yes,min=4,max=15"`
