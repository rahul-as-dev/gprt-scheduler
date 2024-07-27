package main

import (
	"fmt"
	"sort"
	"time"
)

type goprt struct {
	fun   func()
	order int
}

func main() {
	taskList := make([]goprt, 0)
	taskList = append(taskList, goprt{fun: func() { fmt.Println("order is 1") }, order: 1})
	taskList = append(taskList, goprt{fun: func() { fmt.Println("order is 2") }, order: 2})
	taskList = append(taskList, goprt{fun: func() { fmt.Println("order is 3") }, order: 3})
	taskList = append(taskList, goprt{fun: func() { fmt.Println("order is 4") }, order: 4})
	taskList = append(taskList, goprt{fun: func() { fmt.Println("order is 5") }, order: 5})

	sort.Slice(taskList, func(i, j int) bool {
		return i > j
	})

	for _, task := range taskList {
		go task.fun()
	}
	time.Sleep(1 * time.Second)
	return
}
