package main

import "fmt"

func main() {
	var n1, count int
	var issimple = true
	v := 10000
	count = 0
	for i := 0; i < v; i++ {
		if i > 2 {
			issimple = true
			for j := 2; j < (i); j++ {
				n1 = i % j
				if n1 == 0 {
					issimple = false
				}
				//fmt.Println(i, j, n1, issimple)
			}
		}
		if issimple {
			count++
			fmt.Printf(" %5d,", i)
		}
		if count > 19 {
			count = 0
			fmt.Println()
		}
	}
	fmt.Println()
}
