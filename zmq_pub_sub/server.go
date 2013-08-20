package zmq_pub_sub

import (
  // go packages
  "time"
  // 3party packages
  zmq "github.com/alecthomas/gozmq"
)

func Server() {
  println("Starting Chaac's Hand. Giving from 8080 and Receiving from 8081 ;)")

  context, _ := zmq.NewContext()
  receiver, _ := context.NewSocket(zmq.PULL)
  receiver.Bind("tcp://*:8080")
  sender, _ := context.NewSocket(zmq.PUB)
  sender.Bind("tcp://*:8081")

  last := time.Now()
  messages := 0
  for {
    message, _ := receiver.Recv(0)
    sender.Send(message, 0)
    messages += 1
    now := time.Now()
    if now.Sub(last).Seconds() > 10 {
      println(messages, "msg/sec")
      last = now
      messages = 0
    }
  }

}
