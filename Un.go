package main

import (
	"fmt"
)

func main() {
	var x, y, z string

	fmt.Println("Masuk Pak eko")
	fmt.Fscanln(&x)
	fmt.Println("Masuk Pak eko")
	fmt.Fscanln(&y)
	fmt.Println("Masuk Pak eko")
	fmt.Fscanln(&z)
	
	ko := [5]int {2, 7, 4, 1, 9}
	i := 1
	for i < 5 {
		j :=i-1
		te := ko[i]
		for te<ko[j] && j>0 {
			ko[j+1] = ko[j]
			j--
		}
		if ko[j]<=te {
			ko[j+1] = te
		} else {
			ko[j+1] = ko[j]
			ko[j] = te
		}
		i++
	}
	for i:=0; i<5; i++ {
		fmt.Print(ko[i], " ")
	}
}