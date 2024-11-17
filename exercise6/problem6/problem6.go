 Eldar
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

// initializeConf initializes the configuration
func initializeConf() {
	fmt.Println("Configuration initialized.")
	// Here, we would put the actual initialization code
}

// runTasks runs the tasks and ensures configuration is initialized only once
func runTasks(wg *sync.WaitGroup) {
	defer wg.Done()

	// Initialize configuration only once
	once.Do(initializeConf)

	// Simulate task running
	fmt.Println("Task is running.")
}

func main() {
	var wg sync.WaitGroup

	// Start multiple tasks concurrently
	numTasks := 5
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go runTasks(&wg)
	}

	wg.Wait() // Wait for all tasks to complete
	fmt.Println("All tasks completed.")

package problem6

import (
	"sync"
)

func runTasks(init func()) {
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			//TODO: modify so that load function gets called only once.
			init()
		}()
	}
	wg.Wait()
 main
}
