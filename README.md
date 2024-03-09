[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/go-tiny-optional.svg)](https://pkg.go.dev/github.com/hymkor/go-minimum-optional)

go-tiny-optional
===================

This package has only two constructors (Some, None) and four methods (IfSome, IsNone, Each, Match).

It requires Go 1.18 or later.

```example.go
package main

import (
    "github.com/hymkor/go-tiny-optional"
)

func test(x optional.Value[int]) {
    x.IfSome(func(v int) {
        println("   IfSome: it has a value:", v)
    })

    for v := range x.Each {
        println("   for-range(v1.22 X:rangefunc): it has a value:", v)
    }

    if x.IsNone() {
        println("   IsNone: it does not have a value")
    }

    x.Match(func(v int) {
        println("   Match: it has a value:", v)
    }, func() {
        println("   Match: it does not hava a value")
    })

    println()
}

func main() {
    println("None[int]")
    test(optional.None[int]())

    println("Some(4)")
    test(optional.Some(4))
}
```

**go run example.go**

```env GOEXPERIMENT=rangefunc go run example.go|
None[int]
   IsNone: it does not have a value
   Match: it does not hava a value

Some(4)
   IfSome: it has a value: 4
   for-range(v1.22 X:rangefunc): it has a value: 4
   Match: it has a value: 4

```

Comparison with [go-minimal-optional]
-------------------------------------

### Type

#### go-minimal-optional

```
type Value[T any] []T
```

#### go-tiny-optional

```
type Value[T any] struct {
    value T
    ok    bool
}
```

### Speed

#### go-minimal-optional

```
$ go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: github.com/hymkor/go-minimal-optional
cpu: Intel(R) Core(TM) i5-6500T CPU @ 2.50GHz
BenchmarkIfSome-4       656701584                1.790 ns/op           0 B/op          0 allocs/op
BenchmarkIfNone-4       1000000000               0.7019 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/hymkor/go-minimal-optional   2.315s
```

#### go-tiny-optional

```
$ go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: github.com/hymkor/go-tiny-optional
cpu: Intel(R) Core(TM) i5-6500T CPU @ 2.50GHz
BenchmarkIfSome-4       675800485                1.787 ns/op           0 B/op          0 allocs/op
BenchmarkIfNone-4       1000000000               0.7008 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/hymkor/go-tiny-optional      2.325s
```

[go-minimal-optional]: https://github.com/hymkor/go-minimal-optional
