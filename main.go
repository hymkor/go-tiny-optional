package optional

type Value[T any] struct {
	value T
	ok    bool
}

func Some[T any](v T) Value[T] {
	return Value[T]{value: v, ok: true}
}

func None[T any]() Value[T] {
	return Value[T]{}
}

func (x Value[T]) IfSome(f func(T)) {
	if x.ok {
		f(x.value)
	}
}

func (x Value[T]) Each(f func(T) bool) {
	if x.ok {
		f(x.value)
	}
}

func (x Value[T]) IsNone() bool {
	return !x.ok
}

func (x Value[T]) Match(then func(T), otherwise func()) {
	if x.ok {
		then(x.value)
	} else {
		otherwise()
	}
}
