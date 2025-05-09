package main

import "time"

// Function to format time.Time
func formatDate(t time.Time) string {
	return t.Format("Jan 1, 2006 15:04")
}

func add(i int, j int) int {
	return i + j
}

func getTime() string {
	// convert to string
	t := time.Now().Format("2006-01-0215:04:05:999999999")
	return t
}
