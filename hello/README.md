**Follow this document to create GO project which will receive published messages**

https://github.com/GoogleCloudPlatform/functions-framework-go

**Follow this document to use existing python sample script to publish messages**

https://cloud.google.com/pubsub/docs/emulator

Ths message gets delivered in below JSON format

```
{
    "message": {
        "attributes": {
            "greeting": "Hello from the Cloud Pub/Sub Emulator!"
        },
        "data": "SGVsbG8gQ2xvdWQgUHViL1N1YiEgSGVyZSBpcyBteSBtZXNzYWdlIQ==",
        "messageId": "136969346945"
    },
    "subscription": "projects/myproject/subscriptions/mysubscription"
}
```