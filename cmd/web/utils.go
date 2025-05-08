package main

import "time"

// Function to format time.Time
func formatDate(t time.Time) string {
	return t.Format("Jan 1, 2006 15:04")
}
