package main

import (
	"fmt"
	"os"
	"strings"

	"io/ioutil"
)

func main() {

	// Set both display texts to 'error' by default.
	var fullText string = "error"
	var shortText string = "error"

	// Read current load average information from kernel
	// pseudo-file-system mounted at /proc.
	load, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {

		// Write fallback string to STDOUT and fail.
		fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
		os.Exit(1)
	}

	// Remove surrounding space and split at inner spaces.
	loadRaw := strings.Split(strings.TrimSpace(string(load)), " ")

	// Build final output string.
	output := fmt.Sprintf("%s, %s, %s", loadRaw[0], loadRaw[1], loadRaw[2])

	fullText = output
	shortText = output

	// Write out gathered information to STDOUT.
	fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
	os.Exit(0)
}
