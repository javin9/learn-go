package main

import (
	"fmt"
	"sync"
)

type WorkerStruct struct {
	In chan string
	Wg *sync.WaitGroup
}

func doWork(i int, worker *WorkerStruct) {
	n := <-worker.In
	fmt.Printf("index=%d,channel=%s \n", i, n)
	defer worker.Wg.Done()
}

func createWorker(i int, wg *sync.WaitGroup) WorkerStruct {
	worker := WorkerStruct{
		In: make(chan string),
		Wg: wg,
	}
	go doWork(i, &worker)
	return worker
}

func chanDemo() {
	fmt.Println("more work")
	var wg sync.WaitGroup

	var workerList [10]WorkerStruct
	for i := 0; i < 10; i++ {
		workerList[i] = createWorker(i, &wg)
		wg.Add(1)
	}

	for i := 0; i < 10; i++ {
		workerList[i].In <- fmt.Sprintf("hello,%d", i+'a')
		// <-workerList[i].Done
	}

	wg.Wait()
	fmt.Println("结束啦")

}

func main() {
	fmt.Println("使用Channel来等待goroutine的结束")
	chanDemo()
}
