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
		switch p2 {
		case 'X':
			points += getChoice(p1, 2)
		case 'Y':
			points += getChoice(p1, 0) + 3
		case 'Z':
			points += getChoice(p1, 1) + 6
		}
	}
	fmt.Println(points)
}

func getChoice(p1 rune, modifier int) int {
	// modifier = 2 to lose, 0 to draw, 1 to win
	switch p1 {
	case 'A':
		return (0+modifier)%3 + 1
	case 'B':
		return (1+modifier)%3 + 1
	case 'C':
		return (2+modifier)%3 + 1
	}
	return 0
}
