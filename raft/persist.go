package raft

import "sync"

// Persister is a struct for saving and loading persistent state.
type Persister struct {
	mu        sync.Mutex
	raftstate []byte
	snapshot  []byte
}

// persist saves the Raft node's persistent state to stable storage.
// This function should be called whenever the persistent state changes.
func (rf *Raft) persist() {

}

// readPersist restores the previously persisted state into memory.
// This function should be called when a Raft node starts up or recovers from a crash.
func (rf *Raft) readPersist(data []byte) {

}
