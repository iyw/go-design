package main

import "fmt"

type (
	Event struct {
		Data int64
	}
	Observer interface {
		Update(Event)
	}
	Subject interface {
		Attach(Observer)
		Dettch(Observer)
		Notify(Event)
	}
)

type (
	eventObserver struct {
		Id int
	}
	eventSubject struct {
		observers map[Observer]struct{}
	}
)

//Update
func (o *eventObserver) Update(d Event) {
	fmt.Printf("Observer %d from receiver: %d \n", o.Id, d.Data)
}

//attch
func (s *eventSubject) Attach(o Observer) {
	s.observers[o] = struct{}{}
}

//dettch
func (s *eventSubject) Dettch(o Observer) {
	delete(s.observers, o)
}

//notify
func (s *eventSubject) Notify(d Event) {
	for o := range s.observers {
		o.Update(d)
	}
}
func main() {
	s := eventSubject{
		observers: map[Observer]struct{}{},
	}
	s.Attach(&eventObserver{Id: 1})
	s.Attach(&eventObserver{Id: 2})
	s.Notify(Event{Data: 321})
}
