package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	points := 0
	for s.Scan() {
		str := s.Text()
		a := strings.Split(str, " ")
		p1, p2 := rune(a[0][0]), rune(a[1][0])
		switch p1 {
		case 'A':
			if p2 == 'X' {
				points += 3
			} else if p2 == 'Y' {
				points += 6
			}
		case 'B':
			if p2 == 'Y' {
				points += 3
			} else if p2 == 'Z' {
				points += 6
			}
		case 'C':
			if p2 == 'Z' {
				points += 3
			} else if p2 == 'X' {
				points += 6
			}
		}
		points += (int(p2) - 'W')
	}
	fmt.Println(points)
}
