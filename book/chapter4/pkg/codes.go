package pkg

import "fmt"

func F4_1() {
	a := 65536
	pa := &a
	b := 65536
	p := &b
	fmt.Printf("%d\n%v\n%x\n%v\n%T\n", a, pa, p, p, p)

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d %v\n", arr[i], &arr[i])
	}
}

func F4_2() {
	s := []int{1, 2, 3}
	fmt.Printf("%T %v 长度：%d, 容量：%d\n", s, s, len(s), cap(s))
	s1 := make([]int, 5, 10)
	fmt.Printf("%T %v 长度：%d, 容量：%d\n", s1, s1, len(s1), cap(s1))
}

func F4_3() {
	s1 := make([]int, 5, 10)
	for i := 0; i < len(s1); i++ {
		s1[i] = 100
	}
	for _, e := range s1 {
		fmt.Println(e)
	}
}

func F4_4() {
	s1 := make([]int, 5, 10)
	for i := 0; i < len(s1); i++ {
		s1[i] = i
	}
	fmt.Printf("%p %v 长度：%d, 容量：%d\n", s1, s1, len(s1), cap(s1))
	s1 = append(s1, 5, 6, 7, 8, 9)
	fmt.Printf("%p %v 长度：%d, 容量：%d\n", s1, s1, len(s1), cap(s1))
	s1 = append(s1, 10)
	fmt.Printf("%p %v 长度：%d, 容量：%d\n", s1, s1, len(s1), cap(s1))
}

func F4_5() {
	s1 := make([]int, 5, 10)
	fmt.Println(s1)
	_ = append(s1, 5, 6, 7, 8, 9)
	fmt.Println(s1)
}

func F4_7() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%p %v 长度：%d, 容量：%d\n", a, a, len(a), cap(a))
	s1 := a[:]
	fmt.Printf("%p %v 长度：%d, 容量：%d\n", s1, s1, len(s1), cap(s1))
	s2 := a[2:4]
	fmt.Printf("%p %v 长度：%d, 容量：%d\n", s2, s2, len(s2), cap(s2))
}

func F4_8() {
	s := make([]int, 4)
	fmt.Printf("%p\n", s)

	s1 := s[1:]
	fmt.Printf("%p\n", s1)
}

func F4_9() {
	s := []int{1, 2, 3, 4, 5}
	s1 := s[1:]
	s1[0] = -2
	fmt.Printf("%T, %v\n", s, s)

	for i, num := range s1 {
		fmt.Printf("%d %d\n", i, num)
	}

	s2 := []int{1}
	fmt.Println(s2)
	fmt.Println(s2[1:])
	fmt.Println(s2[0])
	s3 := s2
	fmt.Printf("%p %p\n", s2, s3)
}

func F4_10() {
	charCountMap := make(map[string]int)

	charCountMap["a"] = 1
	charCountMap["b"] = 2

	for k, v := range charCountMap {
		fmt.Printf("k:%s v:%d\n", k, v)
	}
}

func F4_12() {
	charCountMap := make(map[string]int)

	charCountMap["a"] = 1
	charCountMap["b"] = 2

	for k, v := range charCountMap {
		fmt.Printf("k:%s v:%d\n", k, v)
	}
	fmt.Println(charCountMap["c"])
}

func F4_13() {
	var charCountMap map[string]int
	delete(charCountMap, "a")
}

func F4_14() {
	count := 5
	msgChannel := make(chan int, count)

	for i := 0; i < count; i++ {
		msgChannel <- i
	}
	for i := 0; i < count; i++ {
		fmt.Println(<-msgChannel)
	}
}

type Persion struct {
	Name  string
	Age   int
	Email string
	sex   string
}

func F4_15() {
	persion := Persion{
		Name:  "张三",
		Age:   18,
		Email: "zhangsan@163.com",
	}
	fmt.Println(persion)

	var persion1 = Persion{}
	fmt.Println(persion1)

	persion1.Name = "李四"
	persion1.Age = 20
	persion1.Email = "lisi@163.com"

	fmt.Println(persion1)

	var persion2 = Persion{}
	var persion3 = new(Persion)
	fmt.Printf("%T %T\n", persion2, persion3)
}
