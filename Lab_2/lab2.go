package main

import (
	"bufio"
	"fmt"
	"github.com/gammazero/deque"
	"log"
	"os"
	"strconv"
	"strings"
)

type bone struct {
	start  string
	end    string
	weight int
}

func parseTable(name string) []bone {
	var table = make([]bone, 0, 0)

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		str := strings.Split(s.Text(), " ")
		distance, err := strconv.Atoi(str[2])
		if err != nil {
			log.Fatal(err)
		}
		b := bone{start: str[0], end: str[1], weight: distance}
		c := bone{start: str[1], end: str[0], weight: distance}
		var _ []bone
		table = append(table, b)
		table = append(table, c)
	}
	return table
}

func bfs(table []bone, startName string, endName string) string {
	queue := createQueue(table, startName)
	var path string = "кратчайший путь: " + startName + " "
	for queue.Len() != 0 {
		currentCity := queue.PopFront()
		switch v := currentCity.(type) {
		case string:
			if !strings.Contains(path, v) {
				path = path + v + " "
				if currentCity == endName {
					return path
				} else {
					newQueue := createQueue(table, currentCity)
					for newQueue.Len() != 0 {
						queue.PushBack(newQueue.PopFront())
					}
				}
			}
		}
	}
	return path
}

var path = "кратчайший путь: "
var visited string

func dfs(graph map[string][]string, start string, end string) bool {
	path = path + start + " "
	if start == end {
		return true
	}
	if strings.Contains(visited, start) {
		return false
	}
	visited = visited + start
	for i := 0; i <= len(graph[start])-1; i++ {
		if !strings.Contains(visited, graph[start][i]) {
			reached := dfs(graph, graph[start][i], end)
			if reached {
				return true
			}
		}
	}
	return false
}
var pathDLS = "кратчайший путь: "
var visitedDLS string

func dls (graph map[string][]string, start string, end string, limit int) bool {
	if strings.Count(pathDLS, start) == 0 {
		pathDLS = pathDLS + start + " "
	}
	var pathDLS = "кратчайший путь: "
	if start == end {return true}
	if limit >= 0 {
		if strings.Contains(visitedDLS, start) {return false}
			visitedDLS = visitedDLS + start
		for i:=0; i <= len(graph[start])-1; i++ {
			if !strings.Contains(visitedDLS, graph[start][i]) && strings.Count(pathDLS, " ")-3 <= limit{
				reached := dls(graph, graph[start][i], end, limit -1)
				if reached {return true}
			}
		}
	}
	return false
}

func iddfs(graph map[string][]string, start string, end string) bool {
	limit := 0
	pathDLS = "кратчайший путь: "
	visitedDLS = " "
	if dls(graph,start,end,limit) {return true}
	if start == end {return true}
	for i := 0; i <= 500; i++ {
		pathDLS = "кратчайший путь: "
		visitedDLS = " "
		if dls(graph,start,end,i) {
			fmt.Println("Поиск в глубину c итеративным увеличением глубины дал  ",pathDLS, " глубина: ", i)
			return true
		}
	}
	return false
}

func biDirectionalSearch(table []bone, start string, end string) string {
	visited := ""
	queueStart := createQueue(table, start)
	queueEnd := createQueue(table, end)
	var pathFront string =  start + " "
	var pathBack string = end + " "
	for queueStart.Len() != 0 {
		currentCityFront := queueStart.PopFront()
		switch v := currentCityFront.(type) {
		case string:
			if strings.Contains(visited, v) {return pathFront + pathBack}
			if !strings.Contains(pathFront, v) {
				pathFront = pathFront + v + " "
				visited = visited + v
				if currentCityFront == end {
					return pathFront
				} else {
					newQueue := createQueue(table, currentCityFront)
					for newQueue.Len() != 0 {
						queueStart.PushBack(newQueue.PopFront())
					}
				}
			}
		}
		currentCityBack := queueEnd.PopFront()
		switch v := currentCityBack.(type) {
		case string:
			if strings.Contains(visited, v) {return pathFront + pathBack}
			if !strings.Contains(pathBack, v) {
				pathBack = v + " " + pathBack
				visited = visited + v
			if currentCityBack == start {
				return pathBack
			} else {
					newQueue := createQueue(table, currentCityBack)
					for newQueue.Len() != 0 {
						queueEnd.PushBack(newQueue.PopFront())
					}
				}
			}
		}
	}
	return pathFront

}

func printDfs() string {
	return path
}
func printDls() string {
	return pathDLS
}

func createMap(table []bone) map[string][]string {
	graphMap := make(map[string][]string)
	for i := 0; i <= len(table)-1; i++ {
		var a = make([]string, 0, 0)
		graphMap[table[i].start] = a
		for j := 0; j <= len(table)-1; j++ {
			if table[j].start == table[i].start {
				a = append(a, table[j].end)
				graphMap[table[i].start] = a
			}
		}
	}
	return graphMap
}

func createQueue(table []bone, starName interface{}) deque.Deque {
	var q deque.Deque
	for i := 0; i <= len(table)-1; i++ {
		if table[i].start == starName {
			q.PushBack(table[i].end)
		}
	}
	return q
}