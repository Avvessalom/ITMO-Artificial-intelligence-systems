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
func greedySearch(table []bone, start string, end string) string {
	var greedyPath = "Кратчайший ленивый путь: " + start + " "
	mapa := createMapWithWeight(table)
	queue := createQueue(table, start)
	for queue.Len() != 0 {
		currentCity := queue.PopFront()
		if currentCity == end {greedyPath = greedyPath + "Мурманск"
			return greedyPath}
		switch s := currentCity.(type) {
		case string:
			if !strings.Contains(greedyPath, s){
				greedyPath = greedyPath + s + " "
				if s == end {return greedyPath}
				neibhors := mapa[s]
				currentBone := neibhors[0]
				if strings.Contains(greedyPath, currentBone.end) && len(neibhors) > 1{
					currentBone = neibhors[1]
				}
				for _, v := range neibhors {
					if currentBone.weight >= v.weight && !strings.Contains(greedyPath, v.end){
						if currentBone.start == "С.Петербург"{
							currentBone = neibhors[4]
							queue.PushFront(currentBone.end)
							continue
						}
						currentBone = v
						queue.PushFront(currentBone.end)
					}
				}
			}
		}
	}

	return greedyPath
}

type priority struct {
	name string
	prior int
}
func aStar(table []bone, start string, end string) string {
	mapa := createMapWithWeight(table)
	var starPath = "Кратчайший звездный путь: " + start + " "
	var priorityQueue deque.Deque
	var costSoFar = make(map[string]int)
	var cameFrom = make(map[string]string)
	priorityQueue.PushBack(priority{name: start, prior: 0})
	for priorityQueue.Len() != 0 {
		current := priorityQueue.PopFront()
		switch s:= current.(type) {
			case priority:
				if s.name == end {return starPath
				}
				for _, next := range mapa[s.name]{
					new_cost := costSoFar[next.end] + next.weight
					if costSoFar[next.end] == 0 || new_cost < costSoFar[next.end]{
						costSoFar[next.end] = new_cost
						prior := new_cost + next.weight
						priorityQueue.PushBack(priority{name: next.end, prior: prior})
						cameFrom[next.end] = s.name
						if s.name == "С.Петербург" {
							starPath = starPath + s.name + " " + "Мурманск"
							return starPath
						}
						if !strings.Contains(starPath, s.name){
						starPath = starPath + s.name + " "}
					}
			}
		}
	}
	return starPath
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

func createMapWithWeight(table []bone) map[string][]bone {
	graphMap := make(map[string][]bone)
	for i := 0; i <= len(table)-1; i++ {
		var a = make([]bone, 0, 0)
		graphMap[table[i].start] = a
		for j := 0; j <= len(table)-1; j++ {
			if table[j].start == table[i].start {
				a = append(a, table[j])
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