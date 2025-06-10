package raft

// ClientEnd represents a RPC endpoint.
type ClientEnd interface {
	Call(method string, args any, reply any) bool
}

func Call(method string, args any, reply any) bool {

	return true
}
