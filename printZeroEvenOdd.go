package main

type ZeroEvenOdd struct {
	n       int
	canZero chan bool
	canEven chan int
	canOdd  chan int
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:       n,
		canZero: make(chan bool, 1),
		canEven: make(chan int, 1),
		canOdd:  make(chan int, 1)}
	zeo.canZero <- true
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		<-z.canZero
		printNumber(0)
		if i%2 == 0 {
			z.canEven <- i
		} else {
			z.canOdd <- i
		}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	if z.n < 2 {
		return
	}
	for i := range z.canEven {
		printNumber(i)
		z.canZero <- true
		if i >= z.n-1 {
			break
		}
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for i := range z.canOdd {
		printNumber(i)
		z.canZero <- true
		if i >= z.n-1 {
			break
		}
	}
}
