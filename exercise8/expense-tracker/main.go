package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello")
	ctx, cancel := context.WithCancel(context.Background())

	_ := godotenv.Load()

}
