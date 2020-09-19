package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bone struct {
	start string
	end string
	weight int
}

func parseTable() {
	var table = make([]bone, 0, 0)

	file, err := os.Open("table.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		//fmt.Printf("%s\n", s.Text())
		str := strings.Split(s.Text(), " ")
		distance, err := strconv.Atoi(str[2])
		if err != nil {
			log.Fatal(err)
		}
		b := bone{start: str[0], end: str[1], weight: distance}
		var _ []bone
		table = append(table, b)
	}
	fmt.Print(table)
}
