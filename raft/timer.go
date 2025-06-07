package raft

// This file outlines the timer management aspects of the Raft.

import (
	"math/rand"
	"time"
)

// Constants for timer durations.
const (
	BaseElectionTimeoutMS  = 150
	RandomElectionSpreadMS = 150
	HeartbeatIntervalMs    = 100
)

// randomizedElectionTimeout generates a random duration for the election timeout.
// The duration is typically chosen from a fixed base plus a random component.
func randomizedElectionTimeout() time.Duration {
	// Example: 150ms base + (0-150ms random) = 150-300ms.
	return time.Duration(BaseElectionTimeoutMS+(rand.Intn(RandomElectionSpreadMS))) * time.Millisecond
}

// fixedHeartbeatInterval returns the fixed interval for leader heartbeats.
// This is typically much shorter than the election timeout.
func fixedHeartbeatInterval() time.Duration {
	// This value should be less than the minimum election timeout to ensure stability.
	// Example: If election timeouts are 150-300ms, heartbeat interval could be 50ms or 100ms.
	return time.Duration(HeartbeatIntervalMs) * time.Millisecond
}
