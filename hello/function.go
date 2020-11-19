package hello

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"

	"cloud.google.com/go/functions/metadata"
)

/*
"message": {
    "attributes": {
      "greeting": "Hello from the Cloud Pub/Sub Emulator!"
    },
    "data": "SGVsbG8gQ2xvdWQgUHViL1N1YiEgSGVyZSBpcyBteSBtZXNzYWdlIQ==",
    "messageId": "136969346945"
  },
  "subscription": "projects/myproject/subscriptions/mysubscription"
*/
// PubSubMessageWrapper is the payload of a Pub/Sub event.
type PubSubMessageWrapper struct {
	Message pubsub.Message `json:"message"`
}

func BackgroundEventFunction(ctx context.Context, m PubSubMessageWrapper) error {
	// Do something with ctx and data.
	//fmt.Println("msg: ", m)
	//fmt.Println("ctx: ", ctx)
	msgData := string(m.Message.Data) // Automatically decoded from base64.
	//if msgData == "" {
	//	msgData = "Hard Coded World"
	//}
	log.Printf("Hello, %s!", msgData)
	log.Printf("m.Message.ID %s", m.Message.ID)

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

// HelloWorld writes "Hello, World!" to the HTTP response.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!\n")
}
