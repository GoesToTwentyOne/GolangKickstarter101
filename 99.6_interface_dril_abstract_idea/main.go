package main

import "fmt"

func main() {
	var s, d writeMessage
	s.Name = "Rowjatul Jannat Neha"
	s.Message = "Hello How are you?"
	s.To = "Nihad"
	d.Name = "Neha"
	d.Message = "You are small baby"
	d.To = "Nihad Hossain"
	info(s)
	info(d)

	// s.send()

}

//Abstrction is the process of hiding the implementation details and showing only functionality to the users like sendmessage
//Atm booth,Mobile call function etc

type userMessage interface {
	send()
}
type writeMessage struct {
	Name    string
	Message string
	To      string
}

func (s writeMessage) send() {
	fmt.Println("Hi ! I am  ", s.Name)
	fmt.Println("Message ", s.Message)
	fmt.Println("Send to ", s.To)

}
func info(u userMessage) {
	fmt.Println(u)
	u.send()
}
