package main

import "tracker/internal/api"

func main() {
	err := api.StartServer()
	if err != nil {
		return
	}
}
