package main

import "fmt"

func main() {
	table := parseTable("table.txt")
	fmt.Printf("%s", dfs(createMap(table), "Витебск", "Брест"))
	fmt.Printf("%s",createMap(table))
}
