package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

// Set display texts to defaults.
var fullText = "unknown"
var shortText = "unknown"

func main() {
	var iface = &net.Interface{}

	// Check if interface parameter was passed in
	var ifaceNameFlag = flag.String("ifaceName", "", "-ifaceName wlp2s0")
	flag.Parse()

	if ifaceNameFlag == nil || *ifaceNameFlag == "" {
		// ifaceName parameter was not passed, use first interface that is up
		desiredIface, err := firstNonLoopbackInterfaceThatIsUp()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			writeAndExit()
		}
		iface = desiredIface
	} else {
		// ifaceName was passed, use the name to get interface
		desiredIface, err := net.InterfaceByName(*ifaceNameFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[i3blocks-go] Failed to find interface with name %s. %s", *ifaceNameFlag, err.Error())
			writeAndExit()
		}
		iface = desiredIface
	}

	// Retrieve all addresses belonging to
	// found active interface.
	addrs, err := iface.Addrs()
	if err != nil || len(addrs) == 0 {
		fmt.Fprintf(os.Stderr, "[i3blocks-go] Failed to retrieve IP addresses associated with %s: %v", iface.Name, err)
		writeAndExit()
	}

	// Extract the actual IP from address string
	// and remove subnet information.
	ipRaw := strings.Split(addrs[0].String(), "/")

	fullText = ipRaw[0]
	shortText = ipRaw[0]
	writeAndExit()
}

func firstNonLoopbackInterfaceThatIsUp() (*net.Interface, error) {
	// Get all interfaces associated with this machine.
	ifaces, err := net.Interfaces()
	if err != nil {
		// Write an error to STDERR, fallback display values
		// to STDOUT and writeAndExit with failure code.
		return &net.Interface{}, fmt.Errorf("[i3blocks-go] Failed to retrieve local interfaces: %s", err.Error())
	}

	for _, iface := range ifaces {
		// If current interface is not marked as being
		// active, continue to next for loop iteration.
		if (iface.Flags & net.FlagUp) != 1 {
			continue
		}
		// If this interface is the machine's loopback
		// interface, continue to next for loop iteration.
		if (iface.Flags & net.FlagLoopback) == 4 {
			continue
		}
		return &iface, nil
	}
	return &net.Interface{}, errors.New("non loopback interface that is up not found")
}

// Write out gathered information to STDOUT.
func writeAndExit() {
	fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
	os.Exit(0)
}
