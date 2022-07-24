package observerpattern

import (
	"fmt"
)

type Subject interface {
	Attach(newObserver Observer)
	Detach(observerToDetach Observer)
	Notify(postTitle string)
}

type Observer interface {
	NotifyNewArticle(postTitle, authorName string)
	GetId() int
}

type MediumAutor struct {
	authorName string
	observers  []Observer
}

func (a *MediumAutor) Attach(newObserver Observer) {
	a.observers = append(a.observers, newObserver)
}

func (a *MediumAutor) Detach(observerToDetach Observer) {
	idToDetach := observerToDetach.GetId()
	for i, observer := range a.observers {
		if observer.GetId() == idToDetach {
			a.observers = append(a.observers[:i], a.observers[i+1:]...)
		}
	}
}

func (a MediumAutor) Notify(postTitle string) {
	for _, observer := range a.observers {
		observer.NotifyNewArticle(postTitle, a.authorName)
	}
}

func NewMediumAuthor(name string) (author MediumAutor) {
	return MediumAutor{authorName: name}
}

type Subscriber struct {
	Id   int
	name string
}

func (s Subscriber) NotifyNewArticle(postTitle, authorName string) {
	fmt.Printf("Hey %s: a new article named '%s' has been published by %s.\n", s.name, postTitle, authorName)
}

func (s Subscriber) GetId() int {
	return s.Id
}

var numSubscribers = 0

func NewSubscriber(subjects []Subject, name string) (sub Subscriber) {
	sub.Id = numSubscribers
	sub.name = name
	numSubscribers += 1

	for _, subject := range subjects {
		subject.Attach(sub)
	}

	return sub
}

func RunExample() {
	fmt.Println("Creating authors Rob Pike and Bill Gates...")
	robPike := NewMediumAuthor("Rob Pike")
	billGates := NewMediumAuthor("Bill Gates")

	andrea := NewSubscriber([]Subject{&robPike, &billGates}, "Andrea")
	andres := NewSubscriber([]Subject{&robPike}, "Andres")
	lorena := NewSubscriber([]Subject{&billGates}, "Lorena")
	jennifer := NewSubscriber([]Subject{&robPike, &billGates}, "Jennifer")

	fmt.Println("Andrea, andres and jennifer are suscribed to Rob Pike")
	fmt.Println("Andrea, lorena and jennifer are suscribed to Bill Gates")
	fmt.Println()

	fmt.Println("Rob Pike has just published his article named 'Golang generics v1.18'")
	robPike.Notify("Golang generics v1.18")
	fmt.Println()

	fmt.Println("Bill Gates has just published his article named 'Climate change challenges for humanity'")
	billGates.Notify("Climate change challenges for humanity")
	fmt.Println()

	robPike.Detach(andrea)
	robPike.Detach(andres)
	robPike.Detach(jennifer)
	fmt.Println("Andrea, Andres and Jennifer just unsubscribed from Rob Pike")
	fmt.Println("Rob Pike has just published his article named 'I do not have subs'")
	robPike.Notify("I do not have subs")
	fmt.Println()

	billGates.Detach(andrea)
	billGates.Detach(lorena)
	fmt.Println("Andrea, and Lorena just unsubscribed from Bill Gates")
	fmt.Println("Bill Gates has just published his article named 'just Jennifer is suscribed to me'")
	billGates.Notify("just Jennifer is suscribed to me")
}
