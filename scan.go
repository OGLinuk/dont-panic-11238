package main

import (
	"fmt"
	"log"
	"net"
	"sort"
	"time"
)

// scan dials a (network) "tcp" (addr) "localhost:9001". If the dial returns
// a response (err == nil), return true; else return false (default).

// Notes: Changed from net.Dial to net.DialTimeout
// and added 700 microsecond timeout. Reduced scan
// tc from ~9.2 to ~7.2s. Was tested w/ 18 active ports.

// Any amount of time saved is crucial as may end up having to scan
// a lot and often. It should ideally take < 500ms. Total time
// to complete all scans will still remain to be O(n)(*0.5?).
func scan(addr string) bool {
	// TODO: Improve ...
	_, err := net.DialTimeout("tcp", addr, time.Microsecond*400)
	if err == nil {
		log.Printf("Service running on port %s ...", addr)
		return true
	}
	return false
}

// scanLocalhost ports
func scanLocalhost() []string {
	var active []string
	maxPort := 65535

	// No point in checking < 22 really, but leaving
	// at 2 until absolutely necessary
	for i := 2; i < maxPort; i++ {
		// TODO: Scan concurretly, *and consistently*.
		localAddr := fmt.Sprintf("localhost:%d", i)
		if scan(localAddr) == true {
			active = append(active, localAddr)
		}
	}

	sort.Strings(active)

	return active
}
