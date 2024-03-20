package function

import (
	"context"
	"fmt"

	"github.com/cloudevents/sdk-go/v2/event"
)

type Purchase struct {
	CustomerId string `json:"customerId"`
	ProductId  string `json:"productId"`
}

// Handle an event.
func Handle(ctx context.Context, e event.Event) (*event.Event, error) {
	/*
	 * YOUR CODE HERE
	 *
	 * Try running `go test`.  Add more test as you code in `handle_test.go`.
	 */

	fmt.Println("Received event")
	fmt.Println(e) // echo to local output
	purchase := &Purchase{}
	if err := e.DataAs(purchase); err != nil {
		fmt.Println("Failed to decode event: ", err.Error())
		return nil, err
	}

	fmt.Printf("Successfully decoded purchase: %+v", purchase)
	return &e, nil // echo to caller
}

/*
Other supported function signatures:

	Handle()
	Handle() error
	Handle(context.Context)
	Handle(context.Context) error
	Handle(event.Event)
	Handle(event.Event) error
	Handle(context.Context, event.Event)
	Handle(context.Context, event.Event) error
	Handle(event.Event) *event.Event
	Handle(event.Event) (*event.Event, error)
	Handle(context.Context, event.Event) *event.Event
	Handle(context.Context, event.Event) (*event.Event, error)

*/
