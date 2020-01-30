package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const one = 1

func find(num int) []string {
	s := make([]string, 0)

	var x, y, z int

	if num < one {
		s = append(s, fmt.Sprintf("%v - number must be positive", num))
		return s
	}

loop0:
	for {
		y = 0
	loop1:
		for {
			z = 0
		loop2:
			for {
				if x*x+y*y+z*z == num {
					s = append(s, fmt.Sprintf("for %v : x = %v, y = %v, z = %v\n", num, x, y, z))
				}

				if z*z > num {
					break loop2
				}
				z++
			}
			if y*y > num {
				break loop1
			}
			y++
		}
		if x*x > num {
			break loop0
		}
		x++
	}

	return s
}

func check(numbers []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for _, number := range numbers {
		nint, err := strconv.Atoi(number)
		if err != nil {
			log.Println("args must be intengers\n", err)
			continue
		}

		log.Printf("number %v checked\n", nint)
		fmt.Printf("Find for %v? : ", nint)
		scanner.Scan()
		text := scanner.Text()

		if text == "Yes" || text == "Y" || text == "yes" || text == "y" {
			fmt.Println(find(nint))
		}
	}

	fmt.Print("Do you want to try another numbers? ")
	repeat()
}

func repeat() {
	var n []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)

	if text == "Yes" || text == "Y" || text == "yes" || text == "y" {
		for {
			fmt.Print("Input number: ")
			scanner.Scan()
			text = scanner.Text()
			n = append(n, text)

			if text == "" {
				n = n[0 : len(n)-1]
				check(n)

				break
			}
		}
	}
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
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
	filename := flag.String("f", "", "for taking number from file")

	flag.Parse()

	args := os.Args[1:]

	switch {
	case len(*filename) != 0:
		f := readFile(*filename)
		check(f)
	case len(args) == 0:
		fmt.Println("Input manually? :")
		repeat()

	default:
		check(args)
	}
}
