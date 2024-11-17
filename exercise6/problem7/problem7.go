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

package problem7

import (
	"fmt"
	"math/rand"
	"time"
)

func task() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(
		randomDuration(),
		func() {
			fmt.Println(time.Now().Sub(start))
			t.Reset(randomDuration())
		},
	)
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
 main
}
