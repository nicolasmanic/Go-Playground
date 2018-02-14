package main

import "fmt"

type notifier interface {
	notify()
}

// user-defined admin that implements notifier
type admin struct {
	name string
}

//notifier interface is implemented with pointer receiver
//then only pointers of admin type will implement the notifier interface
func (u *admin) notify() {
	fmt.Printf("Admin: %v notified", u.name)
}

//Notifier is implemented by
func SendNotification(n notifier) {
	n.notify()
}
