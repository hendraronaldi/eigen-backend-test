package main

import (
	"fmt"
	"strings"
)

func main() {
	reverseAlphabet()
	longest("Saya sangat senang mengerjakan soal algoritma")
	inputQuery([]string{"xc", "dz", "bbb", "dz"}, []string{"bbb", "ac", "dz"})
	diagonalMatrixDiff([][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}})
}

func reverseAlphabet() {
	s := "NEGIE1"

	var rs string
	cs := strings.Split(s[:len(s)-1], "")
	for i := range cs {
		rs += cs[len(cs)-i-1]
	}
	rs += string(s[len(s)-1])

	fmt.Println("1. Reverse Alphabet")
	fmt.Printf("Input: %s\nOutput: %s\n\n", s, rs)
}

func longest(s string) {
	var rs string

	for _, c := range strings.Split(s, " ") {
		if len(c) > len(rs) {
			rs = fmt.Sprintf("%s: %d", c, len(c))
		}
	}

	fmt.Println("2. Longest")
	fmt.Printf("Input: %s\nOutput: %s\n\n", s, rs)
}

func inputQuery(input, query []string) {
	output := []int{}
	mr := make(map[int]int)

	for i, q := range query {
		if _, ok := mr[i]; !ok {
			mr[i] = 0
		}

		for _, inp := range input {
			if q == inp {
				mr[i] += 1
			}
		}

		output = append(output, mr[i])
	}

	fmt.Println("3. InputQuery")
	fmt.Printf("Input: %v\nQuery: %v\nOutput: %v\n\n", input, query, output)
}

func diagonalMatrixDiff(m [][]int) {
	d1 := 0
	d2 := 0

	if len(m) > 1 {
		for i := range m {
			for j := range m {
				if i == j {
					d1 += m[i][j]
				}

				if i+j == len(m)-1 {
					d2 += m[i][j]
				}
			}
		}
	}

	fmt.Println("4. DiagonalMatrixDiff")
	fmt.Printf("Input: %v\nOutput: %d\n\n", m, (d1 - d2))
}
