package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type IpInfo struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Region   string `json:"region"`
}

func main() {
	// Request whats-my-ip service to return this
	// machine's public IP address via TLS.
	resp, err := http.Get("http://ipinfo.io/json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "server returned %s not 200 OK", resp.Status)
		os.Exit(1)
	}

	var info IpInfo
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&info); err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode json: %v", err)
		os.Exit(2)
	}

	// Set display texts to defaults.
	fullText := "unknown"
	shortText := "unknown"

	// If response indeed contained an IP address,
	// set i3bar protocols fields accordingly.
	if info.IP != "" {
		fullText = info.IP
		shortText = info.IP
	}

	// Write out gathered information to STDOUT.
	fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
	os.Exit(0)
}
