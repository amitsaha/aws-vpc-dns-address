package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

// Get the Amazon DNS IPv4 address for a VPC
// https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_DHCP_Options.html#AmazonDNS
// Basically a golang version of the comment by Dusan Bajic in
// https://stackoverflow.com/questions/39100395/getting-the-dns-ip-used-within-an-aws-vpc

// Running the program will print the DNS server, which you can use for example
// to set the DNS server in docker to be able to resolve private DNS names.

func main() {
	res, err := http.Get("http://169.254.169.254/latest/meta-data/mac")
	if err != nil {
		log.Fatal(err)
	}
	macAddress, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("http://169.254.169.254/latest/meta-data/network/interfaces/macs/%s/vpc-ipv4-cidr-block", macAddress)
	res, err = http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	primaryVpcCIDR, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	ipv4Addr, _, err := net.ParseCIDR(string(primaryVpcCIDR))
	if err != nil {
		log.Fatal(err)
	}

	// Add 2 to the base ipv4Addr to get the Amazon DNS server
	// Convert IP address to a byte array, then convert it into a unit32, add and then do the reverse conversion
	ipv4Uint := binary.BigEndian.Uint32([]byte(ipv4Addr.To4()))
	dnsUint := ipv4Uint + 2
	buff := make(net.IP, 4)
	binary.BigEndian.PutUint32(buff, dnsUint)
	fmt.Printf("%v\n", buff)
}
