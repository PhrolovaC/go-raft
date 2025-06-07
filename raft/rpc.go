package raft

// This file contains the definitions of Args and Reply for RPC between Raft nodes,
// as well as messages passed to the application layer.

type AppendEntriesArgs struct {
}

type AppendEntriesReply struct {
}

type RequestVoteArgs struct {
}

type RequestVoteReply struct {
}

type ApplyMsg struct {
}
