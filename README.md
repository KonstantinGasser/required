# required [![GoDoc](https://godoc.org/github.com/KonstantinGasser/required?status.png)](http://godoc.org/github.com/KonstantinGasser/required) ![Go Report Card](https://goreportcard.com/badge/github.com/KonstantinGasser/required) [![codecov](https://codecov.io/gh/KonstantinGasser/required/branch/main/graph/badge.svg)](https://codecov.io/gh/KonstantinGasser/required) ![](https://travis-ci.com/KonstantinGasser/required.svg?branch=main)


Small module helping you validating structs in Go. By adding `required:"yes"` to a struct field, you can ensure that it will not be the default value and if required add conditions.

## Usage

- `required:"yes"`: only checks if value not the default of type
- `required:"yes, min=2, max=4"`: checks if value not the default and checks for conditions

`min` and `max` works for following types:
- `string` -> length check
- `intX/uintX/floatX` -> obvious, right?
- `slices` -> length check



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
    if err := required.Atomic(&user); err != nil {
        log.Fatalf("Not all fields in the User struct satisfy the tag conditions: %v", err)
    }

    ...

}
```

## Benchmarks
| Benchmark                | iters  | ns/op      | B/op      | allocs/op     |
|:------------------------ |-------:|-----------:| ---------:| -------------:|
|BenchmarkAllMinMax-16     | 288800 |3756 ns/op  | 1200 B/op | 40 allocs/op  |
|BenchmarkAllNoOptions-16  | 934557 |1217 ns/op  | 240 B/op  | 20 allocs/op  |
|BenchmarkAllMix-16        | 687088 |1698 ns/op  | 456 B/op  | 20 allocs/op  |

\*first benchmark is only a struct with `required:"yes"`, the second one with `required:"yes,min=4,max=15"`

