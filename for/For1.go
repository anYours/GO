package main

import "fmt"

func test1() {
	var i int
	for {
		i++
		if i == 100 {
			break
		}
		fmt.Println(i)
	}
}
