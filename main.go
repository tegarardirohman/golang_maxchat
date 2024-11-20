package main

import (
	"fmt"
	"maxchat/internal/router"
	"maxchat/internal/services"
	"net/http"
)

func main() {
	// Memuat data awal dari file
	err := service.LoadDataFromFile("internal/data/data.txt")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	// Setup dan jalankan routes
	router.SetupRoutes()

	// Menjalankan server
	fmt.Println("Server listening on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
