package main

import (
	"fmt"
)

type topic map[string][]chan interface{}

// PubSub acts as a publish/subscribe server
type PubSub struct {
	topics topic
}

func NewPubSub() *PubSub {
	return &PubSub{
		topics: make(topic),
	}
}

func (ps *PubSub) Subscribe(topic string) chan interface{} {
	ch := make(chan interface{}, 1)
	var subchs []chan interface{}
	subchs, _ = ps.topics[topic]
	subchs = append(subchs, ch)
	ps.topics[topic] = subchs
	return ch
}

func (ps *PubSub) Publish(payload interface{}, topic string) {
	if chs, ok := ps.topics[topic]; ok {
		for _, ch := range chs {
			ch <- payload
		}
	}
}

func main() {
	fmt.Println("Starting PUBSUB server...")	
	server := NewPubSub()
	sub := server.Subscribe("news")
	server.Publish("what is up?", "news")
	fmt.Println(<-sub)
}