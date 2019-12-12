package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"sort"
	"time"
)

func dnsQuery(host string) ([]string, error) {
	// Sets up the regex needed to validate an IPv4 IP address.
	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regEx := regexp.MustCompile(numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock)

	// Sets an empty slice (array) to hold all ips we find.
	var ipList []string

	// Executes the DNS lookup, returns array of resolved IPs from hostname.
	ips, err := net.LookupIP(host)

	// Returns an error if unable to resolve DNS.
	if err != nil {
		fmt.Println("[-] DNS Resolution Error", err)
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

func generateOutput(now time.Time, outFormat string, hostNameArg string, dnsQueryResult []string) string {
	// Create an empty string, we can concatenate into, later.
	outString := ""

	// Check to see if the desired output format is json
	if outFormat == "json" {

		// Create an empty map, ready to add data
		jsonData := make(map[string]interface{})

		// Add data to json map.
		jsonData["scriptExecution"] = now.Format(time.RFC3339)
		jsonData["resolvedIps"] = dnsQueryResult
		jsonData["queriedHostname"] = hostNameArg

		// Marshall the JSON data so we can convert it to string
		outData, err := json.Marshal(jsonData)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}

		// Convert the marshalled json data to string, and return it
		outString += string(outData)
		return outString
	}

	// Check to see if the desired output format is stdout
	if outFormat == "stdout" {
		outString += "[+] Domain: " + hostNameArg + "\n"
		outString += "[+] Run Time: " + now.Format(time.RFC3339) + "\n"

		ctr := 0
		for _, ip := range dnsQueryResult {
			ctr ++
			outString += fmt.Sprintf("%v", ctr)+". "+ip+"\n"
		}
		return outString
	}

	// Didnt receive a valid outFormat, so return this string.
	return "No Output Generated"
}

func main() {
	// Gets the current script execution time.
	now := time.Now()

	// Parses input args from command line arguments, ignores program.
	cliArgs := os.Args[1:]

	// Makes sure we get at least 2 items, and only 2 items, as an argument.
	if len(cliArgs) != 2 {
		fmt.Println("[-] Invalid Command Line Arguments. Usage: ./go-dns-client google.com json")
		os.Exit(1)
	}

	// Parses the hostname as 0th element from cliArgs
	hostNameArg := cliArgs[0]
	// sets the desired output format. Valid formats: stdout or json
	outputArg := cliArgs[1]

	// Queries the hostname from CLI args from DNS. Exit if returns error.
	dnsQueryResult, err := dnsQuery(hostNameArg)
	if err != nil {
		os.Exit(1)
	}

	// verify input for output format is a valid type
	output := generateOutput(now, outputArg, hostNameArg, dnsQueryResult)

	// prints out the return value from generateOutput()
	fmt.Println(output)
}