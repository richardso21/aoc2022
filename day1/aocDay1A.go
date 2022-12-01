package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	max, curr := 0, 0
	for s.Scan() {
		str := s.Text()
		if str == "" {
			if curr > max {
				max = curr
			}
			curr = 0
			continue
		}
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		curr += int(i)
	}
	if curr > max {
		max = curr
	}
	fmt.Println(max)
}
