package main

import "fmt"

type client struct {
}

func (c *client) insertComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertPort()
}

type computer interface {
	insertPort()
}

type mac struct {
}

func (m *mac) insertPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type windows struct{}

func (w *windows) insertUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertUSBPort()
}

func main() {

	client := &client{}
	mac := &mac{}

	client.insertComputer(mac)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}

	client.insertComputer(windowsMachineAdapter)
}
