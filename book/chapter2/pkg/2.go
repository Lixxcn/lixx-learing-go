package chapter2

import "fmt"

func F2_4() {
	const (
		a = 10
		b
		c
	)
	fmt.Println(a, b, c)
}

func F2_5() {
	const (
		Sunday    = iota
		Monday    = iota
		Tuesday   = iota
		Wednesday = iota
		Thursday  = iota
		Friday    = iota
		Saturday  = iota
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

func F2_6() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

func F2_7() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday = "x"
		Thursday  = iota
		Friday    = iota
		Saturday  = iota
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

}

func F2_8() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday = "x"
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

}

func F2_9() {
	const (
		_ = iota
		Sunday
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
