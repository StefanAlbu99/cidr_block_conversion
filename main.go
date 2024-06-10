package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
)

func main() {
	// Open the JSON file
	file, err := os.Open("ip_ranges_unorganized.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Decode the JSON data into a slice of strings
	var ips []string
	if err := json.NewDecoder(file).Decode(&ips); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	// Convert subnets to IPNet type
	var ipNets []*net.IPNet
	for _, s := range ips {
		_, ipNet, err := net.ParseCIDR(s)
		if err != nil {
			fmt.Printf("Error parsing subnet %s: %v\n", s, err)
			continue
		}
		ipNets = append(ipNets, ipNet)
	}

	// Sort IPNets based on IP address and mask
	sort.Slice(ipNets, func(i, j int) bool {
		if bytes.Compare(ipNets[i].IP, ipNets[j].IP) != 0 {
			return bytes.Compare(ipNets[i].IP, ipNets[j].IP) < 0
		}
		return ipNets[i].Mask.String() < ipNets[j].Mask.String()
	})

	// Aggregate contiguous IP ranges
	var aggregated []*net.IPNet
	for _, ipNet := range ipNets {
		if len(aggregated) == 0 {
			aggregated = append(aggregated, ipNet)
			continue
		}
		last := aggregated[len(aggregated)-1]
		if isContiguous(last, ipNet) {
			aggregated[len(aggregated)-1] = mergeIPNets(last, ipNet)
		} else {
			aggregated = append(aggregated, ipNet)
		}
	}

	// Print aggregated subnets
	for _, a := range aggregated {
		fmt.Println(a.String())
	}
}

// compareIP compares two IP addresses
func compareIP(a, b net.IP) int {
	return bytes.Compare(a, b)
}

// isContiguous checks if two IP ranges are contiguous
// isContiguous checks if two IP ranges are contiguous
func isContiguous(a, b *net.IPNet) bool {
	// Convert IP addresses and masks to binary representations
	aIP := a.IP.To4()
	bIP := b.IP.To4()
	aMask := a.Mask
	bMask := b.Mask

	// Ensure that the IP addresses share the same network portion except for the last byte
	for i := 0; i < len(aIP)-1; i++ {
		if aIP[i]&aMask[i] != bIP[i]&bMask[i] {
			return false
		}
	}

	// Ensure that the prefix lengths differ by exactly one
	aLen, _ := a.Mask.Size()
	bLen, _ := b.Mask.Size()
	return aLen == bLen || aLen+1 == bLen || aLen-1 == bLen
}

// mergeIPNets merges two contiguous IP ranges
func mergeIPNets(a, b *net.IPNet) *net.IPNet {
	aIP := a.IP.To16()
	_, aBits := a.Mask.Size()
	return &net.IPNet{
		IP:   aIP,
		Mask: net.CIDRMask(aBits-1, 32),
	}
}
