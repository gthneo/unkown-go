package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connected:", pc.name)
}

func main() {
	a := PhoneConnector("PhoneConnector")
	a.Connect()
	Disconnect(a)
}

func Disconnect(usb USB) {
	fmt.Println("Disconnected.")
}
