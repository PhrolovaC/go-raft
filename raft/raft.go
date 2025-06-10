package raft

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Raft is the main struct representing a Raft node.
type Raft struct {
	mu        sync.Mutex  // Lock to protect shared access to this peer's state
	peers     []ClientEnd // RPC end points of all peers
	persister *Persister  // Object to hold this peer's persisted state
	me        int         // this peer's index in the peers[]
	dead      int32       // flag to indicate if the node is running

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
func Make(peers []ClientEnd, me int, persister *Persister, applyCh chan ApplyMsg) *Raft {
	// Create and initialize a raft instance.
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

	// Restore persisted state if available.
	// rf.readPersist()

	// Start background goroutine.
	go rf.ticker()

	return rf
}

// RequestVote is a RPC handler which candidate calls to request a vote.
func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	reply.VoteGranted = false
	reply.Term = rf.currenTerm
	// Receiver implementation 1
	if args.Term < rf.currenTerm {
		return
	}

	if args.Term > rf.currenTerm {
		rf.becomeFollower(args.Term)
	}

	// Receiver implementation 2
	upToDate := func() bool {
		lastLogIndex, lastLogTerm := rf.getLastIndexAndTerm()
		if args.LastLogTerm != lastLogTerm {
			return args.LastLogTerm > lastLogTerm
		}
		return args.LastLogIndex >= lastLogIndex
	}()
	if (rf.voteFor != -1 || rf.voteFor == args.CandidateId) && upToDate {
		rf.voteFor = args.CandidateId
		reply.VoteGranted = true
	}

	rf.electionTimeout = randomizedElectionTimeout()
}

// AppendEntries is a RPC handler which leader calls to replicate log entries or send heartbeats.
func (rf *Raft) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) {

}

// send a RequestVote RPC to a server.
// server is the index of the target server in rf.peers[].
func (rf *Raft) sendRequestVote(server int, args *RequestVoteArgs, reply *RequestVoteReply) bool {
	ok := rf.peers[server].Call("Raft.RequestVote", args, reply)
	return ok
}

// send a AppendEntries RPC to a server.
func (rf *Raft) sendAppendEntries(server int, args *AppendEntriesArgs, reply *AppendEntriesReply) bool {
	ok := rf.peers[server].Call("Raft.AppendEntries", args, reply)
	return ok
}

// startElection is the method called by a follower if doesn't receive heartbeat within electionTimeout.
func (rf *Raft) startElection() {
	rf.mu.Lock()
	rf.becomeCandidate()
	rf.currenTerm++
	rf.voteFor = rf.me

	numPeers := len(rf.peers)
	me := rf.me
	currentTerm := rf.currenTerm
	lastLogIndex, lastLogTerm := rf.getLastIndexAndTerm()
	rf.mu.Unlock()

	var cnt int32 = 1
	for i := range numPeers {
		if i == me {
			continue
		}

		go func(server int) {
			args := &RequestVoteArgs{
				Term:         currentTerm,
				CandidateId:  me,
				LastLogIndex: lastLogIndex,
				LastLogTerm:  lastLogTerm,
			}
			reply := &RequestVoteReply{}

			ok := rf.sendRequestVote(server, args, reply)
			if ok {
				rf.mu.Lock()
				if reply.Term > rf.currenTerm {
					rf.becomeFollower(reply.Term)
					rf.mu.Unlock()
					return
				}
				rf.mu.Unlock()

				if reply.VoteGranted {
					if atomic.AddInt32(&cnt, 1) > int32(numPeers)/2 {
						rf.mu.Lock()
						if rf.role == Candidate && rf.currenTerm == currentTerm {
							rf.becomeLeader()
						}
						rf.mu.Unlock()
					}
				}
			}
		}(i)
	}
}

// Kill is called to gracefully shut down the Raft node.
func (rf *Raft) Kill() {
	atomic.StoreInt32(&rf.dead, 1)
}

// ticker is a goroutine that manages election and heatbeat.
func (rf *Raft) ticker() {
	for atomic.LoadInt32(&rf.dead) != 1 {
		rf.mu.Lock()
		timeout := rf.electionTimeout
		last := rf.lastHeartbeatTime
		role := rf.role
		rf.mu.Unlock()

		if role != Leader && time.Since(last) > timeout {
			rf.startElection()
		}

		// pause for a random amount of time.
		ms := 10 + (rand.Int63() % 10)
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}
}
