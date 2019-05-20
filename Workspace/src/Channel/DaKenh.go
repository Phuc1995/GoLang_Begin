package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const(
	numWorker = 4
	numTasks = 10
)

var wg sync.WaitGroup

func worker(tasks chan string, worker int)  {
	defer wg.Done()


	for{
		//Cho nhan cong viec
		task, ok := <- tasks

		if !ok{
			//Het viec
			fmt.Printf("Nguoi thu %d: Hoan thanh cac cong viec duoc giao!\n",worker)
			return
		}
	}
}

func init(){
	rand.Seed(time.Now().Unix())
}


