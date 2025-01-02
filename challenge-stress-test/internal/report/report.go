package report

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Results struct {
	TotalRequests   int
	SuccessRequests int
	StatusCodes     map[int]int
}

func GenerateReport(results Results, duration time.Duration) {
	fmt.Println("----------------------------")
	fmt.Println("------- Test Results -------")
	fmt.Print("----------------------------")
	headerFmt := color.New(color.FgGreen).SprintfFunc()

	tbl := table.New("", "")
	tbl.WithHeaderFormatter(headerFmt)

	tbl.AddRow("Duration:", duration)
	tbl.AddRow("Total Requests:", results.TotalRequests)
	tbl.AddRow("Successful Requests:", results.SuccessRequests)
	tbl.Print()

	fmt.Println()

	fmt.Println("----------------------------")
	fmt.Println("- Status Code Distribution -")
	fmt.Println("----------------------------")

	tbl2 := table.New("Status", "Total")
	tbl2.WithHeaderFormatter(headerFmt)
	for code, count := range results.StatusCodes {
		tbl2.AddRow(code, count)
	}
	tbl2.Print()
}
