package main

import "fmt"

func multable(num int) {
	if num >= 1 {
		for i := 1; i <= num; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("%d * %d = %-2d  ", i, j, i*j)
			}
			fmt.Println("")
		}
	}
}

func main() {
	var num int = 9
	multable(num)
}
