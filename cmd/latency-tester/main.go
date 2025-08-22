package main

import (
	"fmt"
	"net/http"

	"github.com/AmrmDev/latency-tester/internal/adapter/cli"
	"github.com/AmrmDev/latency-tester/internal/infrastructure/net"
	"github.com/AmrmDev/latency-tester/internal/usecase"
)

func main() {
	tester := netinfra.NewHttpLatencyTester()
	usecase := usecase.NewLatencyTesterUseCase(tester)

	http.HandleFunc("/latency", handler.LatencyHandler(usecase))

	fmt.Println("Latency Tester API is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}