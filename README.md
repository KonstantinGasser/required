# required [![GoDoc](https://godoc.org/github.com/KonstantinGasser/required?status.png)](http://godoc.org/github.com/KonstantinGasser/required) ![Go Report Card](https://goreportcard.com/badge/github.com/KonstantinGasser/required) [![codecov](https://codecov.io/gh/KonstantinGasser/required/branch/main/graph/badge.svg)](https://codecov.io/gh/KonstantinGasser/required) ![](https://travis-ci.com/KonstantinGasser/required.svg?branch=main)


Small module helping you validating structs in Go. By adding `required:"yes"` to a struct field, you can ensure that it will not be the default value and satisfies the provided options `min` and `max`.

## Usage

- `required:"yes"`: only checks if value not the default of type
- `required:"yes" min:"2" max"4"`: checks if value not the default and checks for conditions

`min` and `max` works for following types (min and max are both inclusive):
- `string` -> length check
- `intX/uintX/floatX` -> obvious, right?
- `slices` -> length check



## Example
```go
import (
     "encoding/json"
     "log"
     "net/http"
    
     "github.com/KonstantinGasser/required"
)

type User struct {
    Username string `required:"yes" min:"6" max"25"`
    Password string `required:"yes" min:"12"`
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
| Benchmark           | iters  | ns/op      | B/op    | allocs/op     |
|:------------------- |-------:|-----------:| -------:| -------------:|
|BenchmarkRequired-16 | 699244 |1473 ns/op  | 80 B/op | 10 allocs/op  |
|BenchmarkMinMax-16   | 592780 |2012 ns/op  | 80 B/op | 10 allocs/op  |

\*first benchmark is only a struct with `required:"yes"`, the second one with `required:"yes" min:"4" max:"15"`

