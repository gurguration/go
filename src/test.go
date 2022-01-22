package main

import "log"

type subsciber struct {
	subids string
}

func (s subsciber) id() string {
	return s.subids
}
func (s subsciber) react(msg string) {
	log.Printf("ID %s - received: %s", s.id(), msg)
}

func (p publisher) removeSubscriber(subid string) {
	delete(p.subscribers, subid)
}
func newSubscriber(id string) subsciber {
	return subsciber{subids: id}
}

type Subscriber interface {
	id() string
	react(msg string)
}
type publisher struct {
	subscribers map[string]Subscriber
}

func (p publisher) addSubscriber(subscriber Subscriber) {
	p.subscribers[subscriber.id()] = subscriber
}
func (p publisher) broadcast(msg string) {
	for _, v := range p.subscribers {
		v.react(msg)
	}
}

func newPublisher() publisher {
	return publisher{make(map[string]Subscriber)}
}

type Publisher interface {
	broadcast(msg string)
	addSubscriber(subcriber Subscriber)
	removeSubscriber(id string)
}

func main() {
	var p Publisher
	p = newPublisher()
	p.broadcast("hello")
	s := newSubscriber("23")
	s2 := newSubscriber("444")
	p.addSubscriber(s)
	p.addSubscriber(s2)
	p.removeSubscriber("23")
	p.broadcast("helllo")
}
