package foo

func Foo() int {
	a := 1 + 1

	if a == 2 {
		panic("oops")
	}

	return a
}
