package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	ipHosts := flag.String("h", "", "print all IPs of a host")
	ipFlag := flag.String("ip", "", "print all hosts of an IP")
	interfaceFlag := flag.Bool("i", false, "print all interfaces and IPs")
	flag.Parse()

	if *interfaceFlag {
		printInterfaces()
	}

	if *ipFlag != "" {
		printCIDR(ipFlag)
		printHosts(ipFlag)
	}

	if *ipHosts != "" {
		printIps(ipHosts)
	}
}

// IP addresses can be parsed for errors
// Besides net package functions, there is
// an IP struct that has several more
// helper methods.
func printCIDR(ipFlag *string) {
	address := net.ParseIP(*ipFlag)
	fmt.Println(address)
	mask := net.CIDRMask(24, 32)
	network := address.Mask(mask)
	fmt.Println(mask, network)
}

// net package has several Lookup functions
// that can convert between domain names
// and IP address formats
func printIps(ipHosts *string) {
	ips, _ := net.LookupHost(*ipHosts)
	fmt.Println(ips)
}

func printHosts(ipFlag *string) {
	hosts, _ := net.LookupAddr(*ipFlag)
	fmt.Println(hosts)
}

// net.Interfaces() returns a slice of all
// available interfaces on a system, each
// with a list of IP addresses as well as
// boolean flags for various interface settings
func printInterfaces() {
	interfaces, _ := net.Interfaces()
	//fmt.Println(interfaces)

	for _, i := range interfaces {
		iName, _ := net.InterfaceByName(i.Name)
		addresses, _ := iName.Addrs()
		fmt.Println(iName.Name, addresses)
	}
}
