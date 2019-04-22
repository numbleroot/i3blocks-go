package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"io/ioutil"
	"net/http"
)

func main() {
	// Url to get IP info
	var urlForIp string = "https://ifconfig.co/ip"
	var urlForCity string = "https://ifconfig.co/city"

	// Set display texts to defaults.
	var fullText string = "unknown"
	var shortText string = "unknown"
	var whatIsMyIpUrl string

	var cityFlag = flag.Bool("city", false, "pass -city to report City instead of IP")
	flag.Parse()

	if cityFlag != nil && *cityFlag {
		whatIsMyIpUrl = urlForCity
	} else {
		whatIsMyIpUrl = urlForIp
	}

	// Request whats-my-ip service to return this
	// machine's public IP address via TLS.
	resp, err := http.Get(whatIsMyIpUrl)
	if err != nil {

		// Write an error to STDERR, fallback display values
		// to STDOUT and exit with failure code.
		fmt.Fprintf(os.Stderr, "[i3blocks-go] Failed to get response from public IP service: %s", err.Error())
		fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
		os.Exit(0)
	}

	// Read-in body part of response containing
	// the raw IP address.
	ipRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[i3blocks-go] Could not read body part of IP service response: %s", err.Error())
		fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
		os.Exit(0)
	}
	resp.Body.Close()

	// Remove surrounding space.
	ip := strings.TrimSpace(string(ipRaw))

	fullText = ip
	shortText = ip

	// Write out gathered information to STDOUT.
	fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
	os.Exit(0)
}
