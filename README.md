hand-zmq
========

#### Chaac's Hand to handle Axe with ZMQ edge

This is the repo for Server acting as Subscriber for ZeroMQ aware Publisher clients.

This will be following types of Server:
* hand-zmq : There will be a ZeroMQ aware Server. (intended language: golang)

Then there will also be following types of servers:
* hand-socket : POSIX Socket-based Server accepting "axe-n0de-socket" data and feeding it to "hand-zmq". (intended language: golang)
* hand-http : HTTP HEAD aware server (custom) handling "axe-n0de-http" data and feeding it to "hand-zmq". (intended language: golang)


