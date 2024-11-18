package main

import (
	"fmt"

	pkg1 "lixx.cn/go-book-1/c14/package1"
)

func main() {
	m := map[string]int{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 3,
		"k5": 3,
		"k6": 3,
	}

	fmt.Println(m["k1"])
	fmt.Println(m["k3"])

	_, ok := m["k3"]
	if !ok {
		fmt.Println("k3 does not exist")
	}
	fmt.Println(m)

	// delete(m, "k3")
	// fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, " ", v)
	}

	var s1 []string
	for k, _ := range m {
		s1 = append(s1, k)
	}
	fmt.Println(s1)

	pkg1.F1()
}
