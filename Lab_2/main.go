package main

import "fmt"

func main() {
	table := parseTable("table.txt")
	fmt.Println("Поиск в ширину дал ", bfs(table, "С.Петербург", "Брест"))
	dfs(createMap(table), "С.Петербург", "Брест")
	fmt.Println("Поиск в глубину дал ", printDfs())
	//fmt.Printf("%s",createMap(table))
}
