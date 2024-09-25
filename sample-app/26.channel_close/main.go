package main

import "fmt"

type WorkerStruct struct {
	In   chan string
	Done chan bool
}

func doWork(i int, worker *WorkerStruct) {
	for {
		n := <-worker.In
		fmt.Printf("index=%d,channel=%s \n", i, n)
		worker.Done <- true
	}
}

func createWorker(i int) WorkerStruct {
	// c := make(chan string)
	// var worker WorkerStruct
	// worker.In = make(chan string)
	worker := WorkerStruct{
		In:   make(chan string),
		Done: make(chan bool),
	}
	go doWork(i, &worker)
	return worker
}

func chanDemo() {
	fmt.Println("more work")
	var workerList [10]WorkerStruct
	for i := 0; i < 10; i++ {
		workerList[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workerList[i].In <- fmt.Sprintf("hello,%d", i+'a')
		<-workerList[i].Done
	}
	for i := 0; i < 10; i++ {
		workerList[i].In <- fmt.Sprintf("Hello,%d", i+'A')
		<-workerList[i].Done
	}
}

func main() {
	chanDemo()
}
