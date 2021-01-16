package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Heartbeat dials a (network) "tcp" (addr) "localhost:9001". If the dial returns
// a response (err == nil), return true; else return false (default).
func Heartbeat(addr string) bool {
	// TODO: Improve ...
	_, err := net.Dial("tcp", addr) //DialTimeout("tcp", addr, time.Microsecond*400)
	if err == nil {
		log.Printf("Service running on port %s ...", addr)
		return true
	}
	return false
}

// ScanLocalhost ports
func ScanLocalhost() []string {
	var active []string

	// TODO: add manifest of standard/reserved service ports to check initially
	// as a heartbeat analytics measure
	sTime := time.Now()

	// No point in checking < 22 really, but leaving
	// at 2 until absolutely necessary
	for i := 2; i < 65535; i++ {
		localAddr := fmt.Sprintf("localhost:%d", i)
		if Heartbeat(localAddr) == true {
			active = append(active, localAddr)
		}
	}

	timeTaken = time.Since(sTime)
	timeSince = time.Now()
	log.Printf("ScanLocalhost took %s ...", timeTaken)

	return active
}
