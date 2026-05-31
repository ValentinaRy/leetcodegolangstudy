package main

type Foo struct {
	chanel1to2 chan bool
	chanel2to3 chan bool
}

func NewFoo() *Foo {
	return &Foo{make(chan bool), make(chan bool)}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	f.chanel1to2 <- true
	close(f.chanel1to2)
}

func (f *Foo) Second(printSecond func()) {
	<-f.chanel1to2
	printSecond()
	f.chanel2to3 <- true
	close(f.chanel2to3)
}

func (f *Foo) Third(printThird func()) {
	<-f.chanel2to3
	printThird()
}
