package main

import (
	"fmt"
	"os"
	"strconv"
)

type Day struct {
	Part1 (func(string) int)
	Part2 (func(string) int)
}

func (d *Day) Run(inputFile string) {
	fmt.Printf(
		"\033[3mPart 1\033[0m\n  >>\033[1;34m%d\033[0m<<\n",
		d.Part1(inputFile),
	)
	fmt.Printf(
		"\033[3mPart 2\033[0m\n  >>\033[1;34m%d\033[0m<<\n",
		d.Part2(inputFile),
	)
}

func main() {
	days := []Day{
		day01,
		day02,
	}
	if len(os.Args) < 2 {
		days[len(days)-1].Run(fmt.Sprintf("Day%02d-input", len(days)))
		os.Exit(0)
	}
	switch os.Args[1] {
	case "-h", "--help":
		fmt.Println(`
These are my solutions to Advent of code 2024.
To get the answers for a particular day, please run
    [1madvent24 {DAY#}[0m
        `)
	default:
		dayNum, err := strconv.Atoi(os.Args[1])
		check(err)
		days[dayNum-1].Run(fmt.Sprintf("Day%02d-input", dayNum+1))
	}
}
