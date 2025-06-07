# Raft consensus algorithm implementation in Golang

## Project Overview
This project is an implementation of the Raft consensus algorithm in Go which use the standard library as much as possible. The core Raft logic is located in the `raft/` directory, and the main executable is built from `cmd/main.go`.

> Not fully implemented yet.

## Getting Started
To get started with this Raft implementation, follow these instructions:

### Prerequisites
- Go (version 1.18 or later)

### Building the project
1. Clone the repository:
   ```bash
   git clone https://github.com/PhrolovaC/go-raft.git
   cd go-raft
   ```
2. Build the executable:
   ```bash
   go build ./cmd/main.go
   ```
   This will create an executable file named `main` (or `main.exe` on Windows) in the current directory.

### Running the project
To run a Raft node, you need to provide a unique ID for the node and a list of peer addresses.

Execute the binary with the required flags:
```bash
./main -id <node-id> -peers <peer1-addr,peer2-addr,...>
```
For example:
```bash
./main -id node1 -peers localhost:8001,localhost:8002,localhost:8003
```
You will need to run multiple instances of the application on different terminals (or machines) to form a cluster, each with a unique `-id` and the same list of `-peers`.

## Usage
The main executable `cmd/main.go` accepts the following command-line arguments:

- `-id <string>`: **(Required)** The unique identifier for the Raft node. This can be any string that uniquely identifies the node within the cluster (e.g., "node1", "serverA").
- `-peers <string>`: **(Required)** A comma-separated list of peer addresses for all nodes in the Raft cluster. Each address should be in the format `host:port` (e.g., "localhost:8001,localhost:8002,localhost:8003"). All nodes in the cluster should be started with the same list of peers.

## References
- https://raft.github.io/raft.pdf