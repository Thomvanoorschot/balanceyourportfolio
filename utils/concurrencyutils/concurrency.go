package concurrencyutils

type Tuple2[A any, B any] struct {
	Value A
	Error B
}

func T2[A any, B any](value A, err B) Tuple2[A, B] {
	return Tuple2[A, B]{Value: value, Error: err}
}

func Async[A any](f func() A) <-chan A {
	ch := make(chan A, 1)
	go func() {
		ch <- f()
	}()
	return ch
}

func Async0(f func()) <-chan struct{} {
	ch := make(chan struct{}, 1)
	go func() {
		f()
		ch <- struct{}{}
	}()
	return ch
}

func Async1[A any](f func() A) <-chan A {
	return Async(f)
}

func Async2[A any, B any](f func() (A, B)) <-chan Tuple2[A, B] {
	ch := make(chan Tuple2[A, B], 1)
	go func() {
		ch <- T2(f())
	}()
	return ch
}
