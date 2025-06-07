package raft

import (
	"sync"
	"time"
)

// Raft is the main struct representing a Raft node.
type Raft struct {
	mu        sync.Mutex   // Lock to protect shared access to this peer's state
	peers     []*ClientEnd // RPC end points of all peers
	persister *Persister   // Object to hold this peer's persisted state
	me        int          // this peer's index in the peers[]

	// Persistent state on all servers
	currenTerm int         // lastest term server has seen
	voteFor    int         // candidateId that received vote in current term
	log        []*LogEntry // log entries

	// Volatile state on all servers
	commitIndex int // index of highest log entry known to be committed
	lastApplied int // index of highest log entry applied to state machine

	// Volatile state on leaders
	nextIndex  []int // for each server, index of the next log entry to send to that server
	matchIndex []int // for each server, index of highest log entry known to be replicated on server

	role Role // current role of this peer in the cluster

	// timers
	electionTimeout   time.Duration // timeout for starting a election
	lastHeartbeatTime time.Time     // time of last heartbeat

	applyCh chan ApplyMsg // channel for sending committed log entries to the state machine
}

// Role defines the possible role of a Raft server.
type Role int

const (
	Leader Role = iota
	Follower
	Candidate
)

// Make creates a new Raft server and initializes its state.
func Make(peers []*ClientEnd, me int, persister *Persister, applyCh chan ApplyMsg) *Raft {
	rf := &Raft{}
	rf.peers = peers
	rf.me = me
	rf.persister = persister
	rf.applyCh = applyCh

	rf.currenTerm = 0
	rf.voteFor = -1
	rf.log = make([]*LogEntry, 0)

	rf.commitIndex = 0
	rf.lastApplied = 0

	rf.nextIndex = make([]int, len(peers))
	rf.matchIndex = make([]int, len(peers))

	rf.role = Follower

	rf.electionTimeout = randomizedElectionTimeout()
	rf.lastHeartbeatTime = time.Now()

	return rf
}

// RequestVote is a RPC handler which candidate calls to request a vote.
func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {

}

// AppendEntries is a RPC handler which leader calls to replicate log entries or send heartbeats.
func (rf *Raft) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) {

}
