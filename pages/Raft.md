alias:: Raft , Algorithm

# What's is Raft ?

 **For Short**: Raft is consensus algorithms just like *Paxos* , but more simple and easier to understand.
# How to implement a Raft ?
	-
	  #+BEGIN_IMPORTANT
	  **The implement steps are follow by question**
	  #+END_IMPORTANT
	- Cores
		- DOING Leader election
		  doing:: 1622714063936
			-
			  #+BEGIN_QUOTE
			  Leaders send periodic
			  heartbeats (AppendEntries RPCs that carry no log entries)
			  to all followers in order to maintain their authority.
			  #+END_QUOTE
		- TODO VoteRequest
			-
		- TODO ((60bf2316-d07a-4e6a-b196-f9258090739f)) Building a Client
		  id:: 60b8a680-048d-4de9-8cf0-24ff3625405a
		  todo:: 1622713555915
		- TODO ((60b997b2-776b-4287-8450-c1aad5a25d0f)) Building a Server
		  doing:: 1622796287168
		  todo:: 1622796288089
		  collapsed:: true
			- DONE Broker
			- TODO Server need handle RPC requests
			- TODO  Server need send heartbeat
			  todo:: 1622773594108
		- TODO  *election timeout*
			-
			  #+BEGIN_QUOTE
			  If a follower receives no communication over a period of time called the election timeout, then it assumes there is no viable leader and begins an election to choose a new leader
			  #+END_QUOTE
		- TODO Split Votes
			-
			  #+BEGIN_QUOTE
			  The third possible outcome is that a candidate neither
			  wins nor loses the election: if many followers become
			  candidates at the same time, votes could be split so that
			  no candidate obtains a majority. When this happens, each
			  candidate will time out and start a new election by incrementing its term and initiating another round of RequestVote RPCs. However, without extra measures split votes
			  could repeat indefinitely.
			  
			  #+END_QUOTE
	- Client-Questions
	  id:: 60bf2316-d07a-4e6a-b196-f9258090739f
		- What's happen if request a server , not on leader state ?
			-
			  #+BEGIN_QUOTE
			  Clients of Raft send all of their requests to the leader.
			  When a client first starts up, it connects to a randomlychosen server. If the client’s first choice is not the leader,
			  that server will reject the client’s request and supply information about the most recent leader it has heard from
			  (AppendEntries requests include the network address of
			  the leader). If the leader crashes, client requests will time
			  out; clients then try again with randomly-chosen servers.
			  #+END_QUOTE
	- Server-Questions
	  id:: 60b997b2-776b-4287-8450-c1aad5a25d0f
		- Using Broker Server as the register center and registry center
	- Building Server
		- Step One:
			-
			  #+BEGIN_QUOTE
			  When servers start up, they begin as followers. A server remains in follower state as long as it receives valid RPCs from a leader or candidate.
			  #+END_QUOTE
			-
		- Implement a Server , and how Client find the Server
			- Using Configuration on Client
	- Golang-Tips
		- What's diff between *make(chan int)* vs *make(chan int , 1)*
			- If the channel is unbuffered, the sender blocks until the receiver has received the value. If the channel has a buffer, the sender blocks only until the value has been copied to the buffer; if the buffer is full, this means waiting until some receiver has retrieved a value.
	- Final
		- https://github.com/holicc/go-raft
-
# Material
	- [Github](https://raft.github.io/)
	- [raft.pdf](../assets/raft_1622712969638_0.pdf)