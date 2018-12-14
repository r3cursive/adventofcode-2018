package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("hi")
	dat, err := os.Open("./input")
	check(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	freq := 0

	var results []int
	var operations []int

	for scanner.Scan() {
		text := scanner.Text()
		i, err := strconv.Atoi(text)
		check(err)
		freq += i
		if intInSlice(freq, results) {
			panic(freq)
		}
		results = append(results, freq)
		operations = append(operations, i)

	}
	for true {
		for _, i := range operations {
			freq += i
			if intInSlice(freq, results) {
				panic(freq)
			}
			results = append(results, freq)

		}
	}
	check(scanner.Err())
}
