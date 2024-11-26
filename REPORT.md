# Gobyd - A simple bidding REST API

I didn't acctually finish this program but this report is written as how i wanted it to end up.

This is a simple bidding REST API that allows users to create auctions and place bids on them.
The API is build using mutual exclusion to ensure that only one bid is accepted at a time.
We also use replication to ensure that the data is not lost in case of failure on a node.
Nodes addresses are hardcoded in the code, to make service discovery easier.

Nodes all have a node ID. This is a UUID that is generated when the node starts. This ID is used to identify the node in the system.
The reason we are not using the process ID og the running process is because we want to be able to use the code in a containerized environment.
Here the process ID is not necessarily unique across nodes, since each container is an isolated environment. In some very unlikly event,
the process ID of 2 nodes might be the same in different containers, and therefore we cannot use it to identity the nodes uniquely.

The system consists of 2 REST APIs. The first REST API is used by users for bidding on auctions.
The other REST API is used for internal communication in the system. The reason they are seperated is because it will make it much easier to limit
access since we don't want normal users to be able to access the internal communication. This means that the system exposes 2 ports, where only 1
of the ports is open to every on the internet. The other port is only open to other nodes on the same network.

## Vector Clocks

This implementation uses a vector clock to keep track of the logical time of events in the system.
A vector clock is better than a lamport clock because it can be used to detect causality between events.

A vector clock works be keeping a vector of logical clocks, with an entry for one for each node in the system.
When a node sends a message, it increments its own logical clock and sends the message with the updated vector clock.
When a node receives a message, it also increments its own logical clock. It then merges the two clocks by making sure
each entry in the clock is at least as high as the received counterpart.

In this implementation, the vector clock is sent through HTTP headers. It uses a custom header called `V-Clock` to set the vector clock in the request.
The value of the header is formatted as `NodeID=Timestamp` (E.g. `0423c67a-b7c9-4086-a485-ec76832af11f=5`). The header is set once for each node that
the sender knowns. Each node also sets a `Requester-ID` header in the request to let the receiver know who sent the request.
If a node is not known by the local clock, it assumes that it is just starting and sets the clock to 0 (or whatever was received in the headers).

Here is an example of an request to get access to the critical section containing the vector clock and requester ID headers:

```http
GET /mutex/request HTTP/1.1
User-Agent: GoByd Mutex Client
Accept: application/json
Requester-ID: 333026fd-0d73-442f-b4a6-0b8016535732
V-Clock: 333026fd-0d73-442f-b4a6-0b8016535732=2
V-Clock: bb76d29a-f275-42c8-b3bf-10b61b187a84=4
V-Clock: ee7b905c-1dc3-44f8-9e14-59888ab61760=1
```

## Mutual exclusion

This system uses the Ricart & Agrawala algorithm to obtain mutual exclusion. This means that whenever a node want access to the critical zone,
it sends a message to all other nodes in the system. Each node then reply if it can have access to the critical zone. When a node have obtained
access from all other nodes, it can then proceed to access the critical zone. If multiple nodes request access at the same time, then the node with
the earliest timestamp receives access first. Whenever the nodes is done using the critical section, a message is sent to all other nodes that they
can now request access to the critical section again.

How i implemented mutual exclusion is by each node having thier own REST API seperated from the main bidding REST API, used to communicate between
eachother. The nodes uses 3 endpoints to handle the different events needed to obtain mutual exclusion:

- `mutex/request`: To request access to the critical zone.
- `mutex/release`: To release access to the critical zone.
- `mutex/notify`: To tell others that they can request access again.

When all these endpoints are defined, mutual exclusion can be obtained for the system.

## Replication

This system uses Mast-Master replication. This means that each node is its own master and can perform both read and write operations.
Whenever this means that if we don't handle the reading and writing properly, we might run in to race conditions over the network.
We fix these issues by using [mutual exlusion](#mutual-exclusion). This means that when a node want to read or write to the critical section,
it obtains a network lock that stops other nodes from reading and writing to the system at the same time. It then first releases this lock when
it have replicated its data to all other nodes, ensuring that no race conditions can happend.

## Correctness 1

The implementation that this system uses is more aligned with ensuring sequential consistency, as the use of vector clocks can track causality and
maintain consistent sequencing of events. I don't think lineariability might be fully achived since the systems main focus is mutual exclusion.

## Correctness 2

In the absence of failures, the implementation would ensure correctness through the use of mutual exclusion via the Ricart & Agrawala algorithm. Each node must obtain permission from all other nodes before accessing a critical section, ensuring no race conditions, which means maintaining the system's integrity.

## Others

[GitHub repo: https://github.com/Kanerix/gobyd](https://github.com/Kanerix/gobyd)

[Logs: https://github.com/Kanerix/gobyd/logs/bidding.log](https://github.com/Kanerix/gobyd/logs/bidding.log)
