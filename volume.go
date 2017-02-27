package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var volRegex = regexp.MustCompile(`\[(\d{1,3})\%\]`)

func main() {
	vol, err := volume()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get volume: %v", err)
		os.Exit(1)
	}

	// Set display texts to defaults.
	full := "unknown"
	short := "unknown"

	if vol >= 0 {
		full = fmt.Sprintf("%d%%", vol)
		short = fmt.Sprintf("%d%%", vol)
	}

	if vol == 0 {
		full = fmt.Sprintf(" %s", full)
		short = fmt.Sprintf("  %s", short)
	}
	if 49 > vol && vol > 0 {
		full = fmt.Sprintf(" %s", full)
		short = fmt.Sprintf(" %s", short)
	}
	if vol >= 50 {
		full = fmt.Sprintf(" %s", full)
		short = fmt.Sprintf(" %s", short)
	}

	// Write out gathered information to STDOUT.
	fmt.Fprintf(os.Stdout, "%s\n%s\n", full, short)
	os.Exit(0)
}

func volume() (int, error) {
	cmd := exec.Command("amixer", "sget", "Master")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return -1, err
	}

	vol := 0

	scanner := bufio.NewScanner(bytes.NewBuffer(out))
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "Front Left: Playback") {
			matches := volRegex.FindStringSubmatch(text)
			if len(matches) != 2 {
				return -1, fmt.Errorf("expected 2 matches got %d", len(matches))
			}

			vol, err = strconv.Atoi(matches[1])
			if err != nil {
				return -1, err
			}
		}
	}

	return vol, nil
}
