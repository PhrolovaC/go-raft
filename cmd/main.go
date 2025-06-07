package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	id := flag.String("id", "", "ID of the Raft node (required)")
	peersList := flag.String("peers", "", "Comma-separated list of peer addresses (required)")

	flag.Parse()

	if *id == "" {
		log.Fatal("Error: -id flag is required")
	}

	if *peersList == "" {
		log.Fatal("Error: -peers flag is required")
	}

	peers := strings.Split(*peersList, ",")
	if len(peers) == 0 {
		log.Fatal("Error: -peers list cannot be empty")
	}

	fmt.Printf("Node ID: %s\n", *id)
	fmt.Printf("Peer Addresses: %v\n", peers)

	// Placeholder for Raft node initialization and starting
	// raftNode := raft.NewRaftNode(*id, peers, ...) // Adjust parameters as needed
	// raftNode.Start()

	// Keep the main function running (e.g., by listening on a channel or a signal)
	// select {}
}
