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
		- TODO Membership Changes
		- TODO Client
		  todo:: 1622713555915
		  id:: 60b8a680-048d-4de9-8cf0-24ff3625405a
		- TODO ((60b997b2-776b-4287-8450-c1aad5a25d0f)) Building a Server
		  doing:: 1622796287168
		  todo:: 1622796288089
		  collapsed:: true
			- DONE Broker
			- TODO Server need handle RPC requests
			- TODO  Server need send heartbeat
			  todo:: 1622773594108
	- Server-Questions
	  id:: 60b997b2-776b-4287-8450-c1aad5a25d0f
		- Using Broker Server as the register center and registry center
	- Step 1
		- Implement a Server , and how Client find the Server
			- Using Configuration on Client
	- Final
		- https://github.com/holicc/go-raft
-
# Material
	- [Github](https://raft.github.io/)
	- [raft.pdf](../assets/raft_1622712969638_0.pdf)