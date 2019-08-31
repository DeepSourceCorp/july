// File documentation

package main

import (
	"fmt"
	at "github.com/deepsourcelabs/artifacts/types"
	"github.com/streadway/amqp"
)

//
func anotherUnexportedFunction() {
	// documentation for this func will not be counted
	fmt.Println("Another exported function")
}

// Somehow use amqp lib, for external deps (public)
func somehowUseAMQP() {
	// Create connection
	rmqURL := "amqp://localhost:5672"

	_, err := amqp.Dial(rmqURL)
	if err != nil {
		t.Errorf("RabbitMQ: Error establishing connection: %v", err)
	}
}

// Somehow use a go module hosted privately
func somehowUseArtifacts() {
	status := at.Status{
		Code:     2001,
		HMessage: "yo yo message",
		Err:      "meh! no error",
	}

	fmt.Printf("%+v", status)
}
