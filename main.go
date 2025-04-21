package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cli/go-gh/v2/pkg/api"
)

type IssueData struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Labels    []string `json:"labels,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
}

func main() {
	// Define flags
	csvFile := flag.String("file", "", "Path to CSV file containing issue data")
	repo := flag.String("repo", "", "Repository in format owner/repo")
	hasHeader := flag.Bool("header", true, "Whether the CSV file has a header row")
	flag.Parse()

	// Validate flags
	if *csvFile == "" || *repo == "" {
		fmt.Println("Error: Both --file and --repo flags are required")
		flag.Usage()
		os.Exit(1)
	}

	// Open CSV file
	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create CSV reader
	reader := csv.NewReader(file)

	// Skip header if present
	if *hasHeader {
		_, err = reader.Read()
		if err != nil {
			fmt.Printf("Error reading header: %v\n", err)
			os.Exit(1)
		}
	}

	// Create GitHub API client
	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Printf("Error creating GitHub client: %v\n", err)
		os.Exit(1)
	}

	// Process each row
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading row: %v\n", err)
			continue
		}

		// Create issue data from CSV row
		// Expected format: title,body,labels,assignees
		issue := IssueData{
			Title: record[0],
			Body:  record[1],
		}

		// Handle optional labels
		if len(record) > 2 && record[2] != "" {
			issue.Labels = strings.Split(record[2], ";")
		}

		// Handle optional assignees
		if len(record) > 3 && record[3] != "" {
			issue.Assignees = strings.Split(record[3], ";")
		}

		// Create the issue
		var response struct {
			Number int    `json:"number"`
			URL    string `json:"html_url"`
		}

		path := fmt.Sprintf("repos/%s/issues", *repo)
		jsonData, err := json.Marshal(issue)
		if err != nil {
			fmt.Printf("Error marshaling issue data: %v\n", err)
			continue
		}
		err = client.Post(path, strings.NewReader(string(jsonData)), &response)
		if err != nil {
			fmt.Printf("Error creating issue '%s': %v\n", issue.Title, err)
			continue
		}

		fmt.Printf("Created issue #%d: %s\n", response.Number, response.URL)
	}
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
