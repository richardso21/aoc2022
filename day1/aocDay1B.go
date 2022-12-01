package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	h := make([]int, 0)
	var curr int = 0
	for s.Scan() {
		str := s.Text()
		if str == "" {
			h = append(h, curr)
			curr = 0
			continue
		}
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		curr += int(i)
	}
	sort.Ints(h)
	max3 := 0
	for i := 0; i < 3; i++ {
		max3 += h[len(h)-i-1]
	}
	fmt.Println(max3)
}
