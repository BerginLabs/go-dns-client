package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"sort"
)

func dnsLookup(host string) ([]string, error) {
	// Sets up the regex needed to validate an IPv4 IP address.
	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regEx := regexp.MustCompile(numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock)

	// Sets an empty slice (array) to hold all ips we find.
	var ipList []string

	// Executes the DNS lookup, returns array of resolved IPs from hostname.
	ips, err := net.LookupIP(host)

	// Returns an error if unable to resolve DNS.
	if err != nil {
		fmt.Println("[+] DNS Resolution Error", err)
		return nil, errors.New("unable to resolve DNS query")
	}

	// Loops over each result from our DNS query, regexes for valid IPv4 IP, and appends result to ipList.
	for _, ip := range ips {
		ipRegex := regEx.FindString(ip.String())
		if ipRegex != "" {
			ipList = append(ipList, ipRegex)
		}
	}

	// Sorts the list of IP strings smallest to largest in our list.
	sort.Strings(ipList)
	return ipList, nil
}

func main() {
	// Parses input args from command line arguments, ignores program.
	cliArgs := os.Args[1:]

	// Makes sure we get at least 1 item, and only 1 item, as an argument.
	if len(cliArgs) != 1 {
		fmt.Println("[-] Invalid Command Line Arguments. Usage: ./go-dns-client google.com")
		os.Exit(1)
	}

	// Parses the hostname as 0th element from cliArgs
	hostNameArg := cliArgs[0]
	fmt.Println("[+] Starting DNS Lookup for domain:", hostNameArg)

	// Queries the hostname from CLI args from DNS. Exit if returns error.
	dnsQueryResult, err := dnsLookup(hostNameArg)

	// Exits the program with non-zero-status exit code if we cant resolve the hostname via DNS.
	if err != nil {
		os.Exit(1)
	}

	// Loop through query results, and parse and print valid ipv4 addresses.
	for _, ip := range dnsQueryResult {
		fmt.Println("[+] Query result:", ip)
	}

	fmt.Println("[+] DNS Resolution Complete.")
}