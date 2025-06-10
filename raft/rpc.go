package raft

// This file contains the definitions of Args and Reply for RPC between Raft nodes,
// as well as messages passed to the application layer.

// AppendEntriesArgs struct is used by leaders to replicate log entries to followers
// and also as a heartbeat mechanism.
type AppendEntriesArgs struct {
	Term         int
	LeaderId     int
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []*LogEntry
	LeaderCommit int
}

// AppendEntriesReply struct is sent by a follower in response to an AppendEntries RPC.
type AppendEntriesReply struct {
	Term    int
	Success bool
}

// RequestVoteArgs struct is used by candidates to request votes from other peers.
type RequestVoteArgs struct {
	Term         int
	CandidateId  int
	LastLogIndex int
	LastLogTerm  int
}

// RequestVoteReply structure is sent by a peer in response to a RequestVote RPC.
type RequestVoteReply struct {
	Term        int
	VoteGranted bool
}

// ApplyMsg structure is used to send committed log entries (commands) or snapshots
// from the Raft library to the service/application using Raft.
type ApplyMsg struct {
	CommandVaild bool // true if this message delivers a new command to the application
	Command      any  // the command that was agreed upon by Raft
	CommandIndex int  // the log index of the command

	SnapshotValid bool   // true if this message delivers a snapshot to the application
	Snapshot      []byte // raw bytes of the snapshot (application is responsible for interpreting)
	SnapshotTerm  int    // the term of the last log entry included in the snapshot
	SnapshotIndex int    // the log index of the last log entry included in the snapshot
}
