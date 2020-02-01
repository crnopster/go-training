package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

const one = 1

func find(numbers []int) []string {
	var s []string

	for _, num := range numbers {
		if num < one {
			s = append(s, fmt.Sprintf("%v - number must be positive", num))
			continue
		}

		for x := 0; x*x < num+1; x++ {
			for y := 0; y*y < num+1; y++ {
				for z := 0; z*z < num+1; z++ {
					if x*x+y*y+z*z == num {
						s = append(s, fmt.Sprintf("for %v : x = %v, y = %v, z = %v", num, x, y, z))
					}
				}
			}
		}
	}

	return s
}

func check(numbers []string) []int {
	intnumbers := make([]int, 0)

	for _, number := range numbers {
		nint, err := strconv.Atoi(number)
		if err != nil {
			continue
		}

		intnumbers = append(intnumbers, nint)
	}

	return intnumbers
}

func input() []string {
	var n []string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Input number: ")
		scanner.Scan()
		text := scanner.Text()
		n = append(n, text)

		if text == "" {
			n = n[0 : len(n)-1]

			break
		}
	}

	return n
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var numbers []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	return numbers
}

func main() {
	filename := flag.String("f", "", "for taking numbers from file")

	flag.Parse()

	args := os.Args[1:]

	switch {
	case len(*filename) != 0:
		fmt.Println(find(check(readFile(*filename))))
	case len(args) == 0:
		fmt.Println(find(check(input())))
	default:
		fmt.Println(find(check(args)))
	}
}
