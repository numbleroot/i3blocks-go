package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	// Set display texts to defaults.
	var fullText string = "unknown"
	var shortText string = "unknown"

	// Request whats-my-ip service to return this
	// machine's public IP address via TLS.
	resp, _ := http.Get("https://ip.wirelab.org/")

	// Read-in body part of response containing
	// the raw IP address.
	ipRaw, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// Remove surrounding space.
	ip := strings.TrimSpace(string(ipRaw))

	// If response indeed contained an IP address,
	// set i3bar protocols fields accordingly.
	if ip != "" {
		fullText = ip
		shortText = ip
	}

	// Write out gathered information to STDOUT.
	fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
	os.Exit(0)
}
