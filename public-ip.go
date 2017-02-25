package main

import (
	"fmt"
	"os"
	"strings"

	"io/ioutil"
	"net/http"
)

func main() {

	// Per default, exit with a success code.
	var retValue int = 0

	// Set both display texts to 'unknown' by default.
	var fullText string = "unknown"
	var shortText string = "unknown"

	// Request whats-my-ip service to return this
	// machine's public IP adress via TLS.
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

	os.Exit(retValue)
}
