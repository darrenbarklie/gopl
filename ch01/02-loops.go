package main

import (
	"fmt"
)

func main () {
	
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	j := 6

	for j <= 10 {
		fmt.Println(j)
		j++
	}

	k := 11

	for {
		fmt.Println(k)
	
		k++
			
		if k == 16 {
			fmt.Println("Break")
			break
		}
	}

	fmt.Println("End")

}
