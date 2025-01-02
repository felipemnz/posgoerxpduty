package main

import (
	"fmt"
	"os"
	"time"

	"challenge-stresstest/internal/httpclient"
	"challenge-stresstest/internal/report"

	"github.com/spf13/cobra"
)

func main() {
	var url string
	var requests int
	var concurrency int

	rootCmd := &cobra.Command{
		Use:   "stresstest",
		Short: "Stress test a service",
	}
	rootCmd.Flags().StringVar(&url, "url", "", "Service Url to test")
	rootCmd.Flags().IntVar(&requests, "requests", 0, "Number of requests to make")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 0, "Number of concurrent requests to make")

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if url == "" || requests == 0 || concurrency == 0 {
			fmt.Println("Error: Mandatory flags are missing. Please use --help for more information.")
			cmd.Usage()
			os.Exit(1)
		}

		start := time.Now()
		results := runTest(url, requests, concurrency)
		duration := time.Since(start)

		report.GenerateReport(results, duration)
	}

	rootCmd.Execute()
}

func runTest(url string, requests int, concurrency int) report.Results {
	workers := make(chan struct{}, concurrency)
	results := report.Results{
		TotalRequests:   requests,
		SuccessRequests: 0,
		StatusCodes:     make(map[int]int),
	}

	for i := 0; i < requests; i++ {
		workers <- struct{}{}
		go func() {
			defer func() { <-workers }()
			code, err := httpclient.MakeRequest(url)
			results.StatusCodes[code]++
			if err != nil {
				fmt.Printf("Error making the request: %s\n", err)
				return
			}
			results.SuccessRequests++
		}()
	}

	for i := 0; i < cap(workers); i++ {
		workers <- struct{}{}
	}

	return results
}
