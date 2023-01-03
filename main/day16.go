package main

import (
	"fmt"
	"regexp"
	"strings"
)

type pipe struct {
	valve   string
	rate    int
	tunnels []string
}

func main() {
	lines := readLines("./resource/day16.txt")

	valves := []pipe{}
	for _, line := range lines {
		p := pipe{}

		str := strings.Split(line, " ")
		// fmt.Println("valve is: ", str[1])
		p.valve = str[1]

		rg := regexp.MustCompile(`[\d]+`)
		r := rg.FindString(str[4])
		// fmt.Println("rate is: ", r)
		p.rate = toInt(r)

		rg1 := regexp.MustCompile(`tunnels?`)
		rg2 := regexp.MustCompile(`valves?`)
		rg3 := regexp.MustCompile(`leads?`)
		tl1 := rg1.ReplaceAllString(line, "tunnels")
		tl2 := rg2.ReplaceAllString(tl1, "valves")
		tl3 := rg3.ReplaceAllString(tl2, "lead")

		tl := strings.Split(tl3, "tunnels lead to valves ")
		tls := strings.Split(tl[1], " ")
		// fmt.Println("tunnels are: ", tls)
		p.tunnels = tls
		valves = append(valves, p)
	}
	fmt.Println("list of valves: ", valves)
}
