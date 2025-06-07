package raft

// The definination of log entry in the Raft.
// The log is a sequence of commands that the Raft leader replicates to its followers.
// Once a log entry is commited, it can be applied to the state machine.

type LogEntry struct {
	Index   int
	Term    int
	Command any
}
