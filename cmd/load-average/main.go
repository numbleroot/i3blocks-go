package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	// Set display texts to defaults.
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

	// Depending on length of display text, construct
	// final output string.
	fullText = fmt.Sprintf("%s<span foreground=\"#999999\">-</span>%s<span foreground=\"#999999\">-</span>%s", loadStrings[0], loadStrings[1], loadStrings[2])
	shortText = fmt.Sprintf("%s", loadStrings[0])

	// Write out gathered information to STDOUT.
	fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
	os.Exit(0)
}
