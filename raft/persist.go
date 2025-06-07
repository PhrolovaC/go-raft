package raft

import "sync"

// Persister is a struct for saving and loading persistent state.
type Persister struct {
	mu        sync.Mutex
	raftstate []byte
	snapshot  []byte
}
