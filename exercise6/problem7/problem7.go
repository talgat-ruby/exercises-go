package main

import (
	"fmt"
	"sync"
)

var once sync.Once

// initializeConf initializes the configuration
func initializeConf() {
	fmt.Println("Configuration initialized.")
	// Place actual configuration setup code here
}

// runTasks runs the tasks, ensuring configuration is initialized only once
func runTasks(wg *sync.WaitGroup) {
	defer wg.Done()

	// Ensure configuration is initialized only once
	once.Do(initializeConf)

	// Simulate task processing
	fmt.Println("Task is running.")
}

func main() {
	var wg sync.WaitGroup
	numTasks := 5 // Number of concurrent tasks

	// Start multiple tasks concurrently
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go runTasks(&wg)
	}

	wg.Wait() // Wait for all tasks to complete
	fmt.Println("All tasks completed.")
}
