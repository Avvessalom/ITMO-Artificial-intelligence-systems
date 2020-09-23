package main

func main() {
	table := parseTable("table.txt")
	//fmt.Println("Поиск в ширину дал ", bfs(table, "Симферополь", "Мурманск"))
	//dfs(createMap(table), "Симферополь", "Мурманск")
	//fmt.Println("Поиск в глубину дал ", printDfs())
	//dls(createMap(table), "Симферополь", "Мурманск",6)
	//fmt.Println("Поиск в глубину c ограничением глубины дал  ", printDls())
	//iddfs(createMap(table),"Симферополь", "Мурманск")
	//fmt.Print(createMapWithWeight(table))
	//fmt.Println("Двунаправленный поиск дал кратчайший путь: " ,biDirectionalSearch(table,"Симферополь", "Мурманск"))
	//fmt.Print(table)
	greedySearch(createMapWithWeight(table), "Симферополь", "Мурманск")
}
