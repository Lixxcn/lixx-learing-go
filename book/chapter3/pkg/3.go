package pkg

import (
	"fmt"
	"math/big"
	"reflect"
	"unicode/utf8"
)

func F3_1() {
	var a float32

	b := 1.4

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

}

func F3_2() {
	a := 3.14
	b := 3.15
	result := big.NewFloat(a).Cmp(big.NewFloat(b))

	if result < 0 {
		fmt.Println("a < b")
	} else if result > 0 {
		fmt.Println("a > b")
	} else {
		fmt.Println("a == b")
	}
}

func F3_3() {
	var c = '李'
	fmt.Println(c)
	fmt.Printf("%c\r\n", c)
	fmt.Printf("%c\n", c)
	fmt.Println(string(c))
}

func F3_4() {
	q := '\''
	fmt.Printf("%c", q)
	fmt.Println(q)
	l := '\\'
	fmt.Printf("%c", l)
	fmt.Println(l)
}

func F3_5() {
	var s string = "Go 语言字符串"
	fmt.Printf("%v\n", s)
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", reflect.TypeOf(s))
}

func F3_6() {
	s := "字符串"
	len := len(s)
	for i := 0; i < len; i++ {
		fmt.Printf("%X ", s[i])
	}
	println()
}

func F3_7() {
	s := "字符串"

	fmt.Println(utf8.RuneCountInString(s))
	for _, c := range s {
		fmt.Printf("%T %X %c %v %d\n", c, c, c, c, c)
	}
}

func F3_8() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
}

func F3_9() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("a[2]:", a[2])
	a[2] = 100
	fmt.Println(a)
}

func F3_10() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("执行前...", a)
	change(a)
	fmt.Println("执行后...", a)
}

func change(a [5]int) {
	a[2] = 0
	fmt.Println("执行中...", a)
}

func F3_11() {
	s1 := "c:\\Users\\\"Go.pdf\""
	println(s1)
	s2 := `c:\Users\"Go.pdf"`
	println(s2)
}
