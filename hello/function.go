package hello

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/functions/metadata"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

func BackgroundEventFunction(ctx context.Context, m PubSubMessage) error {
	// Do something with ctx and data.
	fmt.Println("msg: ", m)
	fmt.Println("ctx: ", ctx)
	name := string(m.Data) // Automatically decoded from base64.
	if name == "" {
		name = "Hard Coded World"
	}
	log.Printf("Hello, %s!", name)

	meta, err := metadata.FromContext(ctx)
	if err != nil {
		fmt.Println("Error getting metadata.")
	} else {
		fmt.Println(" meta.EventID: ", meta.EventID)
		fmt.Println("meta.EventType: ", meta.EventType)
		fmt.Println("meta.Timestamp: ", meta.Timestamp)
		if meta.Resource != nil {
			fmt.Println("meta.Resource.Service: ", meta.Resource.Service)
			fmt.Println("meta.Resource.Name: ", meta.Resource.Name)
			fmt.Println("meta.Resource.Type: ", meta.Resource.Type)
			fmt.Println("meta.Resource.RawPath: ", meta.Resource.RawPath)
		}

	}

	return nil
}
