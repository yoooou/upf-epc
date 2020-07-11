// SPDX-License-Identifier: Apache-2.0
// Copyright(c) 2020 Intel Corporation

package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

// Increment increments the given net.IP by one bit. Incrementing the last IP in an IP space (IPv4, IPV6) is undefined.
func increment(ip net.IP) net.IP {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		//only add to the next byte if we overflowed
		if ip[i] != 0 {
			break
		}
	}
	return ip
}

/*
func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
*/

func hex2int(hexStr string) uint32 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 32)
	return uint32(result)
}

func getOutboundIP(dstIP string) net.IP {
	conn, err := net.Dial("udp", dstIP+":"+PFCPPort)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
