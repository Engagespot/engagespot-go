# Engagespot Go SDK

This is the official Go SDK for Engagespot, a powerful notification infrastructure for developers.

## Installation

To install the Engagespot Go SDK, use `go get github.com/Engagespot/engagespot-go`

## Usage

### Initializing the client

First, import the package and create a new client:

```go
import "github.com/Engagespot/engagespot-go/engagespot"
client := engagespot.NewClient("YOUR_API_KEY", "YOUR_API_SECRET")
```

Replace `YOUR_API_KEY` and `YOUR_API_SECRET` with your actual Engagespot API credentials.

### Sending a notification

To send a notification using a workflow, use the `Send()` function:

```go
notification := &engagespot.Notification{
Workflow: engagespot.Workflow{
Identifier: "your-workflow-identifier",
},
Data: map[string]interface{}{
"key": "value"
},
SendTo: engagespot.SendTo{
Recipients: []string{"user-id-001"},
},
}
response, err := client.Notifications.Send(notification)
if err != nil {
log.Fatalf("Error sending notification: %v", err)
}
fmt.Printf("Notification sent successfully. Request ID: %s\n", response.RequestID)
```

### Full working example

```go
package main

import (
	"log"
	"fmt"
	"github.com/Engagespot/engagespot-go/engagespot"
)

func main(){
	client := engagespot.NewClient("YOUR_API_KEY", "YOUR_API_SECRET")

	notification := &engagespot.Notification{
		Workflow: engagespot.Workflow{
		Identifier: "workflow-identifier",
		},
		SendTo: engagespot.SendTo{
		Recipients: []string{"user-id-001"},
		},
		}
		response, err := client.Notifications.Send(notification)
		if err != nil {
		log.Fatalf("Error sending notification: %v", err)
		}
		fmt.Printf("Notification sent successfully. Request ID: %s\n", response.RequestID)
}
```