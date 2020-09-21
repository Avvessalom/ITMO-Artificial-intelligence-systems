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
	start string
	end string
	weight int
}

func parseTable(name string) []bone {
	var table = make([]bone, 0, 0)

	file, err := os.Open(name)
	if err != nil{
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
		var _ []bone
		table = append(table, b)
	}
	return table
}

func bfs(table []bone, startName string, endName string) bool {
	queue := createQueue(table, startName)
	var path string = "Кратчайший найденный путь: " + startName + " "
	for queue.Len() != 0 {
		currentCity := queue.PopFront()
		switch v := currentCity.(type) {
		case string:
			path = path + v + " "
		}

		if currentCity == endName {
			fmt.Print("Заебумба")
			fmt.Print(path)
			return true
		} else {
			newQueue := createQueue(table, currentCity)
			for newQueue.Len() != 0 {
				queue.PushBack(newQueue.PopFront())
			}
		}
	}
	fmt.Print("Не заебумба")
	return false
}

func dfs(graph map[string][]string, start string, end string) bool{
	var visited string
	if start == end {
		return true
	}
	if strings.Contains(visited, start) {
		return false
	}
	visited = visited + start
	for i:= 0; i <= len(graph[start]) -1; i++ {
		if !strings.Contains(visited, graph[start][i]) {
			reached := dfs(graph, graph[start][i], end)
			if reached {
				return true
			}
		}
		fmt.Println(graph[start][i])
	}
	return false
}


func createMap(table []bone) map[string][]string {
	graphMap := make(map[string][]string)
	for i := 0; i <= len(table) - 1; i++ {
		var a = make([]string, 0, 0)
		graphMap[table[i].start] = a
		for j := 0; j <= len(table)-1; j++  {
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
	for i := 0; i <= len(table) -1; i++ {
		if table[i].start == starName{
			q.PushBack(table[i].end)
		}
	}
	return q
}