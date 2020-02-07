package main

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type MyMessage struct {
	Name string
}

type MyActor struct {
}

func NewMyActor() actor.Actor {
	fmt.Println("New MyActor")
	return &MyActor{}
}

func (state *MyActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case MyMessage:
		fmt.Printf("Hello %v\n", msg.Name)
		time.Sleep(1 * time.Second)
		fmt.Printf("Bye %v\n", msg.Name)
	}
}

func main() {
	props := actor.PropsFromProducer(NewMyActor)
	rootContext := actor.EmptyRootContext
	pid1 := rootContext.Spawn(props)
	pid2 := rootContext.Spawn(props)

	for true {
		pid1.Tell(MyMessage{Name: "Samuel"})
		pid2.Tell(MyMessage{Name: "Bob"})
		time.Sleep(2 * time.Second)
	}
}
