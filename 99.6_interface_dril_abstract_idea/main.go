package main

import "fmt"

func main() {
	var s writeMessage
	s.Name = "Nihad"
	s.To = "Rowjatul Jannat Neha"

	s.send()

}

//Abstrction is the process of hiding the implementation details and showing only functionality to the users like sendmessage
//Atm both,Mobile call function etc

type userMessage interface {
	send()
}
type writeMessage struct {
	Name string
	To   string
}

func (s writeMessage) send() {
	fmt.Println("Hi ! I am  ", s.Name)
	fmt.Println("Send to ", s.To)

}
