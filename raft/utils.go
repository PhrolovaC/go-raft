package raft

// This file is intended for utility functions that helpful across the Raft implementation.

func (rf *Raft) becomeLeader() {

}

func (rf *Raft) becomeFollower(term int) {
	rf.currenTerm = term
	rf.voteFor = -1

	rf.role = Follower
}

func (rf *Raft) becomeCandidate() {
	rf.role = Candidate
}

// getLastIndexAndTerm returns the lastLogIndex and lastLogTerm of a Raft node.
func (rf *Raft) getLastIndexAndTerm() (int, int) {
	lastLogIndex := len(rf.log) - 1
	lastLogTerm := 0
	if lastLogIndex >= 0 {
		lastLogTerm = rf.log[lastLogIndex].Term
	}
	return lastLogIndex, lastLogTerm
}
