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
	loadRaw, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {

		// Write fallback string to STDOUT and fail.
		fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
		os.Exit(1)
	}

	// Remove surrounding space and split at inner spaces.
	loadStrings := strings.Split(strings.TrimSpace(string(loadRaw)), " ")

	// Build final output string.
	output := fmt.Sprintf("%s, %s, %s", loadStrings[0], loadStrings[1], loadStrings[2])

	fullText = output
	shortText = output

	// Write out gathered information to STDOUT.
	fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
	os.Exit(0)
}
