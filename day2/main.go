package main

import (
	"bufio"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("hi")
	dat, err := os.Open("./input")
	check(err)
	defer dat.Close()

	//letters = map[string]int{}
	three_times := 0
	two_times := 0
	scanner := bufio.NewScanner(dat)
	alpha := "abcdefghijklmnopqrstuvwxyz"
	for scanner.Scan() {
		text := scanner.Text()
		check(err)
		for _, elem := range alpha {
			if strings.Count(text, string(elem)) == 3 {
				three_times++
				break
			}
		}

		for _, elem := range alpha {
			if strings.Count(text, string(elem)) == 2 {
				two_times++
				break
			}

		}
	}

	fmt.Println(three_times*two_times, " Result")

}
