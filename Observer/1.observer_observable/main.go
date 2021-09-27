package main

import (
	"container/list"
	"fmt"
)

// Observable, Observer
//
// Example: An ill person would send a notification to the doctor whenever it get ill.
// Then the doctor can react to the event.

// The entity that is sending notifications to a set of subscribers
type Observable struct {
	subs *list.List
}

// Add observer to the list of subscribers
func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

// Remove observer to the list of subscribers
func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

// Fires notifications to all subscribers
func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

// All of the entities that wish to be informed
type Observer interface {
	Notify(data interface{}) // Can be any kind of data
}

// The ill patient, that can send events to doctor service.
type Person struct {
	Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		Name:       name,
	}
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

// The DoctorService is the observer, because it wishes to get notifications
type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s",
		data.(string))
}

func main() {
	p := NewPerson("Boris")
	ds := &DoctorService{}

	p.Subscribe(ds) // Doctor gets subcribed to person events.

	p.CatchACold() // Inform the doctor service.
}
