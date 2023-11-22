package main

import fmt "fmt"

// struct can have attributes of other structs
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	return true
}

// one time struct
// do not need it elsewhere in the program
type car struct {
	Make  string
	Model string
	// Wheel is a field containing anonymous struct
	Wheel struct {
		Radius   int
		Material string
	}
}

// method that takes a struct as a receiver
func (c car) getMake() string {
	return c.Make
}

func structTest() {
	myMessage := messageToSend{"ABC", user{}, user{}}
	myMessage.sender.number = 20
	myMessage.sender.name = "Xipeng"
	myMessage.recipient.name = "A"
	fmt.Println(canSendMessage(myMessage))

	myCar := car{
		Make:  "Tesla",
		Model: "Model 3",
		Wheel: struct {
			Radius   int
			Material string
		}{Radius: 20, Material: "Rubber"},
	}

	fmt.Println(myCar)
}
