package main

type FooBar struct {
	n       int
	chanFoo chan bool
	chanBar chan bool
}

func NewFooBar(n int) *FooBar {
	return &FooBar{n: n, chanFoo: make(chan bool), chanBar: make(chan bool)}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		printFoo()
		fb.chanBar <- true
		<-fb.chanFoo
	}
	close(fb.chanBar)
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.chanBar
		printBar()
		fb.chanFoo <- true
	}
	close(fb.chanFoo)
}
