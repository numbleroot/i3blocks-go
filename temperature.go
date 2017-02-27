package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"io/ioutil"
)

func main() {

	// Allow to specify high and critical thresholds.
	highFlag := flag.Int("highTemp", 72, "Specify which temperature threshold in Celsius is considered high.")
	criticalFlag := flag.Int("criticalTemp", 80, "Specify which temperature threshold in Celsius is considered critical.")
	flag.Parse()

	// Gather temperature thresholds.
	criticalTemp := *criticalFlag
	highTemp := *highFlag
	diffTemp := criticalTemp - highTemp

	// Set display texts to defaults.
	var icon string
	var output string
	var fullText string = "error"
	var shortText string = "error"

	// Read CPU temperature information from kernel
	// pseudo-file-system mounted at /sys.
	tempRaw, err := ioutil.ReadFile("/sys/class/hwmon/hwmon0/temp1_input")
	if err != nil {

		// Write fallback string to STDOUT and fail.
		fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
		os.Exit(1)
	}

	// Trim spaces.
	tempString := strings.TrimSpace(string(tempRaw))

	// Convert temperature string to integer.
	temp, err := strconv.Atoi(tempString)
	if err != nil {

		// Write fallback string to STDOUT and fail.
		fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
		os.Exit(1)
	}

	// Normalize temperature value.
	temp = temp / 1000

	// Define temperature values in dependence on
	// specified high and critical values.
	mediumTemp := highTemp - diffTemp
	lowTemp := mediumTemp - diffTemp

	// Depending on current temperature value,
	// set appropriate thermometer icon.
	if temp <= lowTemp {
		icon = ""
	} else if (temp > lowTemp) && (temp <= mediumTemp) {
		icon = ""
	} else if (temp > mediumTemp) && (temp <= highTemp) {
		icon = ""
	} else if (temp > highTemp) && (temp <= criticalTemp) {
		icon = "<span foreground=\"#ffae00\"></span>"
	} else {
		icon = "<span foreground=\"#ff0000\"></span>"
	}

	// Build final output string.
	output = fmt.Sprintf("%s%3d°C", icon, temp)

	fullText = output
	shortText = output

	// Write out gathered information to STDOUT.
	fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
	os.Exit(0)
}
