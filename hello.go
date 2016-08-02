package main

import "fmt"

func main() {
	count := set_count()
	for i := 0; i < 10; i++ {
		count()
	}
	mymap := map[int]int{0: 42}
	mymap[0] = 42
	entry, ok := mymap[0]
	fmt.Println(entry, ok)
}

func set_count() func() {
	var cnt int = 0
	return func() {
		cnt++
		fmt.Println(cnt)
	}
}
