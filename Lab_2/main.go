package main

import "fmt"

func main() {
	table := parseTable("table.txt")
	fmt.Println("Поиск в ширину дал ", bfs(table, "Симферополь", "Мурманск"))
	dfs(createMap(table), "Симферополь", "Мурманск")
	fmt.Println("Поиск в глубину дал ", printDfs())
	dls(createMap(table), "Симферополь", "Мурманск",5)
	fmt.Println("Поиск в глубину c ограничением глубины дал  ", printDls())
	//fmt.Printf("%s",createMap(table))
}
