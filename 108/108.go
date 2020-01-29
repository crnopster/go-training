package main

import (
	"fmt"
	"log"
	"math"
)

func find(n float64) string {
	x := 0.0
	if n < 0 {
		return "number cant be negative"
	} else if n > math.Pow(2, x) {
		for {
			x = x + 1
			log.Println(math.Pow(2, x))
			if math.Pow(2, x) > n {
				break
			}
		}
		//		for {
		//			x = x - 0.1
		//			log.Println(math.Pow(2, x))
		//			if math.Pow(2, x) < n {
		//				x = x + 0.1
		//				break
		//			}
		//		}
		//		for {
		//			x = x - 0.01
		//			log.Println(math.Pow(2, x))
		//			if math.Pow(2, x) < n {
		//				x = x + 0.01
		//				break
		//			}
		//		}
	}
	return fmt.Sprintf("The smallest number for 2^x > n is %v", x)
}

func main() {
	fmt.Println(find(-1))
	fmt.Println(find(129))
	fmt.Println(find(65537))
}
